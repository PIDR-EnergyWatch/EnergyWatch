import {
    requestData,
    fetchEco2Mix,
    fetchWeather,
    fetchOpenWeather,
    fetchOpenWeatherForecast,
    fetchEco2MixGrandEst
} from '$lib/api/requests';

export const getAllGraphs = async() => {
    let graphs = undefined;

    let temp = await requestData('temperature');
    let pat = await requestData('PAT');
    let eco2mix = await fetchEco2Mix();
    let eco2mixGrandEst = await fetchEco2MixGrandEst();
    let weather = await fetchWeather();
    let openWeather = await fetchOpenWeather();
    let openWeatherForecast = await fetchOpenWeatherForecast();

    graphs = [
        {
            title: 'Eco2mix Grand Est',
            type: 'eco2mix',
            res: eco2mixGrandEst,
            from: "Réseau de Transport d'Électricité"
        },
        {
            title: 'Eco2mix National',
            type: 'eco2mix',
            res: eco2mix,
            from: "Réseau de Transport d'Électricité"
        },
        {
            title: 'Température salle 206',
            type: 'line',
            res: temp,
            from: "Capteur ENSEM"
        },
        {
            title: 'PAT',
            type: 'line',
            res: pat,
            from: "Capteur UL"
        },
        {
            title: 'Température extérieure',
            type: 'line',
            res: {
                values: weather.observations.map((value) => ((value.imperial.tempAvg - 32) * 5) / 9),
                labels: weather.observations.map((value) => value.obsTimeLocal),
                field: '°C'
            },
            from: "Station Météo ENSEM"
        },
        {
            title: 'Weather',
            type: 'weather',
            res: {
                weather: weather,
                openWeather: openWeather,
                openWeatherForecast: openWeatherForecast
            },
            from: "OpenWeatherMap"
        }
    ];

    return graphs;
}