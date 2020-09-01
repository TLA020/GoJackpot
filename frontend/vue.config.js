const path = require("path");

module.exports = {
  "assetsDir": "static",
  "transpileDependencies": [
    "vuetify"
  ],
  devServer: {
    host: 'localhost',
    port: 8080, // CHANGE YOUR PORT HERE!
    https: true,
    hotOnly: false,
  },
  configureWebpack: {
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "src/")
      }
    }
  }
}
