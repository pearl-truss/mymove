import React, { Component, Fragment } from 'react';
import { Link } from 'react-router-dom';
import { get } from 'lodash';
import moment from 'moment';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';
import { indexBackupContacts } from 'scenes/ServiceMembers/ducks';
import ppmBlack from 'shared/icon/ppm-black.svg';

import './Review.css';
export class Summary extends Component {
  componentWillMount() {
    const service_member = get(this.props.loggedInUser, 'service_member');
    if (
      service_member &&
      !(this.props.currentBackupContacts && this.props.currentBackupContacts[0])
    ) {
      this.props.indexBackupContacts(service_member.id);
    }
  }
  componentWillUpdate(newProps) {
    const service_member = get(newProps.loggedInUser, 'service_member');
    if (get(this.props.loggedInUser, 'service_member') !== service_member) {
      this.props.indexBackupContacts(service_member.id);
    }
  }
  render() {
    const {
      currentPpm,
      currentBackupContacts,
      loggedInUser,
      currentOrders,
      schemaRank,
      schemaOrdersType,
      moveIsApproved,
    } = this.props;
    const yesNoMap = { true: 'Yes', false: 'No' };

    function getFullName() {
      const service_member = get(loggedInUser, 'service_member');
      if (!service_member) return;
      return `${service_member.first_name} ${service_member.middle_name ||
        ''} ${service_member.last_name} ${service_member.suffix || ''}`;
    }
    function getFullAddress(address) {
      if (address) {
        return (
          <Fragment>
            <div>{address.street_address_1}</div>
            {address.street_address_2 && <div>{address.street_address_2}</div>}
            <div>
              {address.city}, {address.state} {address.postal_code}
            </div>
          </Fragment>
        );
      }
    }
    function getFullContactPreferences() {
      const service_member = get(loggedInUser, 'service_member');
      if (!service_member) return;
      const prefs = {
        phone_is_preferred: 'Phone',
        text_message_is_preferred: 'Text',
        email_is_preferred: 'Email',
      };
      const preferredMethods = [];
      Object.keys(prefs).forEach(propertyName => {
        if (service_member[propertyName]) {
          preferredMethods.push(prefs[propertyName]);
        }
      });
      return preferredMethods.join(', ');
    }
    // TODO: Uncomment function below after backup contact auth is implemented.
    // function getFullBackupPermission(backup_contact) {
    //   const perms = {
    //     NONE: '',
    //     VIEW: 'View all aspects of this move',
    //     EDIT:
    //       'Authorized to represent me in all aspects of this move (letter of authorization)',
    //   };
    //   return `${perms[backup_contact.permission]}`;
    // }
    function formatDate(date) {
      if (!date) return;
      return moment(date, 'YYYY-MM-DD').format('MM/DD/YYYY');
    }

    const rootAddress = `/moves/${this.props.match.params.moveId}/review`;
    const editProfileAddress = rootAddress + '/edit-profile';
    const editBackupContactAddress = rootAddress + '/edit-backup-contact';
    const editContactInfoAddress = rootAddress + '/edit-contact-info';
    const editOrdersAddress = rootAddress + '/edit-orders';
    const editDateAndLocationAddress = rootAddress + '/edit-date-and-location';
    const editWeightAddress = rootAddress + '/edit-weight';
    const privateStorageString = get(
      currentPpm,
      'estimated_storage_reimbursement',
    )
      ? `(spend up to ${
          currentPpm.estimated_storage_reimbursement
        } on private storage)`
      : '';
    const sitDisplay = get(currentPpm, 'has_sit', false)
      ? `${currentPpm.days_in_storage} days ${privateStorageString}`
      : 'Not requested';
    return (
      <Fragment>
        <h3>Profile and Orders</h3>
        <div className="usa-grid-full review-content">
          <div className="usa-width-one-half review-section">
            <table>
              <tbody>
                <tr>
                  <th>
                    Profile{' '}
                    <span className="align-right">
                      <Link to={editProfileAddress}>Edit</Link>
                    </span>
                  </th>
                </tr>
                <tr>
                  <td> Name: </td>
                  <td>{getFullName()}</td>
                </tr>
                <tr>
                  <td>Branch:</td>
                  <td>{get(loggedInUser, 'service_member.affiliation')}</td>
                </tr>
                <tr>
                  <td> Rank/Pay Grade: </td>
                  <td>
                    {get(
                      schemaRank['x-display-value'],
                      get(loggedInUser, 'service_member.rank'),
                    )}
                  </td>
                </tr>
                <tr>
                  <td> DoD ID#: </td>
                  <td>{get(loggedInUser, 'service_member.edipi')}</td>
                </tr>
                <tr>
                  <td> Current Duty Station: </td>
                  <td>
                    {get(loggedInUser, 'service_member.current_station.name')}
                  </td>
                </tr>
              </tbody>
            </table>

            <table>
              <tbody>
                <tr>
                  <th>
                    Orders{moveIsApproved && '*'}
                    {!moveIsApproved && (
                      <span className="align-right">
                        <Link to={editOrdersAddress}>Edit</Link>
                      </span>
                    )}
                  </th>
                </tr>
                <tr>
                  <td> Orders Type: </td>
                  <td>
                    {get(
                      schemaOrdersType['x-display-value'],
                      get(currentOrders, 'orders_type'),
                    )}
                  </td>
                </tr>
                <tr>
                  <td> Orders Date: </td>
                  <td> {formatDate(get(currentOrders, 'issue_date'))}</td>
                </tr>
                <tr>
                  <td> Report-by Date: </td>
                  <td>{formatDate(get(currentOrders, 'report_by_date'))}</td>
                </tr>
                <tr>
                  <td> New Duty Station: </td>
                  <td> {get(currentOrders, 'new_duty_station.name')}</td>
                </tr>
                <tr>
                  <td> Dependents?: </td>
                  <td>
                    {' '}
                    {currentOrders &&
                      yesNoMap[get(currentOrders, 'has_dependents').toString()]}
                  </td>
                </tr>
                <tr>
                  <td> Orders Uploaded: </td>
                  <td>
                    {get(currentOrders, 'uploaded_orders.uploads') &&
                      get(currentOrders, 'uploaded_orders.uploads').length}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <div className="usa-width-one-half review-section">
            <table>
              <tbody>
                <tr>
                  <th>
                    Contact Info{' '}
                    <span className="align-right">
                      <Link to={editContactInfoAddress}>Edit</Link>
                    </span>
                  </th>
                </tr>
                <tr>
                  <td> Best Contact Phone: </td>
                  <td>{get(loggedInUser, 'service_member.telephone')}</td>
                </tr>
                <tr>
                  <td> Alt. Phone: </td>
                  <td>
                    {get(loggedInUser, 'service_member.secondary_telephone')}
                  </td>
                </tr>
                <tr>
                  <td> Personal Email: </td>
                  <td>{get(loggedInUser, 'service_member.personal_email')}</td>
                </tr>
                <tr>
                  <td> Preferred Contact Method: </td>
                  <td>{getFullContactPreferences()}</td>
                </tr>
                <tr>
                  <td> Current Mailing Address: </td>
                  <td>
                    {getFullAddress(
                      get(loggedInUser, 'service_member.residential_address'),
                    )}
                  </td>
                </tr>
                <tr>
                  <td> Backup Mailing Address: </td>
                  <td>
                    {getFullAddress(
                      get(
                        loggedInUser,
                        'service_member.backup_mailing_address',
                      ),
                    )}
                  </td>
                </tr>
              </tbody>
            </table>
            {currentBackupContacts.map(contact => (
              <table key={contact.id}>
                <tbody>
                  <tr>
                    <th>
                      Backup Contact{' '}
                      <span className="align-right">
                        <Link to={editBackupContactAddress}>Edit</Link>
                      </span>
                    </th>
                  </tr>
                  <tr>
                    <td> Backup Contact: </td>
                    <td>
                      {contact.name} <br />
                      {/* getFullBackupPermission(contact) */}
                    </td>
                  </tr>
                  <tr>
                    <td> Email: </td>
                    <td> {contact.email} </td>
                  </tr>
                  <tr>
                    <td> Phone: </td>
                    <td> {contact.telephone}</td>
                  </tr>
                </tbody>
              </table>
            ))}
          </div>
        </div>
        {currentPpm && (
          <div className="usa-grid-full ppm-container">
            <h3>
              <img src={ppmBlack} alt="PPM shipment" /> Shipment - You move your
              stuff (PPM)
            </h3>
            <div className="usa-width-one-half review-section ppm-review-section">
              <table>
                <tbody>
                  <tr>
                    <th>
                      Dates & Locations
                      <span className="align-right">
                        <Link to={editDateAndLocationAddress}>Edit</Link>
                      </span>
                    </th>
                  </tr>
                  <tr>
                    <td> Move Date: </td>
                    <td>{formatDate(get(currentPpm, 'planned_move_date'))}</td>
                  </tr>
                  <tr>
                    <td> Pickup ZIP Code: </td>
                    <td> {currentPpm && currentPpm.pickup_postal_code}</td>
                  </tr>
                  {currentPpm.has_additional_postal_code && (
                    <tr>
                      <td> Additional Pickup: </td>
                      <td> {currentPpm.additional_pickup_postal_code}</td>
                    </tr>
                  )}
                  <tr>
                    <td> Delivery ZIP Code: </td>
                    <td> {currentPpm && currentPpm.destination_postal_code}</td>
                  </tr>
                  <tr>
                    <td> Storage: </td>
                    <td>{sitDisplay}</td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div className="usa-width-one-half review-section ppm-review-section">
              <table>
                <tbody>
                  <tr>
                    <th>
                      Weight
                      <span className="align-right">
                        <Link to={editWeightAddress}>Edit</Link>
                      </span>
                    </th>
                  </tr>
                  <tr>
                    <td> Estimated Weight: </td>
                    <td> {currentPpm && currentPpm.weight_estimate} lbs</td>
                  </tr>
                  <tr>
                    <td> Estimated PPM Incentive: </td>
                    <td> {currentPpm && currentPpm.estimated_incentive}</td>
                  </tr>
                  {currentPpm.has_requested_advance && (
                    <tr>
                      <td> Advance: </td>
                      <td> ${currentPpm.advance.requested_amount}</td>
                    </tr>
                  )}
                </tbody>
              </table>
            </div>
          </div>
        )}
        {moveIsApproved && (
          <div className="approved-edit-warning">
            *To change these fields, contact your local PPPO office.
          </div>
        )}
      </Fragment>
    );
  }
}

Summary.propTypes = {
  currentPpm: PropTypes.object,
  currentBackupContacts: PropTypes.array,
  currentOrders: PropTypes.object,
  schemaRank: PropTypes.object,
  schemaOrdersType: PropTypes.object,
  moveIsApproved: PropTypes.boolean,
};

function mapDispatchToProps(dispatch) {
  return bindActionCreators({ indexBackupContacts }, dispatch);
}

function mapStateToProps(state) {
  return {
    ...state.ppm,
    ...state.loggedInUser,
    currentBackupContacts: state.serviceMember.currentBackupContacts,
    currentOrders:
      get(state.orders, 'currentOrders') ||
      get(state.loggedInUser, 'loggedInUser.service_member.orders[0]'),
    schemaRank: get(state, 'swagger.spec.definitions.ServiceMemberRank', {}),
    schemaOrdersType: get(state, 'swagger.spec.definitions.OrdersType', {}),
    //todo: this should be a selector
    moveIsApproved:
      get(
        state,
        'loggedInUser.loggedInUser.service_member.orders.0.moves.0.status',
        'SUBMITTED',
      ) === 'APPROVED',
  };
}
export default withRouter(
  connect(mapStateToProps, mapDispatchToProps)(Summary),
);
