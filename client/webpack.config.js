
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
		host: "127.0.0.1",
		port: 5005,
		hot: true,
		static:{
			directory: path.join(__dirname, 'public')
		},
		allowedHosts: [
			'developer-libery-labs.com'
		],
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
		AUTH_SERVER: process.env.AUTH_SERVER,
		USERS_SERVER: process.env.USERS_SERVER,
		PAYMENTS_SERVER: process.env.PAYMENTS_SERVER,
		JD_SERVER: process.env.JD_SERVER,
		WHATSAPP_REDIRECTION: process.env.WHATSAPP_REDIRECTION,
	}

	config.plugins.push(
		new webpack.DefinePlugin({
			"AUTH_SERVER": JSON.stringify(build_config.AUTH_SERVER),
			"USERS_SERVER": JSON.stringify(build_config.USERS_SERVER),
			"PAYMENTS_SERVER": JSON.stringify(build_config.PAYMENTS_SERVER),
			"JD_SERVER": JSON.stringify(build_config.JD_SERVER),
			"WHATSAPP_REDIRECTION": JSON.stringify(build_config.WHATSAPP_REDIRECTION),
		})
	);

	return config
} 