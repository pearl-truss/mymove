import React, { Component, Fragment } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { get, isNull, toUpper } from 'lodash';
import PropTypes from 'prop-types';
import { reduxifyWizardForm } from 'shared/WizardPage/Form';
import Alert from 'shared/Alert';
import { formatCentsRange } from 'shared/formatters';
import { loadEntitlementsFromState } from 'shared/entitlements';
import {
  loadPPMs,
  updatePPM,
  selectActivePPMForMove,
  updatePPMEstimate,
  getPpmWeightEstimate,
  selectPPMEstimateRange,
} from 'shared/Entities/modules/ppms';
import IconWithTooltip from 'shared/ToolTip/IconWithTooltip';
import RadioButton from 'shared/RadioButton';
import 'react-rangeslider/lib/index.css';
import styles from './Weight.module.scss';
import { withContext } from 'shared/AppContext';
import RangeSlider from 'shared/RangeSlider';
import { hasShortHaulError } from 'shared/incentive';
import carGray from 'shared/icon/car-gray.svg';
import trailerGray from 'shared/icon/trailer-gray.svg';
import truckGray from 'shared/icon/truck-gray.svg';

const WeightWizardForm = reduxifyWizardForm('weight-wizard-form');

export class PpmWeight extends Component {
  constructor(props) {
    super(props);

    this.state = {
      pendingPpmWeight: 0,
      includesProgear: 'No',
      isProgearMoreThan1000: 'No',
      hasEstimateError: false,
    };

    this.onWeightSelected = this.onWeightSelected.bind(this);
  }

  getDefaultWeightClassMedian() {
    const { entitlement, currentPPM } = this.props;
    if (isNull(entitlement) || isNull(currentPPM)) {
      return null;
    }
    const defaultWeightRange = {
      min: 500,
      max: 2500,
    };

    const median = defaultWeightRange.min + (defaultWeightRange.max - defaultWeightRange.min) / 2;
    return median;
  }

  componentDidMount() {
    const { currentPPM } = this.props;
    const moveId = this.props.match.params.moveId;
    this.props.loadPPMs(moveId);

    if (currentPPM) {
      this.setState(
        {
          pendingPpmWeight:
            currentPPM.weight_estimate && currentPPM.weight_estimate !== 0
              ? currentPPM.weight_estimate
              : this.getDefaultWeightClassMedian(),
        },
        this.updateIncentive,
      );
    }
  }

  componentDidUpdate(prevProps, prevState) {
    const { currentPPM, hasLoadSuccess } = this.props;
    if (!prevProps.hasLoadSuccess && hasLoadSuccess && currentPPM) {
      this.setState(
        {
          pendingPpmWeight:
            currentPPM.weight_estimate && currentPPM.weight_estimate !== 0
              ? currentPPM.weight_estimate
              : this.getDefaultWeightClassMedian(),
        },
        this.updateIncentive,
      );
    }
  }

  // this method is used to set the incentive on page load
  // it runs even if the incentive has been set before since data changes on previous pages could
  // affect it
  updateIncentive = () => {
    const { currentPPM, originDutyStationZip, tempCurrentPPM } = this.props;
    const weight = this.state.pendingPpmWeight;

    // TODO this is a work around till we refactor more SM data...
    const origMoveDate =
      currentPPM && currentPPM.hasOwnProperty('original_move_date')
        ? currentPPM.original_move_date
        : tempCurrentPPM.original_move_date;
    // TODO this is a work around till we refactor more SM data...
    const pickupPostalCode =
      currentPPM && currentPPM.hasOwnProperty('pickup_postal_code')
        ? currentPPM.pickup_postal_code
        : tempCurrentPPM.pickup_postal_code;

    this.props
      .getPpmWeightEstimate(origMoveDate, pickupPostalCode, originDutyStationZip, this.props.orders.id, weight)
      .catch(() => this.setState({ hasEstimateError: true }));
  };

