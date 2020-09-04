export default {
  SET_CURRENT_USERS(state, currentUsers) {
    state.currentUsers = currentUsers;
  },
  SET_GAME(state, game) {
    state.game = game;
  },
  SET_TIME_LEFT(state, time) {
    state.timeLeft = time;
  },
  SET_WINNER(state, data) {
    state.currentWinner.user = data.winner;
    state.currentWinner.amount = data.amount;
  }
};
