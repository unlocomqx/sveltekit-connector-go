import { query, form } from '$app/server';

export const queryTodos = query(() =>{
	return fetch('http://localhost:9999/rpc/src/routes/todos.remote.go', {
		headers: {
			'Content-Type': 'application/json',
		},
	}).then(res => res.json());
});


export const queryTodoByID = query(() =>{
	return fetch('http://localhost:9999/rpc/src/routes/todos.remote.go', {
		headers: {
			'Content-Type': 'application/json',
		},
	}).then(res => res.json());
});

