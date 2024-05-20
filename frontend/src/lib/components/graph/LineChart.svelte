<script>
	import { Line } from 'svelte-chartjs';

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

	export let requestResult;
	export let title;

	console.log(requestResult);

	let data = {
		labels: requestResult.labels,
		datasets: [
			{
				label: requestResult.measurement,
				data: requestResult.values,
				fill: false,
				borderColor: 'rgb(75, 192, 192)',
				tension: 0.1,
				pointRadius: 0
			}
		]
	};

	let options = {
		plugins: {
			title: {
				text: title,
				display: true,
				font: {
					size: 20
				}
			},
			legend: {
				display: false
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
				}
			},
			y: {
				title: {
					display: true,
					text: '(en ' + requestResult.field + ')'
				}
			}
		}
		// responsive: true
	};
</script>

<Line {data} {options} />
<div class = "source">
	<p>Source : <i>{requestResult.source}</i></p>
</div>
