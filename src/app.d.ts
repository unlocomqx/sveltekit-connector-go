// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};

declare module '*.remote.go' {
	import { form, query } from '$app/server';
	export const queryTodos: query<any>;
	export const queryTodoByID: query<any>;
	export const formCreateTodo: form<any>;
}
