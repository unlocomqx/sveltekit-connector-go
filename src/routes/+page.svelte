<script lang="ts">
	// still can't get .d.ts to work
	import { formUpdateTodo, queryTodos } from './todos.remote.go';

	let todos = $state(await queryTodos());
</script>

<h1>SvelteKit Connector for Go</h1>
<p>Connect your SvelteKit frontend to Go backend easily using this library.</p>

<button onclick={async () => (todos = await queryTodos())}> Fetch todos</button>

<div>
	<h2>Todos list:</h2>
	{#each todos as { id, title } (id)}
		<form {...formUpdateTodo}>
			<input name="" type="text" value={title}>
			<button type="submit">Update</button>
		</form>
		<br>
	{/each}
</div>
