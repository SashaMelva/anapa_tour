import { appStateType, InferActionsTypes } from './Store';
import { ThunkAction } from 'redux-thunk';

const SET_TOAST = "toast/SET_TOAST"
const CLEAR_TOAST = "toast/CLEAR_TOAST"

let initialState = {
  toastObj: null
}

const toastReducer = (state = initialState, action) => {
  switch (action.type) {
    case SET_TOAST:
      return {
        ...state,
        toastObj: action.toastObj
      }
    case CLEAR_TOAST:
      return {
        ...state,
        toastObj: null
      }
    default:
      return state;
  }
}


export const setToastAc = (toastObj) => ({ type: SET_TOAST, toastObj } )
export const clearToastAc = () => ({ type: CLEAR_TOAST } )


export default toastReducer;