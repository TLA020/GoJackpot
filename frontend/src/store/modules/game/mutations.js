export default {
  SET_CURRENT_USERS(state, currentUsers) {
    state.currentUsers = currentUsers;
  },
  SET_GAME(state, game) {
    state.game = Object.assign({}, game);
  },
  SET_TIME_LEFT(state, time) {
    state.timeLeft = time;
  }
};
