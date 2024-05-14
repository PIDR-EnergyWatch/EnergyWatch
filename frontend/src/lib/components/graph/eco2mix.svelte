<script>
	export let res;
	export let title;
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
		backgroundColor: 'rgb(196, 131, 242)',
		borderColor: 'rgb(196, 131, 242)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};
	const gaz = {
		label: 'Gaz',
		data: res.map((record) => record.gaz),
		backgroundColor: 'rgba(242, 90, 90, 1)',
		borderColor: 'rgba(242, 90, 90, 1)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};

	const pompage = {
		label: 'Pompage',
		data: res.map((record) => record.pompage),
		backgroundColor: 'rgb(47, 103, 145)',
		borderColor: 'rgb(47, 103, 145)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};

	const ech_physiques = {
		label: 'Export / Import',
		data: res.map((record) => record.ech_physiques),
		backgroundColor: 'rgb(200, 200, 200)',
		borderColor: 'rgb(200, 200, 200)',
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
		backgroundColor: 'rgb(252, 215, 114)',
		borderColor: 'rgb(252, 215, 114)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};
	const eolien = {
		label: 'Eolien',
		data: res.map((record) => record.eolien),
		backgroundColor: 'rgb(173, 247, 231)',
		borderColor: 'rgb(173, 247, 231)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};
	const solaire = {
		label: 'Solaire',
		data: res.map((record) => record.solaire),
		backgroundColor: 'rgb(227, 152, 86)',
		borderColor: 'rgb(227, 152, 86)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};
	const hydraulique = {
		label: 'Hydraulique',
		data: res.map((record) => record.hydraulique),
		backgroundColor: 'rgb(43, 158, 252)',
		borderColor: 'rgb(43, 158, 252)',
		borderWidth: 0,
		fill: true,
		pointRadius: 0
	};
	const bioenergies = {
		label: 'Bioénergies',
		data: res.map((record) => record.bioenergies),
		backgroundColor: 'rgb(101, 199, 74)',
		borderColor: 'rgba( 101, 199, 74, 1)',
		borderWidth:0,
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
				text: title,
				font: {
					size: 20
				}
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
