import fg from 'fast-glob';
import Parser from 'tree-sitter';
import goLang from 'tree-sitter-go';
import path from 'node:path';
import * as fs from 'node:fs';
import { dtsCommandFn, dtsFormFn, dtsQueryFn, imports, queryFn } from './templates.js';

const parser = new Parser();
parser.setLanguage(goLang);

/**
 * @typedef {Object} RemoteFunction
 * @property {string} name - The name of the remote function
 * @property {'query'|'form'|'command'} type - The type of the remote function
 */

/**
 * @typedef {Object} GoKitOptions
 * @property {string} [endpoint] - The Go endpoint URL
 * @property {Object} [headers] - Custom headers to send with requests
 * @property {number} [timeout] - Request timeout in milliseconds
 */

/**
 * Parses Go code to extract remote function declarations and emits corresponding JavaScript code
 * @param {string} code - The Go source code to parse
 * @param {string} file_path - The absolute path to the Go file
 * @param {import('vite').ResolvedConfig} config - The resolved Vite configuration
 * @param {GoKitOptions} options - Configuration options for the Go connector
 */
function transform_code(code, file_path, config, options) {
	const tree = parser.parse(code);

	/** @type {RemoteFunction[]} */
	const remote_functions = [];
	for (const node of tree.rootNode.children) {
		if (node.type === 'function_declaration') {
			// Find the function name node (usually the second child after 'function' keyword)
			const function_name_node = node.children.find((child) => child.type === 'identifier');
			const name = function_name_node ? function_name_node.text : 'unknown';
			if (name) {
				let type = '';
				if (name.startsWith('query')) type = 'query';
				else if (name.startsWith('form')) type = 'form';
				else if (name.startsWith('command')) type = 'command';

				if (type) {
					remote_functions.push({
						name,
						type
					});
				}
			}
		}
	}

	const {code: js_code} = emit_remote_functions({ config, file_path, remote_functions, options });
	return js_code;
}

/**
 * Emits or processes remote functions found in a Go file
 * @param {Object} params - The parameters object
 * @param {import('vite').ResolvedConfig} params.config - The resolved Vite configuration
 * @param {string} params.file_path - The absolute path to the Go file containing remote functions
 * @param {RemoteFunction[]} params.remote_functions - Array of remote function objects
 * @param {GoKitOptions} params.options - Configuration options for the Go connector
 */
function emit_remote_functions({ config, file_path, remote_functions, options }) {
	console.log(`Remote functions in ${file_path}:`, remote_functions);
	const js_path = path.join(
		config.build.outDir,
		path.relative(config.root, file_path).replace(/\\/g, '/').replace(/\.go$/, '.js')
	);
	const dts_path = path.join(
		config.build.outDir,
		path.relative(config.root, file_path).replace(/\\/g, '/').replace(/\.go$/, '.d.ts')
	);
	fs.mkdirSync(path.dirname(dts_path), { recursive: true });

	const js_code = remote_functions.map((remote_fn) => {
		const relative_path = path.relative(config.root, file_path);
		if (remote_fn.type === 'query') {
			return queryFn
				.replace('_name_', remote_fn.name)
				.replace('_endpoint_', options.endpoint.replace(/\/$/, '') || '')
				.replace('_path_', relative_path);
		}
		return '';
	});

	const dts_code = remote_functions.map((remote_fn) => {
		if (remote_fn.type === 'query') {
			return dtsQueryFn.replace('_name_', remote_fn.name);
		} else if (remote_fn.type === 'form') {
			return dtsFormFn.replace('_name_', remote_fn.name);
		} else if (remote_fn.type === 'command') {
			return dtsCommandFn.replace('_name_', remote_fn.name);
		}
		return '';
	});

	const imports_list = new Set(remote_functions.map((fn) => fn.type));

	const imports_code = imports.replace('_imports_', Array.from(imports_list).join(', '));
	const js_code_with_imports = [imports_code, ...js_code].join('\n');
	fs.writeFileSync(js_path, js_code_with_imports);

	const dts_module = `declare module '${file_path}' {
		${imports_code}
		${dts_code.map((line) => '\t' + line).join('\n')}
	}`;
	fs.writeFileSync(dts_path, dts_module);

	return {code:js_code_with_imports, path: js_path};
}

/**
 * Creates a Vite plugin for Go integration with SvelteKit
 * @param {GoKitOptions} [options={}] - Configuration options for the Go connector
 * @returns {import('vite').Plugin} Vite plugin instance
 */
export const gokit = function (options = {}) {
	/** @type {import('vite').ResolvedConfig} */
	let cfg;
	const module_regex = /\.remote\.go$/;

	return {
		name: 'vite-plugin-gokit',

		configResolved(config) {
			cfg = config;
			const remote_go_files = fg.sync('**/*.remote.go', { cwd: config.root });
			remote_go_files.forEach((file) => {
				const file_path = path.join(config.root, file);
				const go_code = fs.readFileSync(file_path, 'utf-8');
				transform_code(go_code, file_path, config, options);
			});
		},

		resolveId(id) {
			console.log('resolveId:', id);
			if (module_regex.test(id)) {
				// console.log({id});
			}
		},

		renderDynamicImport(id){

		},

		transform(src, id) {
			if (module_regex.test(id)) {
				return transform_code(src, id, cfg, options);
			}
		},

		load(id) {
			if(module_regex.test(id)) {
				console.log({id});
			}
		}
	};
};
