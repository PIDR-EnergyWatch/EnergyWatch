<script>
	export let graphs;
	import LineChart from '$lib/components/graph/LineChart.svelte';
	import Eco2mix from '$lib/components/graph/eco2mix.svelte';
	import Weather from '$lib/components/graph/CurrentWeather.svelte';
	import '$lib/style/card.css'

</script>

<ul
	role="list"
	class="grid grid-cols-1 gap-x-4 gap-y-8 sm:grid-cols-2 sm:gap-x-6 lg:grid-cols-2 xl:gap-x-8"
>
	{#each graphs as graph (graph.title)}
		<li class="relative">
			<div
				class="group aspect-auto block w-full overflow-hidden rounded-lg bg-gray-100 focus-within:ring-2 focus-within:ring-indigo-500 focus-within:ring-offset-2 focus-within:ring-offset-gray-100"
			>
				{#if graph.type === 'line'}
					<LineChart requestResult={graph.res} title={graph.title} />
					<div class="source">
						<p>Source : <i>{graph.res.source}</i></p>
					</div>
				{:else if graph.type === 'eco2mix'}
					<Eco2mix res={graph.res} title={graph.title}/>
					<div class="source">
						<p>Source : <i>{graph.res.source}</i></p>
					</div>
				{:else if graph.type === 'weather'}
					<Weather data={graph.res} />
					<div class="source">
						<p>Source : <i>{graph.res.weather.source}</i> & <i>{graph.res.openWeather.source}</i></p>
					</div>
				{/if}
				<button type="button" class="absolute inset-0 focus:outline-none">
					<span class="sr-only">View details for {graph.title}</span>
				</button>
			</div>
		</li>
	{/each}
</ul>

<style>
	/* Add any component-specific styles here */
</style>
