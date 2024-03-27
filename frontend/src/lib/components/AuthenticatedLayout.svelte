<script>
	import { onMount } from 'svelte';
	import { isLoggedIn } from '$lib/api/requests';

	let loggedIn = undefined;

	onMount(async () => {
		loggedIn = await isLoggedIn();
	});
</script>

{#if loggedIn === true}
	<slot />
{:else}
	You need to be logged in.
{/if}
