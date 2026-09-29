const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';


const getAuthHeaders = () =>{
    const token = sessionStorage.getItem('auth_token') || localStorage.getItem('auth_token');
    return {
        'Content-Type': 'application/json',
        ...(token ? { Authorization: `Bearer ${token}` } : {}),
        };

    };

const DiscogsAPI = {
    async searchReleases(query) {
        const response = await fetch(`${API_BASE_URL}/discogs/search/releases?q=${encodeURIComponent(query)}`, {
            headers: getAuthHeaders(),
        });
        const data = await response.json();
        if (!response.ok) throw new Error(data.error || `HTTP ${response.status}`);
        return data.data || [];
    },

    async searchMasters(query) {
        const response = await fetch(`${API_BASE_URL}/discogs/search/masters?q=${encodeURIComponent(query)}`, {
            headers: getAuthHeaders(),
        });
        const data = await response.json();
        if (!response.ok) throw new Error(data.error || `HTTP ${response.status}`);
        return data.data || [];
    },

    async searchArtists(query) {
        const response = await fetch(`${API_BASE_URL}/discogs/search/artists?q=${encodeURIComponent(query)}`, {
            headers: getAuthHeaders(),
        });
        const data = await response.json();
        if (!response.ok) throw new Error(data.error || `HTTP ${response.status}`);
        return data.data || [];
    },

};

export default DiscogsAPI;
