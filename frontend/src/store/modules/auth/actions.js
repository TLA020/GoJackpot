import axios from "axios";
import router from "@/router";
const API_URL = (process.env.VUE_APP_API_BASE_URL || "/api/v1") + "/accounts";

export default {
  login({commit}, { email, password }) {
      axios.post(`${API_URL}/login`, {email, password})
        .then(({ data }) => {
          commit("SET_USER", data)
          router.push("/");
        })
        .catch(err => {
          console.log(err)
        })
  },

  register({commit}, { email, password }) {
      axios.post(`${API_URL}/register`, {email, password})
        .then(({data}) => {
          commit("SET_USER", data)
          router.push("/")
        })
        .catch(err => {
          console.log(err)
        })
  }
};


