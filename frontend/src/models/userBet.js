import Bet from "./bet";
import Player from "./player";

/**
 * @property {number} share
 * @property {number} startTicket
 * @property {number} endTicket
 * @property {Bet[]} bets
 * @property {Player} player
 */
export default class UserBet {
  constructor({ share, startTicket, endTicket, bets, player }) {
    /** @type {number} */
    this.share = share;
    /** @type {number} */
    this.startTicket = startTicket;
    /** @type {number} */
    this.endTicket = endTicket;
    /** @type {Bet[]} */
    this.bets = bets.map(b => new Bet(b));
    /** @type {Player} */
    this.player = new Player(player);
  }
}
