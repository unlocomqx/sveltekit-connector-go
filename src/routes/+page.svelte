<script lang="ts">
	// still can't get .d.ts to work
	import { FormCreateTodo, QueryTodos, FormDeleteTodo } from './todos.remote.go';
</script>

<h1>SvelteKit Connector for Go</h1>
<p>Connect your SvelteKit frontend to Go backend easily using this library.</p>

<div>
	<h2>Todos list:</h2>
	{#each await QueryTodos() as { id, title } (id)}
		<form {...FormDeleteTodo.for(id)}>
			<span>{title}</span>
			<button>Delete</button>
			<br />
			<br />
		</form>
	{/each}

	<form method="post" {...FormCreateTodo}>
		<input name="title" type="text" value="New todo" />
		<button type="submit">Add</button>
	</form>
</div>

<p>
	<a href="/users">Users</a>
</p>
