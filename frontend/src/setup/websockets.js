import VueNativeSock from "vue-native-websocket";
import store from "../store";

export default {
  install: Vue => {
    let socketUrl = "";

    if (!location.hostname.includes("localhost")) {
       socketUrl = process.env.VUE_APP_SOCKET_URL || `wss://${location.hostname}:${location.port}/ws`;
    } else {
       socketUrl = process.env.VUE_APP_SOCKET_URL || `ws://${location.hostname}:5001/ws`;
    }
    Vue.use(VueNativeSock, `${socketUrl}`, {
      store: store,
      format: "json"
    });
  }
};
