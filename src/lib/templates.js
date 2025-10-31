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
