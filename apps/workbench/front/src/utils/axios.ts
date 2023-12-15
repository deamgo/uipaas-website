import axios from "axios"

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
