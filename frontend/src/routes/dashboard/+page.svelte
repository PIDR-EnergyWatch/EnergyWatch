<script>
	import AuthenticatedLayout from '../../lib/components/AuthenticatedLayout.svelte';
	import { logout } from '$lib/api/requests';
</script>

<AuthenticatedLayout>
	<h1 class="text-3xl font-bold">Welcome!</h1>
	<button
		class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
		on:click={async () => await logout()}
	>
		Logout
	</button>
</AuthenticatedLayout>
