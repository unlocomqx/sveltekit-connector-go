<script lang="ts">
	// still can't get .d.ts to work
	import { QueryUserInfo } from './users.remote.go';

	let infos = $state({});
</script>

<h1>SvelteKit Connector for Go</h1>
<p>Connect your SvelteKit frontend to Go backend easily using this library.</p>

<button onclick={async () => (infos = await QueryUserInfo())}> Fetch user info</button>

<div>
	{JSON.stringify(infos)}
</div>
