import svelte from 'rollup-plugin-svelte';
import commonjs from '@rollup/plugin-commonjs';
import resolve from '@rollup/plugin-node-resolve';
import livereload from 'rollup-plugin-livereload';
import { terser } from 'rollup-plugin-terser';
import css from 'rollup-plugin-css-only';

const production = !process.env.ROLLUP_WATCH;
const staticFolder = "./static";

let server;
function toExit() { if (server) server.kill(0); }
function serve() {
	return {
		writeBundle() {
			if (server) return;
			server = require('child_process').spawn('npm', ['run', 'start', '--', '-D', '-q'], {
				stdio: ['ignore', 'inherit', 'inherit'],
				shell: true
			});
			process.on('SIGTERM', toExit);
			process.on('exit', toExit);
		}
	};
}

export default {
	input: './lib/boot.js',
	output: {
		sourcemap: true,
		format: 'iife',
		name: 'app',
		file: staticFolder + '/app.js'
	},
	plugins: [
		svelte({ compilerOptions: { dev: !production } }), // enable run-time checks when not in production
		css({ output: 'app.css' }),
		resolve({
			browser: true,
			dedupe: ['svelte']
		}),
		commonjs(),
		!production && serve(), // In dev mode, call `npm run start` once the bundle has been generated
		!production && livereload(".."), // Watch directory and refresh browser
		production && terser() // we're building for production (npm run build instead of npm run dev), minify
	],
	watch: {
		clearScreen: false
	}
};
