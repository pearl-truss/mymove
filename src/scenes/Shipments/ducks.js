import { ShipmentsIndex } from 'shared/api';

// Types
export const SHOW_SHIPMENTS = 'SHOW_SHIPMENTS';
export const SHOW_SHIPMENTS_SUCCESS = 'SHOW_SHIPMENTS_SUCCESS';
export const SHOW_SHIPMENTS_FAILURE = 'SHOW_SHIPMENTS_FAILURE';

// Actions
export const createShowShipmentsRequest = () => ({
  type: SHOW_SHIPMENTS,
});

export const createShowShipmentsSuccess = shipments => ({
  type: SHOW_SHIPMENTS_SUCCESS,
  shipments,
});

export const createShowShipmentsFailure = error => ({
  type: SHOW_SHIPMENTS_FAILURE,
  error,
});

// Action Creator
export function loadShipments(shipmentsStatus) {
  // Interpreted by the thunk middleware:
  return function(dispatch, getState) {
    dispatch(createShowShipmentsRequest());
    ShipmentsIndex(shipmentsStatus)
      .then(shipments => dispatch(createShowShipmentsSuccess(shipments)))
      .catch(error => dispatch(createShowShipmentsFailure(error)));
  };
}

// Reducer
export function shipmentsReducer(
  state = { shipments: null, hasError: false },
  action,
) {
  switch (action.type) {
    case SHOW_SHIPMENTS_SUCCESS:
      return { shipments: action.shipments, hasError: false };
    case SHOW_SHIPMENTS_FAILURE:
      return { shipments: null, hasError: true };
    default:
      return state;
  }
}
