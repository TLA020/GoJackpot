import Vue from "vue";
import Vuex from "vuex";

import AuthStore from "./modules/auth";

Vue.use(Vuex);

export default ({
  strict: true,
  modules: {
    auth: AuthStore,
  }
});