import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:8000/',
});

api.interceptors.request.use(
    (config) => {
        const user = localStorage.getItem('user');
        if (user === null) return config;
        //if (!user) return ;
        // @ts-ignore
        const token = JSON.parse(user).accessToken;
        if (token) {
            config.headers.Authorization = `Bearer ${token.trim()}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);


export default api;