  handleSubmit = () => {
    const ppmBody = {
      weight_estimate: parseInt(this.state.pendingPpmWeight),
      has_requested_advance: false,
      has_pro_gear: toUpper(this.state.includesProgear),
      has_pro_gear_over_thousand: toUpper(this.state.isProgearMoreThan1000),
    };

    // TODO this is a work around till we refactor more SM data...
    const ppmId = this.props.currentPPM.id ? this.props.currentPPM.id : this.props.tempCurrentPPM.id;
    // TODO this is a work around till we refactor more SM data...
    const moveId = this.props.currentPPM.move_id ? this.props.currentPPM.move_id : this.props.tempCurrentPPM.move_id;
    return this.props
      .updatePPM(moveId, ppmId, ppmBody)
      .then(({ response }) => this.props.updatePPMEstimate(moveId, response.body.id).catch((err) => err));
    // catch block returns error so that the wizard can continue on with its flow
  };

  onWeightSelected = (event) => {
    const weight = event.target.value;
    this.setState(
      {
        pendingPpmWeight: weight,
      },
      () => this.updateIncentive(),
    );
  };

  chooseVehicleIcon(currentEstimate) {
    if (currentEstimate < 500) {
      return <img className="icon" src={carGray} alt="car-gray" data-cy="vehicleIcon" />;
    }
    if (currentEstimate >= 500 && currentEstimate < 1500) {
      return <img className="icon" src={trailerGray} alt="trailer-gray" data-cy="vehicleIcon" />;
    }
    if (currentEstimate >= 1500) {
      return <img className="icon" src={truckGray} alt="truck-gray" data-cy="vehicleIcon" />;
    }
  }

  chooseEstimateText(currentEstimate) {
    if (currentEstimate < 500) {
      return <p data-cy="estimateText">Just a few things. One trip in a car.</p>;
    }
    if (currentEstimate >= 500 && currentEstimate < 1000) {
      return (
        <p data-cy="estimateText">
          Studio apartment, minimal stuff. A large car, a pickup, a van, or a car with trailer.
        </p>
      );
    }
    if (currentEstimate >= 1000 && currentEstimate < 2000) {
      return (
        <p data-cy="estimateText">
          1-2 rooms, light furniture. A pickup, a van, or a car with a small or medium trailer.
        </p>
      );
    }
    if (currentEstimate >= 2000 && currentEstimate < 3000) {
      return (
        <p data-cy="estimateText">
          2-3 rooms, some bulky items. Cargo van, small or medium moving truck, medium or large cargo trailer.
        </p>
      );
    }
    if (currentEstimate >= 3000 && currentEstimate < 4000) {
      return <p data-cy="estimateText">3-4 rooms. Small to medium moving truck, or a couple of trips.</p>;
    }
    if (currentEstimate >= 4000 && currentEstimate < 5000) {
      return (
        <p data-cy="estimateText">
          4+ rooms, or just a lot of large, heavy things. Medium or large moving truck, or multiple trips.
        </p>
      );
    }
    if (currentEstimate >= 5000 && currentEstimate < 6000) {
      return (
        <p data-cy="estimateText">
          Many rooms, many things, lots of them heavy. Medium or large moving truck, or multiple trips.
        </p>
      );
    }
    if (currentEstimate >= 6000 && currentEstimate < 7000) {
      return (
        <p data-cy="estimateText">
          Large house, a lot of things. The biggest rentable moving trucks, or multiple trips or vehicles.
        </p>
      );
    }
    if (currentEstimate >= 7000) {
      return (
        <p data-cy="estimateText">
          A large house or small palace, many heavy or bulky items. Multiple trips using large vehicles, or hire
          professional movers.
        </p>
      );
    }
  }

  chooseIncentiveRangeText(hasEstimateError) {
    const { incentiveEstimateMin, incentiveEstimateMax } = this.props;
    if (hasEstimateError) {
      return (
        <div className="incentive">
          Not ready yet{' '}
          <IconWithTooltip toolTipText="We expect to receive rate data covering your move dates by the end of this month. Check back then to see your estimated incentive." />
        </div>
      );
    } else {
      return (
        <div data-cy="incentive-range-values" className="incentive">
          {formatCentsRange(incentiveEstimateMin, incentiveEstimateMax)}
        </div>
      );
    }
  }

