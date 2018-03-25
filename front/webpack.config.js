const path = require('path');
const ExtractTextPlugin = require("extract-text-webpack-plugin");

module.exports = [
  {
    entry: path.resolve(__dirname, 'src/app.js'),
    output: {
      path: path.resolve(__dirname, 'dist'),
      filename: 'app.js'
    },
    module: {
      rules: [
        {
          test: /\.(js|jsx|mjs)$/,
          exclude: /node_modules/,
          loader: 'babel-loader',
          options: {
            presets: [
              'react',
              'env', // es2015からenvになった
            ],
          },
        },
        {
        },
      ],
    },
    plugins: [
      new ExtractTextPlugin('app.css'),
    ],
    devServer: {
      contentBase: path.resolve(__dirname, 'dist'),
      port: 8080
    },
    devtool: 'inline-source-map'
  },
  {
    entry: {
      style: './src/css/index.scss',
    },
    output: {
      path: path.resolve(__dirname, 'dist'),
      filename: 'app.css',
    },
    module: {
      rules: [
        {
          test: /\.scss$/,
          use: ExtractTextPlugin.extract({
            use: [
              {
                loader: 'css-loader',
                options: {
                  url: false,
                  sourceMap: true
                },
              },
              {
                loader: 'sass-loader',
                options: {
                  sourceMap: true
                }
              }
            ],
          }),
        },
      ],
    },
    plugins: [
      new ExtractTextPlugin('app.css'),
    ],
  },
];
