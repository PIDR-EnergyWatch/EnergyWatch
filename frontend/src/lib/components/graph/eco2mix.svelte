<script>
	export let res;
	import { Line } from 'svelte-chartjs';

	import {
		Chart as ChartJS,
		Title,
		Tooltip,
		Legend,
		ArcElement,
		CategoryScale,
		Filler
	} from 'chart.js';

	ChartJS.register(Title, Tooltip, Legend, ArcElement, CategoryScale, Filler);

	let data;
	let options;

	const labels = res.map((record) => record.heure);
	const fioul = {
		label: 'Fioul',
		data: res.map((record) => record.fioul),
		backgroundColor: 'rgb(128, 84, 159)',
		borderColor: 'rgb(128, 84, 159)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};
	const gaz = {
		label: 'Gaz',
		data: res.map((record) => record.gaz),
		backgroundColor: 'rgba(255, 0,0, 1)',
		borderColor: 'rgba(255,0,0, 1)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};

	const pompage = {
		label: 'Pompage',
		data: res.map((record) => record.pompage),
		backgroundColor: 'rgb(14, 66, 105)',
		borderColor: 'rgb(14, 66, 105)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};

	const ech_physiques = {
		label: 'Export / Import',
		data: res.map((record) => record.ech_physiques),
		backgroundColor: 'rgb(150, 150, 150)',
		borderColor: 'rgb(150, 150, 150)',
		borderWidth: 0,
		pointRadius: 0,
		fill: true
	};

	const charbon = {
		label: 'Charbon',
		data: res.map((record) => record.charbon),
		backgroundColor: 'rgba(0,0,0, 1)',
		borderColor: 'rgba(0,0,0, 1)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};
	const nucleaire = {
		label: 'Nucléaire',
		data: res.map((record) => record.nucleaire),
		backgroundColor: 'rgb(228, 167, 1)',
		borderColor: 'rgb(228, 167, 1)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};
	const eolien = {
		label: 'Eolien',
		data: res.map((record) => record.eolien),
		backgroundColor: 'rgb(114, 203, 183)',
		borderColor: 'rgb(114, 203, 183)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};
	const solaire = {
		label: 'Solaire',
		data: res.map((record) => record.solaire),
		backgroundColor: 'rgb(214, 107, 13)',
		borderColor: 'rgb(214, 107, 13)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};
	const hydraulique = {
		label: 'Hydraulique',
		data: res.map((record) => record.hydraulique),
		backgroundColor: 'rgb(38, 114, 176)',
		borderColor: 'rgb(38, 114, 176)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};
	const bioenergies = {
		label: 'Bioénergies',
		data: res.map((record) => record.bioenergies),
		backgroundColor: 'rgba( 35, 156, 0 ,1)',
		borderColor: 'rgba( 35, 156, 0 , 1)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};

	data = {
		labels,
		datasets: [
			ech_physiques,
			pompage,
			bioenergies,
			eolien,
			solaire,
			nucleaire,
			hydraulique,
			gaz,
			charbon,
			fioul
		]
	};

	options = {
		plugins: {
			title: {
				display: true,
				text: 'Eco2mix national (via RTE)'
			}
		},
		responsive: true,
		scales: {
			x: {
				display: true,
				title: {
					display: true,
					text: 'Heure'
				}
			},
			y: {
				display: true,
				title: {
					display: true,
					text: 'Production (en MW)'
				},
				stacked: true,
				beginAtZero: true
			}
		}
	};
</script>

<Line {data} {options} />
