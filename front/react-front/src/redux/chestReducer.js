
const ADD_CHESTS = "chest/ADD_CHESTS"
const OPEN_CHEST = "chest/OPEN_CHEST"

let initialState = {
  chests: {}
}

const chestReducer = (state = initialState, action) => {
  switch (action.type) {
    case ADD_CHESTS:
      return {
        ...state,
        chests: action.chests
      }
    default:
      return state;
  }
}


export const addChests = (chests) => ({ type: ADD_CHESTS, chests } )
export const openChest = () => ({ type: OPEN_CHEST } )



export default chestReducer;