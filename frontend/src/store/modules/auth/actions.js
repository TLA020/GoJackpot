import apiService from "@/services/api";

export default {
  login({ commit }) {
    apiService.users
      .create()
      .then(user => commit("SET_USER", user))
      .catch(e => {
        console.error(e);
      });
  },
};
