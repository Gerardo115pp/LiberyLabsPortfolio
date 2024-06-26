import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import path from 'path';
import fs from 'fs';

export default defineConfig(async ({ command, mode, isSsrBuild, isPreview }) => {
	let is_production = command === 'build';

	let build_config = {
		PORTFOLIO_SERVICE: JSON.stringify(process.env.PORTFOLIO_SERVICE),
		RECAPTCHA_PK: JSON.stringify(process.env.RECAPTCHA_PK),
	}

	/**
	 *  @type {import('vite').UserConfig}
	 */
	let config = {
		server: {
			open: false,
			host: "0.0.0.0",
			port: 13001,
		},
		define: {
			...build_config
		},
		resolve: {
			alias: {
				'@libs': path.resolve(__dirname, 'src/libs'),
				'@components': path.resolve(__dirname, 'src/components'),
				'@pages': path.resolve(__dirname, 'src/pages'),
				'@svg': path.resolve(__dirname, 'src/svg'),
				'@models': path.resolve(__dirname, 'src/models'),
				"@actions": path.resolve(__dirname, 'src/actions'),
				"@events": path.resolve(__dirname, 'src/events'),
				"@stores": path.resolve(__dirname, 'src/stores'),
				"@app": path.resolve(__dirname, 'src'),
			},
		},
		plugins: [
			sveltekit()
		],
		clearScreen: true
	};

	if (!is_production) {
		config.server.https = {
			key: fs.readFileSync(process.env.SSL_KEY_PATH),
			cert: fs.readFileSync(process.env.SSL_CERT_PATH),
		}
	}

	return config;
})
