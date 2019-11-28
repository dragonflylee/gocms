'use strict';

const path = require('path');
const webpack = require('webpack');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const formatMessages = require('webpack-format-messages');

let compiler = webpack({
  mode: 'production',
  entry: {
    admin: path.resolve(__dirname, 'admin.js'),
    login: path.resolve(__dirname, 'login.js'),
    index: path.resolve(__dirname, 'index.js')
  },
  output: {
    path: path.resolve('static'),
    filename: 'js/[name].js'
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: 'css/[name].css'
    }),
    new webpack.ProvidePlugin({
      $: "jquery",
      jQuery: "jquery",
      "window.jQuery": "jquery",
    })
  ],
  module: {
    rules: [
      {
        test: /\.(scss|sass|css)$/,
        use: [MiniCssExtractPlugin.loader, "css-loader"],
      },
      {
        test: /\.(png|gif|eot|svg|ttf|woff|woff2)\w*/,
        use: [{
          loader: 'file-loader',
          options: {
            name: 'img/[hash:hex:8].[ext]',
            publicPath: '..'
          }
        }]
      }
    ]
  },
  optimization: {
    splitChunks: {
      chunks: 'async',
      minChunks: 1 
    }
  },
  performance: {
    hints: 'warning',
    //入口起点的最大体积 整数类型（以字节为单位）
    maxEntrypointSize: 50000000,
    //生成文件的最大体积 整数类型（以字节为单位 300k）
    maxAssetSize: 30000000,
  }
});

compiler.hooks.invalid.tap('invalid', function () {
  console.log('Compiling...');
});

compiler.hooks.done.tap('done', (stats) => {
  const messages = formatMessages(stats);

  if (!messages.errors.length && !messages.warnings.length) {
    console.log('Compiled successfully!');
  }

  if (messages.errors.length) {
    console.log('Failed to compile.');
    messages.errors.forEach(e => console.log(e));
    return;
  }

  if (messages.warnings.length) {
    console.log('Compiled with warnings.');
    messages.warnings.forEach(w => console.log(w));
  }
});

compiler.run((err, stats) => {
  if (err) {
    console.log(err);
  }
})
