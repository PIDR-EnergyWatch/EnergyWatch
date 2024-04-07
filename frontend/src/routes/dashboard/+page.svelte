<script>
	import AuthenticatedLayout from '$lib/components/AuthenticatedLayout.svelte';
	import { logout, requestData, fetchEco2Mix } from '$lib/api/requests';
	import { onMount } from 'svelte';
	import PatLayout from '$lib/components/graph/PATLayout.svelte';
	import Eco2mix from '$lib/components/graph/eco2mix.svelte';

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
	let eco2mix = undefined;

	onMount(async () => {
		res = await requestData('PAT');
		eco2mix = await fetchEco2Mix();
	});
</script>

<AuthenticatedLayout>
	<h1 class="text-3xl font-bold">Welcome!</h1>
	<section>
		{#if res}
			<PatLayout requestResult={res} title={'PAT'}/>
			
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
		{#if eco2mix}
			<Eco2mix res={eco2mix} />
		{/if}

	</section>
	<button
		class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
		on:click={async () => await logout()}
	>
		Logout
	</button>
</AuthenticatedLayout>
