<script>
	export let graph;

	import LineChart from '$lib/components/graph/LineChart.svelte';
	import Eco2mix from '$lib/components/graph/eco2mix.svelte';
	import Weather from '$lib/components/graph/CurrentWeather.svelte';
</script>

<div
	class="overflow-visible group aspect-auto block w-full overflow-hidden rounded-lg bg-gray-100 focus-within:ring-2 focus-within:ring-indigo-500 focus-within:ring-offset-2 focus-within:ring-offset-gray-100"
>
	<span
		class="absolute -m-4 items-center rounded-md bg-purple-50 px-2 py-1 text-s font-medium text-purple-700 ring-1 ring-inset ring-purple-700/10"
	>
		{graph.from}</span
	>

	{#if graph.type === 'line'}
		<LineChart requestResult={graph.res} title={graph.title} />
	{:else if graph.type === 'eco2mix'}
		<Eco2mix res={graph.res} title={graph.title} />
	{:else if graph.type === 'weather'}
		<Weather data={graph.res} />
	{/if}
	<button type="button" class="absolute inset-0 focus:outline-none">
		<span class="sr-only">View details for {graph.title}</span>
	</button>
</div>

<style>
	/* Add any component-specific styles here */
</style>
