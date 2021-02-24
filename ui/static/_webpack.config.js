const webpack = require("webpack");

module.exports = (env) => {
    return require(`./webpack.config.${env}.js`)
}