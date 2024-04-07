<script>
    export let res
    import { onMount } from 'svelte';
    import { Doughnut } from 'svelte-chartjs';
    import { Line } from 'svelte-chartjs';
    import axios from 'axios';
  
    import {
      Chart as ChartJS,
      Title,
      Tooltip,
      Legend,
      ArcElement,
      CategoryScale,
      Filler,
    } from 'chart.js';
  
    import ChartDataLabels from 'chartjs-plugin-datalabels';
  
    ChartJS.register(Title, Tooltip, Legend, ArcElement, CategoryScale, ChartDataLabels, Filler);
  
    let data = {};
    let options = {};
    
    const labels = res.data.results.map((record) => record.heure);
    const fioul = {
      label: 'Fioul',
      data: res.data.results.map((record) => record.fioul),
      backgroundColor: 'rgb(128, 84, 159)',
      borderColor: 'rgb(128, 84, 159)',
      borderWidth: 0,
      fill : true,
      pointRadius: 0,
    };
    const gaz = {
      label: 'Gaz',
      data: res.data.results.map((record) => record.gaz),
      backgroundColor: 'rgba(255, 0,0, 1)',
      borderColor: 'rgba(255,0,0, 1)',
      borderWidth: 0,
      fill : true,
      pointRadius: 0,
    };

   

    const pompage = {
      label: 'Pompage',
      data: res.data.results.map((record) => record.pompage),
      backgroundColor: 'rgb(14, 66, 105)',
      borderColor: 'rgb(14, 66, 105)',
      borderWidth: 0,
      fill : true,
      pointRadius: 0,

    };

    const ech_physiques = {
      label: 'Export / Import',
      data: res.data.results.map((record) => record.ech_physiques),
      backgroundColor: 'rgb(150, 150, 150)',
      borderColor: 'rgb(150, 150, 150)',
      borderWidth: 0,
      pointRadius: 0,
      fill : true,

    };

    const charbon = {
      label: 'Charbon',
      data: res.data.results.map((record) => record.charbon),
      backgroundColor: 'rgba(0,0,0, 1)',
      borderColor: 'rgba(0,0,0, 1)',
      borderWidth: 0,
      fill : true,
      pointRadius: 0,

    };
    const nucleaire = {
      label: 'Nucléaire',
      data: res.data.results.map((record) => record.nucleaire),
      backgroundColor: 'rgb(228, 167, 1)',
      borderColor: 'rgb(228, 167, 1)',
      borderWidth: 0,
      fill : true,
      pointRadius: 0,
    };
    const eolien = {
      label: 'Eolien',
      data: res.data.results.map((record) => record.eolien),
      backgroundColor: 'rgb(114, 203, 183)',
      borderColor: 'rgb(114, 203, 183)',
      borderWidth:0,
      fill: true,
      pointRadius: 0,
    };
    const solaire = {
      label: 'Solaire',
      data: res.data.results.map((record) => record.solaire),
      backgroundColor: 'rgb(214, 107, 13)',
      borderColor: 'rgb(214, 107, 13)',
      borderWidth: 0,
      fill : true,
      pointRadius: 0,

    };
    const hydraulique = {
      label: 'Hydraulique',
      data: res.data.results.map((record) => record.hydraulique),
      backgroundColor: 'rgb(38, 114, 176)',
      borderColor: 'rgb(38, 114, 176)',
      borderWidth: 0,
      fill : true,
      pointRadius: 0,

    };
    const bioenergies = {
      label: 'Bioénergies',
      data: res.data.results.map((record) => record.bioenergies),
      backgroundColor: 'rgba( 35, 156, 0 ,1)',
      borderColor: 'rgba( 35, 156, 0 , 1)',
      borderWidth: 0,
      fill : true,
      pointRadius: 0,

    };

    data = {
      labels,
      datasets: [ech_physiques, pompage, bioenergies, eolien, solaire, nucleaire, hydraulique, gaz, charbon,fioul],
    };

    options = {
      plugins: {
        title: {
          display: true,
          text: 'Eco2mix national (via RTE)',
        },
        legend: {
          display: true,
          position: 'right',
        },
        datalabels: {
          display: false,
        },
      },
      responsive: true,
        maintainAspectRatio: false,
        scales : {
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
            beginAtZero: true,
          }
        }
    };

</script>
  
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
  