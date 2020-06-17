/* eslint-disable  */
import React, { Component, lazy, Suspense } from 'react';
import { Route, Switch } from 'react-router-dom';
import { ConnectedRouter } from 'connected-react-router';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import { history } from 'shared/store';

import 'uswds';
import '../../../node_modules/uswds/dist/css/uswds.css';
import 'scenes/Office/office.scss';

import { getCurrentUserInfo, selectCurrentUser } from 'shared/Data/users';
import { loadInternalSchema, loadPublicSchema } from 'shared/Swagger/ducks';
import { detectIE11 } from 'shared/utils';
import ConnectedLogoutOnInactivity from 'shared/User/LogoutOnInactivity';
import PrivateRoute from 'shared/User/PrivateRoute';
import { isProduction } from 'shared/constants';
import SomethingWentWrong from 'shared/SomethingWentWrong';
import QueueHeader from 'shared/Header/Office';
import FOUOHeader from 'components/FOUOHeader';
import { ConnectedSelectApplication } from 'pages/SelectApplication/SelectApplication';
import { roleTypes } from 'constants/userRoles';
import LoadingPlaceholder from 'shared/LoadingPlaceholder';
import { withContext } from 'shared/AppContext';

// Lazy load these dependencies
const ConnectedOfficeHome = lazy(() => import('pages/OfficeHome'));
const MoveInfo = lazy(() => import('./MoveInfo'));
const Queues = lazy(() => import('./Queues'));
const OrdersInfo = lazy(() => import('./OrdersInfo'));
const DocumentViewer = lazy(() => import('./DocumentViewer'));
const ScratchPad = lazy(() => import('shared/ScratchPad'));
const CustomerDetails = lazy(() => import('scenes/Office/TOO/customerDetails'));
const TOO = lazy(() => import('scenes/Office/TOO/too'));
const TOOMoveTaskOrder = lazy(() => import('pages/TOO/moveTaskOrder'));
const TIO = lazy(() => import('scenes/Office/TIO/tio'));
const TOOVerificationInProgress = lazy(() => import('scenes/Office/TOO/tooVerificationInProgress'));
const PaymentRequestShow = lazy(() => import('scenes/Office/TIO/paymentRequestShow'));
const PaymentRequestIndex = lazy(() => import('scenes/Office/TIO/paymentRequestIndex'));
const MoveDetails = lazy(() => import('pages/TOO/moveDetails'));

export const RenderWithOrWithoutHeader = (props) => {
  const Tag = props.tag;
  const Component = props.component;
  return (
    <>
      <Suspense fallback={<LoadingPlaceholder />}>
        {props.withHeader && <QueueHeader />}
        <Tag role="main" className="site__content">
          <Component {...props} />
        </Tag>
      </Suspense>
    </>
  );
};

export class OfficeWrapper extends Component {
  constructor(props) {
    super(props);

    this.state = { hasError: false };
  }

  componentDidMount() {
    document.title = 'Transcom PPP: Office';
    this.props.loadInternalSchema();
    this.props.loadPublicSchema();
    this.props.getCurrentUserInfo();
  }

  componentDidCatch(error, info) {
    this.setState({
      hasError: true,
      error,
      info,
    });
  }

