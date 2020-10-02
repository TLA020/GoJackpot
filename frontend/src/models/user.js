/**
 * @property {string} email
 * @property {string} username
 * @property {string} avatar
 * @property {string} token
 */
export default class User {
  constructor({ email, username, avatar, token }) {
    /** @type {string} */
    this.email = email;
    /** @type {string} */
    this.username = username;
    /** @type {string} */
    this.avatar = avatar;
    /** @type {string} */
    this.token = token;
  }
}
