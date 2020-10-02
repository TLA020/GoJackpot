/**
 * @property {number} id
 * @property {string} username
 * @property {string} email
 * @property {string} avatar
 */

export default class Player {
  constructor({ id, username, avatar, email }) {
    /** @type {number} */
    this.id = id;
    /** @type {string} */
    this.username = username;
    /** @type {string} */
    this.email = email;
    /** @type {string} */
    this.avatar = avatar;
  }
}
