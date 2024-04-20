import axios from 'axios';

export const instanse = axios.create({
  withCredentials: true,
  baseURL: '',
  header: {
    "API-KEY": ""
  }
})