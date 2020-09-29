import axios from "axios";
import router from "@/router";
const API_URL = (process.env.VUE_APP_API_BASE_URL || "/api/v1") + "/accounts";

export default {
  login({ commit, dispatch }, { email, password, username }) {
    commit("SET_LOADING", true);
    axios
      .post(`${API_URL}/login`, { email, password, username })
      .then(({ data }) => {
        commit("SET_USER", data);
        commit("SET_ERROR", null);
        dispatch("sendSocket", { type: "auth", data }, { root: true });
        localStorage.setItem("user", JSON.stringify(data));
        router.push("/");
      })
      .catch(err => {
        commit("SET_ERROR", "Invalid credentials, try again");
        console.log(err);
      })
      .finally(() => {
        commit("SET_LOADING", false);
      });
  },

  register({ commit, dispatch }, { email, password, username }) {
    commit("SET_LOADING", true);
    axios
      .post(`${API_URL}/register`, { email, password, username })
      .then(({ data }) => {
        commit("SET_USER", data);
        commit("SET_ERROR", null);
        dispatch("sendSocket", { type: "auth", data }, { root: true });
        localStorage.setItem("user", JSON.stringify(data));
        router.push("/");
      })
      .catch(err => {
        commit("SET_ERROR", "email already in use, consider another");
        console.log(err);
      })
      .finally(() => {
        commit("SET_LOADING", false);
      });
  }
};
