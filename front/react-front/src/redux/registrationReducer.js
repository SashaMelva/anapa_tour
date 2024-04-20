
const AUTH_IN = "reg/AUTH_IN"
const AUTH_OUT = "reg/AUTH_OUT"

let initialState = {
  isAuth: false
}

const registrationReducer = (state = initialState, action) => {
  switch (action.type) {
    case AUTH_IN:
      return {
        ...state,
        isAuth: true,
      }
    case AUTH_OUT:
      return {
        ...state,
        isAuth: false
      }
    default:
      return state;
  }
}


export const authIn = () => ({ type: AUTH_IN } )
export const  authOut = () => ({ type: AUTH_OUT } )



export default registrationReducer;