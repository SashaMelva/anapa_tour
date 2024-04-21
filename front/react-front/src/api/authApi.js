import { instanse } from "./api";

export const authApi = {
  
  reg(obj) {
    return instanse.post('/registration/', {...obj})
      .then(res => {return res.data})
  },
  login(obj) {
    return instanse.post('/login/', {...obj})
      .then(res => {return res.data})
  },
  getChest() {
    return instanse.get(`/pins_vew/`)
      .then(res => {return res.data})
  },
}