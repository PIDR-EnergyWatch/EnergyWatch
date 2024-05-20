<script>
	import AuthenticatedLayout from '$lib/components/AuthenticatedLayout.svelte';
	// import Carousel from 'svelte-carousel';
	import Carousel from '$lib/components/Carousel.svelte';
	import { getAllGraphs } from '$lib/api/fetchGraphs';
	import { onMount } from 'svelte';
	import Card from '$lib/components/Card.svelte';

	let graphs = undefined;

	onMount(async () => {
		graphs = await getAllGraphs();
	});
</script>

<AuthenticatedLayout>
	<section class="m-8">
		{#if graphs}
			<Carousel autoplay="5000">
				{#each graphs as graph (graph.title)}
					<div class="p-8 flex items-center justify-center">
						<div class="w-full">
							<Card {graph} />
						</div>
					</div>
				{/each}
			</Carousel>
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
