import UserBet from "./userBet";
/**
 * @property {number} id
 * @property {number} state
 * @property {UserBet[]} userBets
 */
export default class Game {
  constructor({ id, state, userBets }) {
    /** @type {number} */
    this.id = id;
    /** @type {string} */
    this.state = state;
    /** @type {UserBet[]} */
    this.userBets = userBets.map(u => new UserBet(u));
  }
}
