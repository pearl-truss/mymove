import React, { Component, Fragment } from 'react';
import { get } from 'lodash';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';

import { getInternalSwaggerDefinition } from 'shared/Swagger/selectors';
import { loadMove, selectMove } from 'shared/Entities/modules/moves';
import {
  fetchLatestOrders,
  getLatestOrdersLabel,
  selectOrdersFromServiceMemberId,
} from 'shared/Entities/modules/orders';
import { getRequestStatus } from 'shared/Swagger/selectors';

// import { getPPM } from 'scenes/Moves/Ppm/ducks.js';
import { moveIsApproved, lastMoveIsCanceled } from 'scenes/Moves/ducks';
import { loadEntitlementsFromState } from 'shared/entitlements';
import Alert from 'shared/Alert';
import { titleCase } from 'shared/constants.js';

import { checkEntitlement } from './ducks';
import ServiceMemberSummary from './ServiceMemberSummary';
import PPMShipmentSummary from './PPMShipmentSummary';

import './Review.css';
import { selectActivePPMForMove } from '../../shared/Entities/modules/ppms';

export class Summary extends Component {
  componentDidMount() {
    // this.props.fetchLatestOrders(this.props.serviceMember.id);
    if (this.props.onDidMount) {
      this.props.onDidMount(this.props.serviceMember.id);
    }
  }
  componentDidUpdate(prevProps) {
    // Only check entitlement for PPMs, not HHGs
    if (prevProps.currentPPM !== this.props.currentPPM) {
      this.props.onCheckEntitlement(this.props.match.params.moveId);
    }
  }

  render() {
    const {
      currentMove,
      currentPPM,
      currentBackupContacts,
      currentOrders,
      schemaRank,
      schemaAffiliation,
      schemaOrdersType,
      moveIsApproved,
      serviceMember,
      entitlement,
      match,
    } = this.props;
    const currentStation = get(serviceMember, 'current_station');
    const stationPhone = get(currentStation, 'transportation_office.phone_lines.0');

    const rootAddressWithMoveId = `/moves/${this.props.match.params.moveId}/review`;
    // isReviewPage being false is the same thing as being in the /edit route
    const isReviewPage = rootAddressWithMoveId === match.url;
    const editSuccessBlurb = this.props.reviewState.editSuccess ? 'Your changes have been saved. ' : '';
    const editOrdersPath = rootAddressWithMoveId + '/edit-orders';

    const showPPMShipmentSummary =
      (isReviewPage && currentPPM) || (!isReviewPage && currentPPM && currentPPM.status !== 'DRAFT');

    const showProfileAndOrders = isReviewPage || !isReviewPage;
    return (
      <Fragment>
        {get(this.props.reviewState.error, 'statusCode', false) === 409 && (
          <Alert type="warning" heading={editSuccessBlurb + 'Your estimated weight is above your entitlement.'}>
            {titleCase(this.props.reviewState.error.response.body.message)}.
          </Alert>
        )}
        {this.props.reviewState.editSuccess &&
          !this.props.reviewState.entitlementChange &&
          get(this.props.reviewState.error, 'statusCode', false) === false && (
            <Alert type="success" heading={editSuccessBlurb} />
          )}
        {currentMove &&
          this.props.reviewState.entitlementChange &&
          get(this.props.reviewState.error, 'statusCode', false) === false && (
            <Alert type="info" heading={editSuccessBlurb + 'Note that the entitlement has also changed.'}>
              Your weight entitlement is now {entitlement.sum.toLocaleString()} lbs.
            </Alert>
          )}

        {showProfileAndOrders && (
          <ServiceMemberSummary
            orders={currentOrders}
            backupContacts={currentBackupContacts}
            serviceMember={serviceMember}
            schemaRank={schemaRank}
            schemaAffiliation={schemaAffiliation}
            schemaOrdersType={schemaOrdersType}
            moveIsApproved={moveIsApproved}
            editOrdersPath={editOrdersPath}
          />
        )}

        {showPPMShipmentSummary && (
          <PPMShipmentSummary ppm={currentPPM} movePath={rootAddressWithMoveId} orders={currentOrders} />
        )}
        {moveIsApproved && (
          <div className="approved-edit-warning">
            *To change these fields, contact your local PPPO office at {get(currentStation, 'name')}{' '}
            {stationPhone ? ` at ${stationPhone}` : ''}.
          </div>
        )}
      </Fragment>
    );
  }
}

Summary.propTypes = {
  currentBackupContacts: PropTypes.array,
  getCurrentMove: PropTypes.func,
  currentOrders: PropTypes.object,
  currentPPM: PropTypes.object,
  schemaRank: PropTypes.object,
  schemaOrdersType: PropTypes.object,
  moveIsApproved: PropTypes.bool,
  lastMoveIsCanceled: PropTypes.bool,
  error: PropTypes.object,
};

function mapStateToProps(state, ownProps) {
  const moveID = state.moves.currentMove.id;
  const showOrdersRequest = getRequestStatus(state, getLatestOrdersLabel);
  const serviceMemberId = get(state, 'serviceMember.currentServiceMember.id');

  return {
    currentPPM: selectActivePPMForMove(state, moveID),
    serviceMember: state.serviceMember.currentServiceMember,
    currentMove: selectMove(state, ownProps.match.params.moveId),
    currentBackupContacts: state.serviceMember.currentBackupContacts,
    // currentOrders: state.orders.currentOrders, // in master
    currentOrders: selectOrdersFromServiceMemberId(state, serviceMemberId),
    schemaRank: getInternalSwaggerDefinition(state, 'ServiceMemberRank'),
    schemaOrdersType: getInternalSwaggerDefinition(state, 'OrdersType'),
    schemaAffiliation: getInternalSwaggerDefinition(state, 'Affiliation'),
    moveIsApproved: moveIsApproved(state),
    lastMoveIsCanceled: lastMoveIsCanceled(state),
    reviewState: state.review,
    entitlement: loadEntitlementsFromState(state),
    loadDependenciesHasSuccess: showOrdersRequest.isSuccess,
    loadDependenciesHasError: showOrdersRequest.error,
  };
}
function mapDispatchToProps(dispatch, ownProps) {
  return {
    onDidMount: function (smId) {
      const moveID = ownProps.match.params.moveId;
      dispatch(loadMove(moveID, 'Summary.getMove'));
      dispatch(fetchLatestOrders(smId));
    },
    onCheckEntitlement: (moveId) => {
      dispatch(checkEntitlement(moveId));
    },
  };
}
export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Summary));
