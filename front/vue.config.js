'use strict'
let webpack = require('webpack');
let JavaScriptObfuscator = require('webpack-obfuscator');

module.exports = {
  devServer: {
    hot: true,
    port: 8087, // CHANGE YOUR PORT HERE!
    https: false,
    inline: true
  },
  outputDir: process.env.NODE_ENV === 'stage' ? 'stage' : 'dist',
  pluginOptions: {
    quasar: {
      theme: 'mat'
    }
  },
  configureWebpack: {
    devtool: !(process.env.NODE_ENV === 'prod' || process.env.NODE_ENV === 'stage') ? 'cheap-module-eval-source-map' : false,
    watch: !(process.env.NODE_ENV === 'prod' || process.env.NODE_ENV === 'stage'),
    plugins: (process.env.NODE_ENV === 'prod' || process.env.NODE_ENV === 'stage') ? [
      new JavaScriptObfuscator ({
        rotateUnicodeArray: true
      }, ['node_modules/*'])
    ] : []
  }
}