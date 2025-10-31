export const imports = `import { _imports_ } from '$app/server';`;

export const queryFn = `
export const _name_ = query(() =>{
	return fetch('_endpoint_/_path_', {
		headers: {
			'Content-Type': 'application/json',
		},
	}).then(res => res.json());
});
`;

export const dtsQueryFn = `export const _name_: query<any>;`;

export const dtsFormFn = `export const _name_: form<any>;`;

export const dtsCommandFn = `export const _name_: command<any>;`;
