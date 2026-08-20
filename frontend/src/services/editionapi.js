import { apiGet, apiPost, apiPut, apiDelete } from './apiClient';


class EditionAPI {
    static async getAllEditions() {
        try {
            const response = await apiGet('/editions');
            return response?.data || [];
        } catch (error) {
            console.error('Fehler beim Abrufen der Editionen:', error);
            throw error;
        }
    }
}


export default EditionAPI;