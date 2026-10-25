import { apiGet } from './apiClient';

// Filme werden über TMDB gesucht, gespeichert wird aber die IMDb-ID (getImdbId). Serien kommen von TVDB.
const TVDBAPI = {
    async searchSeries(query) {
        const response = await apiGet(`/tvdb/search/series?q=${encodeURIComponent(query)}`);
        return response?.data || [];
    },

    async searchMovies(query) {
        const response = await apiGet(`/tmdb/search/movies?q=${encodeURIComponent(query)}`);
        return response?.data || [];
    },

    async getImdbId(tmdbId) {
        const response = await apiGet(`/tmdb/movies/${encodeURIComponent(tmdbId)}/imdb`);
        return response?.data?.imdb_id || '';
    },
};

export default TVDBAPI;
