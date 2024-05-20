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

  export let res;
  export let title;

  let puissanceOrdinateur = res.puissanceOrdinateur;
  let puissanceEcranOrdinateur = res.puissanceEcranOrdinateur;

  let data = {
    labels: puissanceEcranOrdinateur.labels, // Assuming that both requestResult1 and requestResult2 have the same labels
    datasets: [
      {
        label: puissanceOrdinateur.title,
        data: puissanceOrdinateur.values.map((value, index) => value*18),
        fill: false,
        borderColor: 'rgb(75, 192, 192)',
        tension: 0.1,
        pointRadius: 0,
        borderWidth: 2
      },
      {
        label: puissanceEcranOrdinateur.title,
        data: puissanceEcranOrdinateur.values.map((value, index) => value*18),
        fill: false,
        borderColor: 'rgb(255, 99, 132)', // Different color for the second dataset
        tension: 0.1,
        pointRadius: 0,
        borderWidth: 2

      },
      {
        label: 'Total',
        data: puissanceOrdinateur.values.map((value, index) => (value + puissanceEcranOrdinateur.values[index])*18),
        fill: false,
        borderColor: 'rgb(54, 162, 235)', // Different color for the third dataset
        tension: 0.1,
        pointRadius: 0,
        borderWidth: 2

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
        }
      },
      y: {
        title: {
          display: true,
          text: '(en ' + puissanceEcranOrdinateur.field + ')' // Assuming that both requestResult1 and requestResult2 have the same field
        },
        beginAtZero: true
      }
    }
    // responsive: true
  };
</script>

<Line {data} {options} />
