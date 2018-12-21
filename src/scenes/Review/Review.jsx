import { no_op } from 'shared/utils';
import WizardPage from 'shared/WizardPage';
import React, { Component } from 'react';
import Summary from './Summary';
import { connect } from 'react-redux';
import { get } from 'lodash';
import WizardHeader from 'scenes/Moves/WizardHeader';
import { ProgressTimeline, ProgressTimelineStep } from 'shared/ProgressTimeline';
import reviewGray from 'shared/icon/review-gray.svg';
import './Review.css';

class Review extends Component {
  componentDidMount() {
    window.scrollTo(0, 0);
  }
  render() {
    const { pages, pageKey, isHHGPPMComboMove } = this.props;

    return (
      <div>
        {isHHGPPMComboMove && (
          <WizardHeader
            icon={reviewGray}
            title="Review"
            right={
              <ProgressTimeline>
                <ProgressTimelineStep name="Move Setup" completed />
                <ProgressTimelineStep name="Review" current />
              </ProgressTimeline>
            }
          />
        )}
        <WizardPage handleSubmit={no_op} pageList={pages} pageKey={pageKey} pageIsValid={true}>
          <div className="edit-title">
            <h2>Review move details</h2>
            <p>You're almost done! Please review your details before we finalize the move.</p>
          </div>
          <Summary />
        </WizardPage>
      </div>
    );
  }
}

const mapStateToProps = (state, ownProps) => ({
  isHHGPPMComboMove: get(state, 'moves.currentMove.selected_move_type') === 'HHG_PPM',
  ...ownProps,
});

export default connect(mapStateToProps)(Review);
