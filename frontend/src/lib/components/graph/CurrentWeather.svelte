<script>
    export let data;
    import '$lib/style/weather.css'
    const weather = data.weather;
    const openWeather = data.openWeather;
    const openWeatherForecast = data.openWeatherForecast;

    const current = weather.observations[weather.observations.length - 1]


    function convertTime(dt_text) {
        let date = new Date(dt_text);
        let timeString = date.toLocaleTimeString('fr-FR', { hour: '2-digit', minute: '2-digit' });
        return timeString;
  }
    

</script>

<div class="weatherDiv">
    <h2>Station méteo de l'ENSEM</h2>
        <div class="row">
                <div style="background-color:wheat; width: fit-content; border-radius : 50px; margin-right:5%;">
                    <img alt="Weather Icon" src="https://openweathermap.org/img/wn/{openWeather.weather[0].icon}@2x.png">
                </div>
            <div class="column" style="width: 50%;">
                <div class="row space_between">
                    <p>Température : <b>{((current.imperial.tempAvg - 32) * 5 / 9).toFixed(2)} °C</b></p>
                    <p>Humidité : <b>{current.humidityHigh} %</b></p>
                </div>
                <div class="row space_between">
                    <p>Vitesse du vent : <b>{(current.imperial.windspeedAvg * 1.60934).toFixed(2)} km/h</b></p>
                    <p>Indice UV : <b>{current.uvHigh}</b></p>
                </div>
            </div>
        </div>        
    <h2>Prévision méteo</h2>
    <div class="row space_between">
    {#each openWeatherForecast.list as item, index}
    
        {#if index < 4}
        <div class="column ">
        <p>{convertTime(item.dt_txt)}</p>
        <p><b>{(item.main.temp - 273).toFixed(0)} °C</b></p>
        <div style="background-color:wheat; width: fit-content; border-radius : 50px;">
            <img alt="Weather Icon" src="https://openweathermap.org/img/wn/{item.weather[0].icon}@2x.png">
        </div>
    </div>

    {/if}
       
    {/each}
    </div>
    
</div>
<div class = "source">
    <p>Source : <i>{weather.source}</i> & <i>{openWeather.source}</i></p>
</div>