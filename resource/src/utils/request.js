/* eslint-disable */
import axios from 'axios'

import { ElMessage } from 'element-plus';

import.meta.VITE_APP_BASE_API

// create an axios service
const service = axios.create({
  
  baseURL: import.meta.env.VITE_APP_BASE_API,
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 100000, // request timeout
  withCredentials: true,
})


const requestError = error => {
  ElMessage({
    message: error,
    type: 'error',
    duration: 5 * 1000,
    offset: 120
  })
  return Promise.reject(error)
}

const beforeRequest = config => {
  config.headers['Content-Type'] = 'application/json';
  return config
}

const beforeResponse = res => {
    return res.data
}



service.interceptors.request.use(beforeRequest, requestError)
service.interceptors.response.use(beforeResponse, requestError)


export default service
