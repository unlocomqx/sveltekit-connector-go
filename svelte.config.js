import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	compilerOptions: {
		experimental: {
			async: true
		}
	},
	kit: {
		adapter: adapter({ strict: false }),
		experimental: {
			remoteFunctions: true
		}
	}
};

export default config;
