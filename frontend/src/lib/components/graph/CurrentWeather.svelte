<script>
    export let data;

    const weather = data.weather;
    const openWeather = data.openWeather;
    const openWeatherForecast = data.openWeatherForecast;

    const current = weather.observations[weather.observations.length - 1]

    console.log(openWeather);


    console.log(current);

</script>

<div>
    <h2>Station méteo de l'ENSEM</h2>
    <div>
        <p>Temperature: {((current.imperial.tempAvg - 32)*(5/9)).toFixed(2)}°C</p>
        <p>Humidity: {current.humidityAvg}%</p>
        <p>Wind Speed: {(current.imperial.windspeedAvg * 1.60934).toFixed(2)} km/h</p>
        <p>Indice UV: {current.uvHigh}</p>
        <p>{openWeather.weather[0].description}</p>
        <div style="background-color:wheat; width: fit-content; border-radius : 50px;">
            <img alt="Weather Icon" src="https://openweathermap.org/img/wn/{openWeather.weather[0].icon}@2x.png">
        </div>
    </div>
    <h2>Forecast</h2>
    {#each openWeatherForecast.list as item, index}
        {#if index < 4}
            <p>{item.dt_txt}</p>
            <p>{(item.main.temp - 273).toFixed(2)} °C</p>
            <div style="background-color:wheat; width: fit-content; border-radius : 50px;">
                <img alt="Weather Icon" src="https://openweathermap.org/img/wn/{item.weather[0].icon}@2x.png">
            </div>
        {/if}
    {/each}
</div>