  handleChange = (event, type) => {
    this.setState({ [type]: event.target.value });
  };

  chooseEstimateErrorText(hasEstimateError, rateEngineError) {
    if (hasShortHaulError(rateEngineError)) {
      return (
        <Fragment>
          <div className="error-message">
            <Alert type="warning" heading="Could not retrieve estimate">
              MilMove does not presently support short-haul PPM moves. Please contact your PPPO.
            </Alert>
          </div>
        </Fragment>
      );
    }

    if (rateEngineError || hasEstimateError) {
      return (
        <Fragment>
          <div className="error-message">
            <Alert type="warning" heading="Could not retrieve estimate">
              There was an issue retrieving an estimate for your incentive. You still qualify, but need to talk with
              your local transportation office which you can look up on{' '}
              <a href="move.mil" className="usa-link">
                move.mil
              </a>
            </Alert>
          </div>
        </Fragment>
      );
    }
  }

  render() {
    const {
      incentiveEstimateMin,
      incentiveEstimateMax,
      pages,
      pageKey,
      hasEstimateInProgress,
      error,
      rateEngineError,
    } = this.props;
    const { includesProgear, isProgearMoreThan1000, hasEstimateError } = this.state;

    return (
      <div>
        <div className="grid-container usa-prose">
          <WeightWizardForm
            handleSubmit={this.handleSubmit}
            pageList={pages}
            pageKey={pageKey}
            serverError={error}
            additionalValues={{
              hasEstimateInProgress,
              incentiveEstimateMax,
            }}
            readyToSubmit={!hasShortHaulError(rateEngineError)}
          >
            <h3 data-cy="weight-page-title">How much do you think you'll move?</h3>
            <p>Your weight entitlement: {this.props.entitlement.weight.toLocaleString()} lbs</p>
            <div>
              <RangeSlider
                id="incentive-estimation-slider"
                max={this.props.entitlement.weight}
                step={this.props.entitlement.weight <= 2500 ? 100 : 500}
                min={0}
                defaultValue={this.state.pendingPpmWeight}
                prependTooltipText="about"
                appendTooltipText="lbs"
                onChange={this.onWeightSelected}
              />
              {this.chooseEstimateErrorText(hasEstimateError, rateEngineError)}
            </div>
            <div className={`${styles['incentive-estimate-box']} border radius-lg border-base`}>
              {this.chooseVehicleIcon(this.state.pendingPpmWeight)}
              {this.chooseEstimateText(this.state.pendingPpmWeight)}
              <h4>Your incentive for moving {this.state.pendingPpmWeight} lbs:</h4>
              <h3 className={styles['incentive-range-text']} data-cy="incentive-range-text">
                {this.chooseIncentiveRangeText(hasEstimateError)}
              </h3>
              <p className="text-gray-50">Final payment will be based on the weight you actually move.</p>
            </div>
            <div className="radio-group-wrapper normalize-margins">
              <h3>Do you also have pro-gear to move?</h3>
              <RadioButton
                inputClassName="usa-radio__input inline_radio"
                labelClassName="usa-radio__label inline_radio"
                label="Yes"
                value="Yes"
                name="includesProgear"
                checked={includesProgear === 'Yes'}
                onChange={(event) => this.handleChange(event, 'includesProgear')}
              />

              <RadioButton
                inputClassName="usa-radio__input inline_radio"
                labelClassName="usa-radio__label inline_radio"
                label="No"
                value="No"
                name="includesProgear"
                checked={includesProgear === 'No'}
                onChange={(event) => this.handleChange(event, 'includesProgear')}
              />
              <RadioButton
                inputClassName="usa-radio__input inline_radio"
                labelClassName="usa-radio__label inline_radio"
                label="Not sure"
                value="Not Sure"
                name="includesProgear"
                checked={includesProgear === 'Not Sure'}
                onChange={(event) => this.handleChange(event, 'includesProgear')}
              />
              <p>
                Books, papers, and equipment needed for official duties. <a href="#">What counts as pro-gear?</a>{' '}
              </p>
            </div>
            {(includesProgear === 'Yes' || includesProgear === 'Not Sure') && (
              <>
                <div className={`${styles['incentive-estimate-box']} border radius-lg border-base`}>
                  You can be paid for moving up to 2,500 lbs of pro-gear, in addition to your weight entitlement of{' '}
                  {formatCentsRange(incentiveEstimateMin, incentiveEstimateMax)}.
                </div>
                <div className="radio-group-wrapper normalize-margins">
                  <h3>Do you think your pro-gear weighs 1000 lbs or more?</h3>
                  <p>You can move up to 2000 lbs of qualified pro-gear, plus 500 pounds for your spouse.</p>
                  <RadioButton
                    inputClassName="usa-radio__input inline_radio"
                    labelClassName="usa-radio__label inline_radio"
                    label="Yes"
                    value="Yes"
                    name="isProgearMoreThan1000"
                    checked={isProgearMoreThan1000 === 'Yes'}
                    onChange={(event) => this.handleChange(event, 'isProgearMoreThan1000')}
                  />
                  <RadioButton
                    inputClassName="usa-radio__input inline_radio"
                    labelClassName="usa-radio__label inline_radio"
                    label="No"
                    value="No"
                    name="isProgearMoreThan1000"
                    checked={isProgearMoreThan1000 === 'No'}
                    onChange={(event) => this.handleChange(event, 'isProgearMoreThan1000')}
                  />
                  <RadioButton
                    inputClassName="usa-radio__input inline_radio"
                    labelClassName="usa-radio__label inline_radio"
                    label="Not sure"
                    value="Not Sure"
                    name="isProgearMoreThan1000"
                    checked={isProgearMoreThan1000 === 'Not Sure'}
                    onChange={(event) => this.handleChange(event, 'isProgearMoreThan1000')}
                  />
                </div>
              </>
            )}
            {(isProgearMoreThan1000 === 'Yes' || isProgearMoreThan1000 === 'Not Sure') && (
              <>
                <div className={`${styles['incentive-estimate-box']} border radius-lg border-base`}>
                  Pack your pro-gear separately. It might need to be weighed and verified.
                </div>
              </>
            )}
          </WeightWizardForm>
        </div>
      </div>
    );
  }
}

