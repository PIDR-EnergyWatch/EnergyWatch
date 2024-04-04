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

	let res;
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
		<Line {data} {options} />
	</section>
	<button
		class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
		on:click={async () => await logout()}
	>
		Logout
	</button>
</AuthenticatedLayout>