  render() {
    const ConditionalWrap = ({ condition, wrap, children }) => (condition ? wrap(children) : <>{children}</>);
    const { context: { flags: { too, tio } } = { flags: { too: null } } } = this.props;
    const DivOrMainTag = detectIE11() ? 'div' : 'main';
    const { userIsLoggedIn } = this.props;
    return (
      <ConnectedRouter history={history}>
        <div className="Office site">
          <FOUOHeader />
          <Suspense fallback={<LoadingPlaceholder />}>{!userIsLoggedIn && <QueueHeader />}</Suspense>
          <ConditionalWrap
            condition={!userIsLoggedIn}
            wrap={(children) => (
              <DivOrMainTag role="main" className="site__content">
                {children}
              </DivOrMainTag>
            )}
          >
            <ConnectedLogoutOnInactivity />
            {this.state.hasError && <SomethingWentWrong error={this.state.error} info={this.state.info} />}
            {!this.state.hasError && (
              <Switch>
                <PrivateRoute
                  exact
                  path="/"
                  render={(props) => (
                    <Suspense fallback={<LoadingPlaceholder />}>
                      <QueueHeader />
                      <main role="main" className="site__content">
                        <ConnectedOfficeHome {...props} />
                      </main>
                    </Suspense>
                  )}
                />
                <PrivateRoute exact path="/select-application" component={ConnectedSelectApplication} />
                <PrivateRoute
                  path="/queues/:queueType/moves/:moveId"
                  component={(props) => (
                    <Suspense fallback={<LoadingPlaceholder />}>
                      <RenderWithOrWithoutHeader component={MoveInfo} withHeader={true} tag={DivOrMainTag} {...props} />
                    </Suspense>
                  )}
                  requiredRoles={[roleTypes.PPM]}
                />
                <PrivateRoute
                  path="/queues/:queueType"
                  component={(props) => (
                    <Suspense fallback={<LoadingPlaceholder />}>
                      <RenderWithOrWithoutHeader component={Queues} withHeader={true} tag={DivOrMainTag} {...props} />
                    </Suspense>
                  )}
                  requiredRoles={[roleTypes.PPM]}
                />
                <PrivateRoute
                  path="/moves/:moveId/orders"
                  component={(props) => (
                    <Suspense fallback={<LoadingPlaceholder />}>
                      <RenderWithOrWithoutHeader
                        component={OrdersInfo}
                        withHeader={false}
                        tag={DivOrMainTag}
                        {...props}
                      />
                    </Suspense>
                  )}
                  requiredRoles={[roleTypes.PPM]}
                />
                <PrivateRoute
                  path="/moves/:moveId/documents/:moveDocumentId?"
                  component={(props) => (
                    <Suspense fallback={<LoadingPlaceholder />}>
                      <RenderWithOrWithoutHeader
                        component={DocumentViewer}
                        withHeader={false}
                        tag={DivOrMainTag}
                        {...props}
                      />
                    </Suspense>
                  )}
                  requiredRoles={[roleTypes.PPM]}
                />
                {!isProduction && (
                  <PrivateRoute
                    path="/playground"
                    component={(props) => (
                      <Suspense fallback={<LoadingPlaceholder />}>
                        <RenderWithOrWithoutHeader
                          component={ScratchPad}
                          withHeader={true}
                          tag={DivOrMainTag}
                          {...props}
                        />
                      </Suspense>
                    )}
                  />
                )}
                <Suspense fallback={<LoadingPlaceholder />}>
                  <Switch>
                    {too && (
                      <PrivateRoute
                        path="/moves/queue"
                        exact
                        component={(props) => (
                          <Suspense fallback={<LoadingPlaceholder />}>
                            <RenderWithOrWithoutHeader
                              component={TOO}
                              withHeader={true}
                              tag={DivOrMainTag}
                              {...props}
                            />
                          </Suspense>
                        )}
                        requiredRoles={[roleTypes.TOO]}
                      />
                    )}
                    {too && (
                      <PrivateRoute
                        path="/move/mto/:moveTaskOrderId"
                        exact
                        component={(props) => (
                          <Suspense fallback={<LoadingPlaceholder />}>
                            <RenderWithOrWithoutHeader
                              component={TOOMoveTaskOrder}
                              withHeader={true}
                              tag={DivOrMainTag}
                              {...props}
                            />
                          </Suspense>
                        )}
                        requiredRoles={[roleTypes.TOO]}
                      />
                    )}
                    {too && (
                      <PrivateRoute
                        path="/moves/:locator"
                        exact
                        component={(props) => (
                          <Suspense fallback={<LoadingPlaceholder />}>
                            <RenderWithOrWithoutHeader
                              component={MoveDetails}
                              withHeader={true}
                              tag={DivOrMainTag}
                              {...props}
                            />
                          </Suspense>
                        )}
                        requiredRoles={[roleTypes.TOO]}
                      />
                    )}
                    {too && (
                      <PrivateRoute
                        path="/moves/:moveOrderId/customer/:customerId"
                        component={(props) => (
                          <Suspense fallback={<LoadingPlaceholder />}>
                            <RenderWithOrWithoutHeader
                              component={CustomerDetails}
                              withHeader={true}
                              tag={DivOrMainTag}
                              {...props}
                            />
                          </Suspense>
                        )}
                        requiredRoles={[roleTypes.TOO]}
                      />
                    )}
                    {too && <Route path="/verification-in-progress" component={TOOVerificationInProgress} />}
                    {tio && (
                      <PrivateRoute
                        path="/invoicing/queue"
                        component={(props) => (
                          <Suspense fallback={<LoadingPlaceholder />}>
                            <RenderWithOrWithoutHeader
                              component={TIO}
                              withHeader={true}
                              tag={DivOrMainTag}
                              {...props}
                            />
                          </Suspense>
                        )}
                        requiredRoles={[roleTypes.TIO]}
                      />
                    )}
                    {tio && (
                      <PrivateRoute
                        path="/payment_requests/:id"
                        component={(props) => (
                          <Suspense fallback={<LoadingPlaceholder />}>
                            <RenderWithOrWithoutHeader
                              component={PaymentRequestShow}
                              withHeader={true}
                              tag={DivOrMainTag}
                              {...props}
                            />
                          </Suspense>
                        )}
                        requiredRoles={[roleTypes.TIO]}
                      />
                    )}
                    {tio && (
                      <PrivateRoute
                        path="/payment_requests"
                        component={(props) => (
                          <Suspense fallback={<LoadingPlaceholder />}>
                            <RenderWithOrWithoutHeader
                              component={PaymentRequestIndex}
                              withHeader={true}
                              tag={DivOrMainTag}
                              {...props}
                            />
                          </Suspense>
                        )}
                        requiredRoles={[roleTypes.TIO]}
                      />
                    )}
                  </Switch>
                </Suspense>
              </Switch>
            )}
          </ConditionalWrap>
        </div>
      </ConnectedRouter>
    );
  }
}

OfficeWrapper.defaultProps = {
  loadInternalSchema: () => {},
  loadPublicSchema: () => {},
};

const mapStateToProps = (state) => {
  const user = selectCurrentUser(state);
  return {
    swaggerError: state.swaggerInternal.hasErrored,
    userIsLoggedIn: user.isLoggedIn,
  };
};

const mapDispatchToProps = (dispatch) =>
  bindActionCreators({ loadInternalSchema, loadPublicSchema, getCurrentUserInfo }, dispatch);

export default withContext(connect(mapStateToProps, mapDispatchToProps)(OfficeWrapper));