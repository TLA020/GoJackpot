import Vue from "vue";
import router from "@/router";

const socketEvents = {};

Vue.prototype.$socketEvent = (event, fn) => {
  if (!socketEvents[event]) {
    socketEvents[event] = [];
  }
  socketEvents[event].push(fn);
};

export default store => {
  store.subscribe(mutation => {
    if (mutation.type === "SOCKET_ONMESSAGE") {
      const { type, data } = mutation.payload;
      if (socketEvents[type]) {
        socketEvents[type].forEach(fn => fn(data));
      }

      switch (type) {
        case "current-users":
          store.commit("$game/SET_CURRENT_USERS", data.users);
          break;
        case "current-game":
          store.commit("$game/SET_GAME", data.game);
          break;
        case "new-game":
          store.commit("$game/SET_TIME_LEFT", null);
          store.commit("$game/SET_GAME", data.game);
          break;
        case "start-game":
          store.commit("$game/SET_GAME", data.game);
          break;
        case "bet-placed":
          store.commit("$game/SET_GAME", data.game);
          break;
        case "shares-updated":
          store.commit("$game/SET_GAME", data.game);
          break;
        case "time-left":
          store.commit("$game/SET_TIME_LEFT", data.timeLeft);
          break;
        case "invalid-user":
          localStorage.setItem("user", JSON.stringify(data));
          store.commit("$auth/SET_USER", null);
          router.push("/auth");
          break;
        default:
          return;
      }
    }
  });
};
