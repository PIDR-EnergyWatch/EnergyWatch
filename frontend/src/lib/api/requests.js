import api from '$lib/api/api';
import { goto } from '$app/navigation';

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