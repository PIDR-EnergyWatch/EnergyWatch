<script>
	import AuthenticatedLayout from '$lib/components/AuthenticatedLayout.svelte';
	import { requestData, fetchEco2Mix } from '$lib/api/requests';
	import { onMount } from 'svelte';
	import Card from '$lib/components/Card.svelte';

	let graphs = undefined;

	onMount(async () => {
		let res = await requestData('PAT');
		let eco2mix = await fetchEco2Mix();

		graphs = [
			{
				title: 'PAT',
				type: 'line',
				res
			},
			{
				title: 'Eco2mix',
				type: 'eco2mix',
				res: eco2mix
			}
		];
	});
</script>

<AuthenticatedLayout>
	<section class="m-8">
		{#if graphs}
			<Card {graphs} />
		{:else}
			<div
				class="inline-block h-8 w-8 animate-spin rounded-full border-4 border-solid border-current border-e-transparent align-[-0.125em] text-surface motion-reduce:animate-[spin_1.5s_linear_infinite] dark:text-white"
				role="status"
			>
				<span
					class="!absolute !-m-px !h-px !w-px !overflow-hidden !whitespace-nowrap !border-0 !p-0 ![clip:rect(0,0,0,0)]"
					>Loading...</span
				>
			</div>
		{/if}
	</section>
</AuthenticatedLayout>
