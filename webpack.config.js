const path = require('path');

module.exports = {
  mode: 'development',
  entry: '../hb-server/static/ts/app.ts',                // app.ts should have other imports (.ts or .d.ts) so rest other auto compiled
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: 'ts-loader',
        exclude: path.resolve(__dirname, "node_modules"),
      },
    ],
  },
  resolve: {
    extensions: [ '.tsx', '.ts', '.js' ],
  },
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname+"/static/", 'bundle'),
  },
};