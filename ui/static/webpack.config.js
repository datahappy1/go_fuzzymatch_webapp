// webpack.config.js
const webpack = require("webpack");

// definePlugin takes raw strings and inserts them, so you can put strings of JS if you want.
var definePlugin = new webpack.DefinePlugin({
  __DEV__: JSON.stringify(JSON.parse(process.env.BUILD_DEV || 'true')),
  __PRERELEASE__: JSON.stringify(JSON.parse(process.env.BUILD_PRERELEASE || 'false'))
});

//https://github.com/petehunt/webpack-howto#6-feature-flags

module.exports = {
  entry: './src/index.js',
  output: {
    filename: './bundle.js'       
  },
  plugins: [definePlugin]
};