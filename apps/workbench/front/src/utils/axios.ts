import axios from "axios"


<<<<<<<< <Temporary merge branch 1
const requests = axios.create({
  baseURL: "/api",
  timeout: 5000,
});

requests.interceptors.request.use(
=========
const request = axios.create({
  baseURL: "/api/v1",
  timeout: 5000,
});

interface IResp {
  value: {
    code: number
    msg: string
    data: {
      Token: string
    } | {
      CodeKey: string
    } | string | null
  }
}

request.interceptors.request.use(
  (config) => {
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

request.interceptors.response.use(
  (response) => {
    return response.data;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default request