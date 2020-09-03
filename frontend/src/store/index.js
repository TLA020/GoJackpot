import Vue from "vue";
import Vuex from "vuex";

import authStore from "./modules/auth";
import gameStore from "./modules/game";
import socketHandler from "./socket.handler";

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
      socket: {
        isConnected: false,
        message: {},
        reconnectError: false
      }
    },
    mutations: {
      SOCKET_ONOPEN(state, event) {
        Vue.$socket = event.currentTarget;

        if (state.$auth.user.token) {
          Vue.$socket.sendObj({ event: "auth", data: state.$auth.user })
        }
        state.socket.isConnected = true;
      },
      SOCKET_ONCLOSE(state) {
        state.socket.isConnected = false;
      },
      SOCKET_ONERROR(state, event) {
        console.error(state, event);
      },
      SOCKET_ONMESSAGE(state, message) {
        state.socket.message = message;
      },
      SOCKET_RECONNECT(state, count) {
        console.info(state, count);
      },
      SOCKET_RECONNECT_ERROR(state) {
        state.socket.reconnectError = true;
      }
    },
    actions: {
      sendSocket: function (context, message) {
        Vue.$socket.sendObj(message);
      }
    },
    modules: {
      $auth: authStore,
      $game: gameStore,
    },
    plugins: [socketHandler]
});