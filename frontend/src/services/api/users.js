import axios from "axios";
const BASE_URL = process.env.VUE_APP_API_BASE_URL;

export default {
  create: () => {
    return axios.post(`${BASE_URL}/users/`).then(({ data: user }) => {
      return user;
    });
  },
};