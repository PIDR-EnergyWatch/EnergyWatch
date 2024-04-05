<script>
	import AuthenticatedLayout from '../../lib/components/AuthenticatedLayout.svelte';
	import { logout, requestData } from '$lib/api/requests';
	import { Line } from 'svelte-chartjs';
	import { onMount } from 'svelte';

	import {
		Chart as ChartJS,
		Title,
		Tooltip,
		Legend,
		LineElement,
		LinearScale,
		PointElement,
		CategoryScale,
		TimeScale
	} from 'chart.js';
	import 'chartjs-adapter-date-fns';

	ChartJS.register(
		Title,
		Tooltip,
		Legend,
		LineElement,
		LinearScale,
		PointElement,
		CategoryScale,
		TimeScale
	);

	let res = undefined;
	let data;
	let options;

	onMount(async () => {
		res = await requestData('PAT');

		data = {
			labels: res.labels,
			datasets: [
				{
					label: res.measurement,
					data: res.values,
					fill: false,
					borderColor: 'rgb(75, 192, 192)',
					tension: 0.1
				}
			]
		};

		options = {
			plugins: {
				title: {
					text: 'Chart.js Time Scale',
					display: true
				}
			},
			scales: {
				x: {
					type: 'time',
					display: true,
					title: {
						display: true,
						text: 'Date'
					},
					ticks: {
						autoSkip: true,
						maxRotation: 100,
						major: {
							enabled: true
						}
					},
					y: {
						title: {
							display: true,
							text: res.field
						}
					}
				}
			}
		};
	});
</script>

<AuthenticatedLayout>
	<h1 class="text-3xl font-bold">Welcome!</h1>
	<section>
		{#if res}
			<Line {data} {options} />
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
	<button
		class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
		on:click={async () => await logout()}
	>
		Logout
	</button>
</AuthenticatedLayout>
