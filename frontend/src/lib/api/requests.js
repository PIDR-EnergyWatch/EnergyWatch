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
			`https://odre.opendatasoft.com/api/explore/v2.1/catalog/datasets/eco2mix-national-tr/records?limit=20&refine=date_heure%3A%22${formattedDate}%22`
	);
    return response.data.results;
}