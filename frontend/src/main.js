import Vue from 'vue'
import Vuex from 'vuex'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import App from './App.vue'
import router from "./router";
import store from "./store";
import websocketPlugin from "@/setup/websockets";

Vue.config.productionTip = false
Vue.use(Vuex)
Vue.use(websocketPlugin)

Vue.use(Vuetify)

new Vue({
  router,
  store,
  vuetify : new Vuetify({ theme: {
      dark: false,
    },}),
  render: h => h(App)
}).$mount("#app");