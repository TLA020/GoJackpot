export default {
  SET_USER(state, payload) {
    state.user = payload;
  },
  REMOVE_USER(state) {
    state.user = null;
  },
  SET_AUTHENTICATED(state, payload) {
    state.authenticated = payload;
  },
  SET_ERROR(state, payload) {
    state.error = payload;
  },
  SET_LOADING(state, payload) {
    state.loading = payload;
  }
}