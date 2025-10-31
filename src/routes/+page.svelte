<script lang="ts">
	import { queryTodos } from './todos.remote.go';

	let todos = $state([]);
</script>

<h1>SvelteKit Connector for Go</h1>
<p>Connect your SvelteKit frontend to Go backend easily using this library.</p>

<button onclick={async () => (todos = await queryTodos())}> Fetch todos </button>

<div>
	<h2>Todos</h2>
	{#each todos as { id, title } (id)}
		<p>{title}</p>
	{/each}
</div>
