
const path = require('path');
const htmlWebpackPlugin = require('html-webpack-plugin');
const webpack  = require('webpack');

const config = {
	entry: './src/index.js',
	output: {
		path: path.resolve(__dirname, 'build'),
		filename: 'boundle.js'
	},
	devServer: {
		host: "192.168.0.140",
		port: 5005,
		hot: true,
		static:{
			directory: path.join(__dirname, 'public')
		},
		// allowedHosts: [
		// 	'developer-libery-labs.com'
		// ],
		hot: true, // this
		historyApiFallback: true
	},
	resolve: {
		alias: {
			svelte: path.resolve('node_modules', 'svelte'),
			'@libs': path.resolve(__dirname, 'src/libs'),
			'@components': path.resolve(__dirname, 'src/components'),
			'@pages': path.resolve(__dirname, 'src/pages'),
			'@svg': path.resolve(__dirname, 'src/svg'),
			'@models': path.resolve(__dirname, 'src/models'),
			"@actions": path.resolve(__dirname, 'src/actions'),
			"@events": path.resolve(__dirname, 'src/events'),
			"@stores": path.resolve(__dirname, 'src/stores'),
		},
		extensions: ['*', '.mjs', '.js', '.svelte'],
		mainFields: ['svelte', 'browser', 'module', 'main'],
	},
	module: {
		rules: [
			{
				test:  /\.js?$/,
				exclude: /node_modules/,
				use: {
					loader: 'babel-loader'
				}
			},
			{
				test: /\.svelte$/,
				use: {
					loader: 'svelte-loader',
					options: {
						onwarn: (warning, handler) => {
							if (warning.code.startsWith("a11y") || warning.code.startsWith("css-unused")) return;

							if (warning.code === "unused-export-let") return;
							console.log(`stupid warning: ${warning.code}`);
		
							handler(warning);
						}
					}
				}
				
			},
			{
				test: /\.svg$/,
				exclude: /node_modules/,
				use: {
					loader: 'svg-inline-loader',
					options: {
					  removeSVGTagAttrs: true
					}
				}
			}
		]
	},
	plugins: [
		new htmlWebpackPlugin({
			inject: true,
			template: './public/index.html',
			filename: './index.html'
		})
	]
}


module.exports = (env, argv) => {
	const isProd = argv.mode === 'production';
	const build_config = {
		PORTFOLIO_SERVICE: process.env.PORTFOLIO_SERVICE,
	}

	config.plugins.push(
		new webpack.DefinePlugin({
			"PORTFOLIO_SERVICE": JSON.stringify(build_config.PORTFOLIO_SERVICE)
		})
	);

	return config
} 