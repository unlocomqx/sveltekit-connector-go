import fg from 'fast-glob';
import Parser from 'tree-sitter';
import goLang from 'tree-sitter-go';
import path from 'node:path';
import * as fs from 'node:fs';
import { imports, queryFn } from './templates.js';

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
 * Emits or processes remote functions found in a Go file
 * @param {Object} params - The parameters object
 * @param {string} params.root - The root directory path
 * @param {string} params.file_path - The absolute path to the Go file containing remote functions
 * @param {RemoteFunction[]} params.remote_functions - Array of remote function objects
 * @param {GoKitOptions} params.options - Configuration options for the Go connector
 */
function emit_remote_functions({ root, file_path, remote_functions, options }) {
	console.log(`Remote functions in ${file_path}:`, remote_functions);
	const js_path = path.join(path.dirname(file_path), path.basename(file_path, '.go') + '.js');
	const js_code = remote_functions.map((remote_fn) => {
		const relative_path = path.relative(root, file_path);
		if (remote_fn.type === 'query') {
			return queryFn
				.replace('_name_', remote_fn.name)
				.replace('_endpoint_', options.endpoint.replace(/\/$/, '') || '')
				.replace('_path_', relative_path);
		}
		return '';
	});

	const imports_list = new Set(remote_functions.map((fn) => fn.type));
	const imports_code = imports.replace('_imports_', Array.from(imports_list).join(', '));
	const js_code_with_imports = [imports_code, ...js_code].join('\n');
	fs.writeFileSync(js_path, js_code_with_imports);
}

/**
 * Creates a Vite plugin for Go integration with SvelteKit
 * @param {GoKitOptions} [options={}] - Configuration options for the Go connector
 * @returns {import('vite').Plugin} Vite plugin instance
 */
export const gokit = function (options = {}) {
	const module_regex = /\.remote\.go$/;

	const parser = new Parser();
	parser.setLanguage(goLang);

	return {
		name: 'vite-plugin-gokit',

		configResolved(config) {
			const remote_go_files = fg.sync('**/*.remote.go', { cwd: config.root });
			remote_go_files.forEach((file) => {
				const file_path = path.join(config.root, file);
				const go_code = fs.readFileSync(file_path, 'utf-8');
				const tree = parser.parse(go_code);

				/** @type {RemoteFunction[]} */
				const remote_functions = [];
				for (const node of tree.rootNode.children) {
					console.log(node.type);
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

				emit_remote_functions({ root: config.root, file_path, remote_functions, options });
			});
		},

		resolveId(id) {
			// console.log({ id });
		},

		transform(src, id) {
			if (module_regex.test(id)) {
				console.log({ id, src });
			}
		},

		configureServer(server) {}
	};
};
