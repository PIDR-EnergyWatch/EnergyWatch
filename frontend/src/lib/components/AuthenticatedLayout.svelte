<script>
	import { onMount } from 'svelte';
	import { isLoggedIn } from '$lib/api/requests';
	import Navbar from '$lib/components/Navbar.svelte';

	let loggedIn = undefined;

	onMount(async () => {
		loggedIn = await isLoggedIn();
	});
</script>

{#if loggedIn === true}
	<Navbar />
	<slot />
{:else}
	You need to be logged in.
{/if}
