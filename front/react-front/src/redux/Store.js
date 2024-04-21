import { applyMiddleware, combineReducers, legacy_createStore as createStore } from "redux";
import thunk from "redux-thunk";
import { compose } from "redux"; 
import toastReducer from "./toastReducer";
import registrationReducer from "./registrationReducer";
import chestReducer from "./chestReducer";



const reducers = combineReducers({
  toast: toastReducer,
  registration: registrationReducer,
  chestObj: chestReducer
});


const store = createStore( 
  reducers                         // Для расширения в браузере redux   
);
window.__store__ = store

export default store