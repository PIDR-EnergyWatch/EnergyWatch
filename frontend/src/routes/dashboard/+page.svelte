<script>
	import AuthenticatedLayout from '$lib/components/AuthenticatedLayout.svelte';
	import { requestData, fetchEco2Mix, fetchWeather, fetchOpenWeather, fetchOpenWeatherForecast, fetchEco2MixGrandEst } from '$lib/api/requests';
	import { onMount } from 'svelte';
	import Card from '$lib/components/Card.svelte';

	let graphs = undefined;

	onMount(async () => {
		let res = await requestData('PAT');
		let eco2mix = await fetchEco2Mix();
		let eco2mixGrandEst = await fetchEco2MixGrandEst();
		let weather = await fetchWeather();
		let openWeather = await fetchOpenWeather();
		let openWeatherForecast = await fetchOpenWeatherForecast();
		console.log(weather);
		graphs = [
			{
				title: 'Eco2mix Grand Est',
				type: 'eco2mix',
				res: eco2mixGrandEst
			},
			{
				title: 'Eco2mix National',
				type: 'eco2mix',
				res: eco2mix
			},
			
			{
				title: 'PAT',
				type: 'line',
				res
			},
			{
				title: 'Température',
				type : 'line',
				res : {
					values : weather.observations.map(value => (value.imperial.tempAvg-32)*5/9),
					labels : weather.observations.map(value => value.obsTimeLocal),
					field : '°C'
				}
			},
			{
				title: 'Weather',
				type: 'weather',
				res: {
					weather : weather,
					openWeather : openWeather,
					openWeatherForecast : openWeatherForecast
				}
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
