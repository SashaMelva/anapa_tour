import { instanse } from "./api";

export const authApi = {
  reg(obj) {
    console.log(obj)
    // return instanse.put('/registration', {...obj})
    //   .then(res => {return 'success'})
  },
  isReg(obj) {
    console.log(obj)
    // return instanse.put('/registration', {...obj})
    //   .then(res => {return true})
  },
  user(id) {
    console.log(id)
    // return instanse.get(`/user/${id}`)
    //   .then(res => {return {data: [1, 2, 3]}})
  },
  login(obj) {
    console.log(obj)
    // return instanse.post('/login', {...obj})
    //   .then(res => {return 'success'})
  }
}