PpmWeight.propTypes = {
  currentPpm: PropTypes.shape({
    id: PropTypes.string,
    weight: PropTypes.number,
    incentive: PropTypes.string,
  }),
  hasLoadSuccess: PropTypes.bool.isRequired,
  currentPPM: PropTypes.object.isRequired,
};
function mapStateToProps(state) {
  const schema = get(state, 'swaggerInternal.spec.definitions.UpdatePersonallyProcuredMovePayload', {});
  const originDutyStationZip = state.serviceMember.currentServiceMember.current_station.address.postal_code;
  const moveID = state.moves.currentMove.id;

  const props = {
    ...state.ppm,
    incentiveEstimateMin: selectPPMEstimateRange(state).range_min,
    incentiveEstimateMax: selectPPMEstimateRange(state).range_max,
    currentPPM: selectActivePPMForMove(state, moveID),
    entitlement: loadEntitlementsFromState(state),
    schema: schema,
    originDutyStationZip,
    orders: get(state, 'orders.currentOrders', {}),
    // TODO this is a work around till we refactor more SM data...
    tempCurrentPPM: get(state, 'ppm.currentPpm'),
  };

  return props;
}

function mapDispatchToProps(dispatch) {
  return bindActionCreators(
    {
      loadPPMs,
      getPpmWeightEstimate,
      updatePPM,
      updatePPMEstimate,
    },
    dispatch,
  );
}

export default withContext(connect(mapStateToProps, mapDispatchToProps)(PpmWeight));
