export default {
  Authenticated: state => {
    return state.authenticated;
  },
  userAvatar: state => {
    if (process.env.VUE_APP_DEVELOPMENT === "1") {
      return "http://localhost:5001/" + state.user.avatar;
    }
    return state.user.avatar;
  }
};
