import React, { Component } from 'react';
import { withRouter } from 'react-router';
import ReactTable from 'react-table';
import { connect } from 'react-redux';
import { get } from 'lodash';
import 'react-table/react-table.css';
import Alert from 'shared/Alert';
import { formatTimeAgo } from 'shared/formatters';
import { newColumns, ppmColumns, hhgActiveColumns, defaultColumns } from './queueTableColumns';

import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faSyncAlt from '@fortawesome/fontawesome-free-solid/faSyncAlt';
import { getEntitlements } from '../../shared/entitlements';
import moment from 'moment';
import { formatDate4DigitYear } from '../../shared/formatters';
import { sitDaysUsed, sitTotalDaysUsed } from '../../shared/StorageInTransit/calculator';

class QueueTable extends Component {
  constructor() {
    super();
    this.state = {
      data: [],
      pages: null,
      loading: true,
      refreshing: false, // only true when the user clicks the refresh button
      lastLoadedAt: new Date(),
      lastLoadedAtText: formatTimeAgo(new Date()),
      interval: setInterval(() => {
        this.setState({
          lastLoadedAtText: formatTimeAgo(this.state.lastLoadedAt),
        });
      }, 5000),
    };
    this.fetchData = this.fetchData.bind(this);
  }

  componentDidMount() {
    this.fetchData();
  }

  componentDidUpdate(prevProps) {
    if (this.props.queueType !== prevProps.queueType) {
      this.fetchData();
    }
  }

  openMove(rowInfo) {
    this.props.history.push(`new/moves/${rowInfo.original.id}`, {
      referrerPathname: this.props.history.location.pathname,
    });
  }

  static defaultProps = {
    moveLocator: '',
    firstName: '',
    lastName: '',
  };

  async fetchData() {
    const loadingQueueType = this.props.queueType;

    this.setState({
      data: [],
      pages: null,
      loading: true,
      loadingQueue: loadingQueueType,
    });

    // Catch any errors here and render an empty queue
    try {
      const body = await this.props.retrieveMoves(this.props.queueType);

      // Only update the queue list if the request that is returning
      // is for the same queue as the most recent request.
      if (this.state.loadingQueue === loadingQueueType) {
        this.setState({
          data: body,
          pages: 1,
          loading: false,
          refreshing: false,
          lastLoadedAt: new Date(),
        });
      }
    } catch (e) {
      this.setState({
        data: [],
        pages: 1,
        loading: false,
        refreshing: false,
        lastLoadedAt: new Date(),
      });
    }
  }

  refresh() {
    clearInterval(this.state.interval);

    this.setState({
      refreshing: true,
      lastLoadedAt: new Date(),
      interval: setInterval(() => {
        this.setState({
          lastLoadedAtText: formatTimeAgo(this.state.lastLoadedAt),
        });
      }, 5000),
    });

    this.fetchData();
  }

  render() {
    const titles = {
      new: 'New Moves/Shipments',
      troubleshooting: 'Troubleshooting',
      ppm: 'PPMs',
      hhg_active: 'Active HHGs',
      hhg_approved: 'Approved HHGs',
      hhg_delivered: 'Delivered HHGs',
      all: 'All Moves',
    };

    const showColumns = queueType => {
      switch (queueType) {
        case 'new':
          return newColumns;
        case 'ppm':
          return ppmColumns;
        case 'hhg_active':
          return hhgActiveColumns;
        default:
          return defaultColumns;
      }
    };

    this.state.data.forEach(row => {
      if (this.props.queueType === 'new' && row.ppm_status && row.hhg_status) {
        row.shipments = 'HHG, PPM';
      } else if (row.ppm_status && !row.hhg_status) {
        row.shipments = 'PPM';
      } else {
        row.shipments = 'HHG';
      }

      if (this.props.queueType === 'ppm' && row.ppm_status !== null) {
        row.synthetic_status = row.ppm_status;
      } else {
        row.synthetic_status = row.status;
      }

      if (this.props.queueType === 'hhg_active' && row.storage_in_transits) {
        row.sit_expires = formatDate4DigitYear(
          moment.min(
            row.storage_in_transits.map(sit => {
              return moment(sit.actual_start_date).add(
                getEntitlements(row.rank).storage_in_transit +
                  sitDaysUsed(sit) -
                  sitTotalDaysUsed(row.storage_in_transits),
                'days',
              );
            }),
          ),
        );
      }
    });

    return (
      <div>
        {this.props.showFlashMessage ? (
          <Alert type="success" heading="Success">
            {this.props.flashMessageLines.join('\n')}
            <br />
          </Alert>
        ) : null}
        <h1 className="queue-heading">{titles[this.props.queueType]}</h1>
        <div className="queue-table">
          <span className="staleness-indicator" data-cy="staleness-indicator">
            Last updated {formatTimeAgo(this.state.lastLoadedAt)}
          </span>
          <span className={'refresh' + (this.state.refreshing ? ' focused' : '')} title="Refresh" aria-label="Refresh">
            <FontAwesomeIcon
              data-cy="refreshQueue"
              className="link-blue"
              icon={faSyncAlt}
              onClick={this.refresh.bind(this)}
              color="blue"
              size="lg"
              spin={!this.state.refreshing && this.state.loading}
            />
          </span>
          <ReactTable
            columns={showColumns(this.props.queueType)}
            data={this.state.data}
            loading={this.state.loading} // Display the loading overlay when we need it
            defaultSorted={[{ id: 'move_date', asc: true }]}
            pageSize={this.state.data.length}
            className="-striped -highlight"
            showPagination={false}
            getTrProps={(state, rowInfo) => ({
              'data-cy': 'queueTableRow',
              onDoubleClick: () => this.openMove(rowInfo),
              onClick: () => this.openMove(rowInfo),
            })}
          />
        </div>
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    showFlashMessage: get(state, 'flashMessages.display', false),
    flashMessageLines: get(state, 'flashMessages.messageLines', false),
  };
};

export default withRouter(connect(mapStateToProps)(QueueTable));
