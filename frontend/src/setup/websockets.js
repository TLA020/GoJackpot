import VueNativeSock from "vue-native-websocket";
// import store from "@/store";

export default {
  install: Vue => {
    const socketUrl = process.env.VUE_APP_SOCKET_URL || `ws://${location.hostname}:${location.port}/ws`;

    Vue.use(VueNativeSock, `${socketUrl}`, {
      // store: store,
      format: "json"
    });
  }
};
