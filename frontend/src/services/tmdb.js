const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';



const getAuthHeaders = () =>{
    //const user = JSON.parse(localStorage.getItem('user') || '{}');
    const token = sessionStorage.getItem('auth_token') || localStorage.getItem('auth_token');
    return {
        'Content-Type': 'application/json',
        ...(token ? { Authorization: `Bearer ${token}` } : {}),
        };

    };

const TMDBAPI = {
    async searchSeries(query) {
        const response = await fetch(`${API_BASE_URL}/tmdb/search/series?q=${encodeURIComponent(query)}`, {
            headers: getAuthHeaders(),
        });
        const data = await response.json();
        if (!response.ok) throw new Error(data.error || `HTTP ${response.status}`);
        return data.data || [];
    },
    async searchMovies(query) {
        const response = await fetch(`${API_BASE_URL}/tmdb/search/movies?q=${encodeURIComponent(query)}`, {
            headers: getAuthHeaders(),
        });
        const data = await response.json();
        if (!response.ok) throw new Error(data.error || `HTTP ${response.status}`);
        return data.data || [];
    },
    
};

export default TMDBAPI;