import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8090/api/",
  headers: {
    "Content-Type": "application/json",
  },
});

api.interceptors.request.use(function(config) {
  let t = "";
  try {
    let token = JSON.parse(localStorage.getItem("user")).Token;
    t = `Bearer ${token}`;
  } catch (e) {
    t = "";
  }
  config.headers.Authorization = t;
  return config;
});

export default api;
