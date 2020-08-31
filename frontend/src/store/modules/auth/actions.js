import axios from "axios";
import router from "@/router";
const API_URL = (process.env.VUE_APP_API_BASE_URL || "/api/v1") + "/accounts";

export default {
  login({commit}, { email, password }) {
    commit("SET_LOADING", true);
    axios.post(`${API_URL}/login`, {email, password})
        .then(({ data }) => {
          commit("SET_USER", data);
          commit("SET_ERROR", null);
          router.push("/");
        })
        .catch(err => {
          commit("SET_ERROR", "Invalid credentials, try again");
          console.log(err)
        })
        .finally(() => {
          commit("SET_LOADING", false);
        })
  },

  register({commit}, { email, password }) {
    commit("SET_LOADING", true);
      axios.post(`${API_URL}/register`, {email, password})
        .then(({data}) => {
          commit("SET_USER", data);
          commit("SET_ERROR", null);
          router.push("/")
        })
        .catch(err => {
          commit("SET_ERROR", "email already in use, consider another");
          console.log(err)
        }).finally(() => {
        commit("SET_LOADING", false);
      })
  }
};


