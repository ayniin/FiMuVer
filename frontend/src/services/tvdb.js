const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';

const getAuthHeaders = () => {
    const user = JSON.parse(localStorage.getItem('user') || '{}');
    return {
        'Content-Type': 'application/json',
        ...(user.token ? { Authorization: `Bearer ${user.token}` } : {}),
        };
    };

const TVDBAPI = {
    async searchSeries(query) {
        const response = await fetch(`${API_BASE_URL}/tvdb/search/series?q=${encodeURIComponent(query)}`, {
            headers: getAuthHeaders(),
        });
        const data = await response.json();
        if (!response.ok) throw new Error(data.error || `HTTP ${response.status}`);
        return data.data || [];
    },

    async searchMovies(query) {
        const response = await fetch(`${API_BASE_URL}/tvdb/search/movies?q=${encodeURIComponent(query)}`, {
            headers: getAuthHeaders(),
        });
        const data = await response.json();
        if (!response.ok) throw new Error(data.error || `HTTP ${response.status}`);
        return data.data || [];
    },

};

export default TVDBAPI;
