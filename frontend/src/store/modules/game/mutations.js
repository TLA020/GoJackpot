import Game from "../../../models/game";
export default {
  SET_CURRENT_USERS(state, currentUsers) {
    state.currentUsers = currentUsers;
  },
  SET_GAME(state, game) {
    state.game = new Game(game);
    console.log(state.game);
  },
  SET_TIME_LEFT(state, time) {
    state.timeLeft = time;
  }
};
