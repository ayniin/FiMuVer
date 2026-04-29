const API_BASE_URL = 'http://localhost:8080/api';

class MediaAPI {
  // Alle Medien abrufen (optional gefiltert nach Typ)
  static async getAllMedia(mediaType = null) {
    try {
      const url = mediaType
        ? `${API_BASE_URL}/v1/media?type=${mediaType}`
        : `${API_BASE_URL}/v1/media`;

      const response = await fetch(url);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      return data.data || [];
    } catch (error) {
      console.error('Fehler beim Abrufen der Medien:', error);
      throw error;
    }
  }

  // Medium nach ID abrufen
  static async getMediaById(id) {
    try {
      const response = await fetch(`${API_BASE_URL}/v1/media/${id}`);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      return data.data;
    } catch (error) {
      console.error(`Fehler beim Abrufen des Mediums ${id}:`, error);
      throw error;
    }
  }

  // Neues Medium erstellen
  static async createMedia(media) {
    try {
      const response = await fetch(`${API_BASE_URL}/v1/media`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(media),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      return data.data;
    } catch (error) {
      console.error('Fehler beim Erstellen des Mediums:', error);
      throw error;
    }
  }

  // Medium aktualisieren
  static async updateMedia(id, media) {
    try {
      const response = await fetch(`${API_BASE_URL}/v1/media/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(media),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return await response.json();
    } catch (error) {
      console.error(`Fehler beim Aktualisieren des Mediums ${id}:`, error);
      throw error;
    }
  }

  // Medium löschen
  static async deleteMedia(id) {
    try {
      const response = await fetch(`${API_BASE_URL}/v1/media/${id}`, {
        method: 'DELETE',
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return await response.json();
    } catch (error) {
      console.error(`Fehler beim Löschen des Mediums ${id}:`, error);
      throw error;
    }
  }

  // Nach Medien suchen
  static async searchMedia(query) {
    try {
      const response = await fetch(`${API_BASE_URL}/v1/media/search?q=${query}`);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      return data.data || [];
    } catch (error) {
      console.error('Fehler bei der Mediensuche:', error);
      throw error;
    }
  }
}



export const logout = () => {
  localStorage.removeItem('user');
};

export const getCurrentUser = () => {
  return JSON.parse(localStorage.getItem('user'));
};

export default MediaAPI;

