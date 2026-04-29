const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';

// Helper-Funktion für headers mit token
const getHeaders = (additionalHeaders = {}) => {
  const headers = {
    'Content-Type': 'application/json',
    ...additionalHeaders,
  };

  const user = JSON.parse(localStorage.getItem('user'));
  if (user?.token) {
    headers['Authorization'] = `Bearer ${user.token}`;
  }

  return headers;
};

class CollectionAPI {
    // Alle Collections für den authentifizierten Benutzer abrufen
    static async getAllCollectionsForUser() {
        try {
            const response = await fetch(`${API_BASE_URL}/collections`, {
                headers: getHeaders(),
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            return data.data || [];
        } catch (error) {
            console.error('Fehler beim Abrufen der Collections:', error);
            throw error;
        }
    }

    // Collection nach ID abrufen
    static async getCollectionById(id) {
        try {
            const response = await fetch(`${API_BASE_URL}/collections/${id}`, {
                headers: getHeaders(),
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            return data.data;
        } catch (error) {
            console.error(`Fehler beim Abrufen der Collection ${id}:`, error);
            throw error;
        }
    }

    // Neue Collection erstellen
    static async createCollection(collectionData) {
        try {
            const response = await fetch(`${API_BASE_URL}/collections`, {
                method: 'POST',
                headers: getHeaders(),
                body: JSON.stringify(collectionData),
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            return data.data;
        } catch (error) {
            console.error('Fehler beim Erstellen der Collection:', error);
            throw error;
        }
    }

    // Collection aktualisieren
    static async updateCollection(id, collectionData) {
        try {
            const response = await fetch(`${API_BASE_URL}/collections/${id}`, {
                method: 'PUT',
                headers: getHeaders(),
                body: JSON.stringify(collectionData),
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            return data.data;
        } catch (error) {
            console.error(`Fehler beim Aktualisieren der Collection ${id}:`, error);
            throw error;
        }
    }

    // Collection löschen
    static async deleteCollection(id) {
        try {
            const response = await fetch(`${API_BASE_URL}/collections/${id}`, {
                method: 'DELETE',
                headers: getHeaders(),
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            return data;
        } catch (error) {
            console.error(`Fehler beim Löschen der Collection ${id}:`, error);
            throw error;
        }
    }
}

export const getAllCollectionsForUser = CollectionAPI.getAllCollectionsForUser;
export const getCollectionById = CollectionAPI.getCollectionById;
export const createCollection = CollectionAPI.createCollection;
export const updateCollection = CollectionAPI.updateCollection;
export const deleteCollection = CollectionAPI.deleteCollection;

export default CollectionAPI;
