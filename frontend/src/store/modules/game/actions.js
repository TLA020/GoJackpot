import axios from "axios";
const API_URL = (process.env.VUE_APP_API_BASE_URL || "/api/v1") + "/game";

export default {
  fakeBets() {
    axios
      .post(`${API_URL}/test/bets/random`)
      .then(() => {})
      .catch(err => {
        console.log(err);
      });
  }
};
