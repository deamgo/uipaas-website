import axios from "axios"


const requests = axios.create({
  baseURL: "/api",
  timeout: 5000,
});

requests.interceptors.request.use(
  (config) => {
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

requests.interceptors.response.use(
  (response) => {
    return response.data;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default requests