import api from '$lib/api/api';
import { goto } from '$app/navigation';
import axios from 'axios';

export const login = async(username, password) => {
    try {
        await api.post('/api/auth/login', {
            username: username,
            password: password
        });
        goto('/dashboard');
    } catch (err) {
        return err.response?.data?.message;
    }
}

export const isLoggedIn = async() => {
    try {
        await api.get('/api/loggedIn');
        return true
    } catch (err) {
        // console.log(err)
        return false
    }
}

export const logout = async() => {
    await api.post('/api/auth/logout');
    goto('/login');
}

export const requestData = async(measurement) => {
    try {
        const response = await api.get(`/api/request?measurement=${measurement}`);
        return response;
    } catch (err) {
        return err.response?.data?.message;
    }
}

export const fetchEco2Mix = async() => {
    const today = new Date();
	const year = today.getFullYear();
	const month = String(today.getMonth() + 1).padStart(2, '0');
	const day = String(today.getDate()).padStart(2, '0');
	const formattedDate = `${year}%2F${month}%2F${day}`;
	const response = await axios.get(
			`https://odre.opendatasoft.com/api/explore/v2.1/catalog/datasets/eco2mix-national-tr/records?where=date%3D%22${year}-${month}-${day}%22&order_by=heure&limit=100&refine=date_heure%3A%22${formattedDate}%22 `
	);
    return response.data.results;
}

export const fetchEco2MixGrandEst = async() => {
    const today = new Date();
	const year = today.getFullYear();
	const month = String(today.getMonth() + 1).padStart(2, '0');
	const day = String(today.getDate()).padStart(2, '0');
	const formattedDate = `${year}%2F${month}%2F${day}`;
	const response = await axios.get(
			`https://odre.opendatasoft.com/api/explore/v2.1/catalog/datasets/eco2mix-regional-tr/records?where=date%3D%22${year}-${month}-${day}%22&order_by=heure&limit=100&refine=libelle_region%3A%22Grand%20Est%22&refine=date_heure:%22${formattedDate}%22`
	);
    return response.data.results;
}


export const fetchWeather = async() => {
    const response = await axios.get('https://api.weather.com/v2/pws/observations/hourly/7day?stationId=IVANDU8&format=json&units=e&apiKey=4e78f553ad464503b8f553ad46b50386');
    return response.data;

}

export const fetchOpenWeather = async() => {
    const response = await axios.get('https://api.openweathermap.org/data/2.5/weather?lat=48.650002&lon=6.18333&unit=metric&appid=6c7536caec8cea574d9b049fde683a35')
    return response.data
}

export const fetchOpenWeatherForecast = async() => {
    const response = await axios.get('https://api.openweathermap.org/data/2.5/forecast?lat=48.650002&lon=6.18333&unit=metric&appid=6c7536caec8cea574d9b049fde683a35')
    return response.data
}