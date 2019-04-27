import { createStore, applyMiddleware } from 'redux';
import { appReducer, tspAppReducer, adminAppReducer } from 'appReducer';
import createHistory from 'history/createBrowserHistory';
import { routerMiddleware } from 'react-router-redux';
import thunk from 'redux-thunk';

import { composeWithDevTools } from 'redux-devtools-extension/logOnlyInProduction';

import { isDevelopment, isTspSite, isAdminSite } from 'shared/constants';
import logger from './reduxLogger';
import * as schema from 'shared/Entities/schema';

export const history = createHistory();

const middlewares = [thunk.withExtraArgument({ schema }), routerMiddleware(history)];

if (isDevelopment && !window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__) {
  middlewares.push(logger);
}

const composeEnhancers = composeWithDevTools({});

export const store = composeEnhancers(applyMiddleware(...middlewares))(createStore)(
  isTspSite ? tspAppReducer : isAdminSite ? adminAppReducer : appReducer,
);

export default store;
