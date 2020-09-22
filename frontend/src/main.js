import Vue from "vue";
import Vuetify from "vuetify";
import "vuetify/dist/vuetify.min.css";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import websocketPlugin from "@/setup/websockets";

Vue.config.productionTip = false;
Vue.use(websocketPlugin);

Vue.use(Vuetify);

Vue.prototype.$store = store;

new Vue({
  render: h => h(App),
  store,
  router,
  vuetify: new Vuetify({
    theme: {
      dark: true
    }
  })
}).$mount("#app");
