import Vue from "vue";

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
      const { event, data } = mutation.payload;
      if (socketEvents[event]) {
        socketEvents[event].forEach(fn => fn(data));
      }

      switch (event) {
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
        case "winner-picked":
          store.commit("$game/SET_WINNER", data);
          store.commit("$game/SET_GAME", data.game);
          break;
        default:
          return;
      }
    }
  });
};
