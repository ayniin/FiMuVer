import { apiGet, apiPost, apiPut, apiDelete } from './apiClient';

/**
 * SICHERHEIT: Alle API-Requests nutzen den sicheren apiClient mit:
 * - HttpOnly Cookie Auth (credentials: 'include')
 * - CSRF-Token Protection
 * - Zentrales Error Handling
 */

class CollectionAPI {
  /**
   * Alle Collections für den authentifizierten Benutzer abrufen
   */
  static async getAllCollectionsForUser() {
    try {
      const response = await apiGet('/collections');
      return response?.data || [];
    } catch (error) {
      console.error('Fehler beim Abrufen der Collections:', error);
      throw error;
    }
  }

  /**
   * Collection nach ID abrufen
   */
  static async getCollectionById(id) {
    try {
      const response = await apiGet(`/collections/${id}`);
      return response?.data;
    } catch (error) {
      console.error(`Fehler beim Abrufen der Collection ${id}:`, error);
      throw error;
    }
  }

  /**
   * Neue Collection erstellen
   */
  static async createCollection(collectionData) {
    try {
      const response = await apiPost('/collections', collectionData);
      return response?.data;
    } catch (error) {
      console.error('Fehler beim Erstellen der Collection:', error);
      throw error;
    }
  }

  /**
   * Collection aktualisieren
   */
  static async updateCollection(id, collectionData) {
    try {
      const response = await apiPut(`/collections/${id}`, collectionData);
      return response?.data;
    } catch (error) {
      console.error(`Fehler beim Aktualisieren der Collection ${id}:`, error);
      throw error;
    }
  }

  /**
   * Collection löschen
   */
  static async deleteCollection(id) {
    try {
      const response = await apiDelete(`/collections/${id}`);
      return response;
    } catch (error) {
      console.error(`Fehler beim Löschen der Collection ${id}:`, error);
      throw error;
    }
  }
}

export const getAllCollectionsForUser = CollectionAPI.getAllCollectionsForUser.bind(CollectionAPI);
export const getCollectionById = CollectionAPI.getCollectionById.bind(CollectionAPI);
export const createCollection = CollectionAPI.createCollection.bind(CollectionAPI);
export const updateCollection = CollectionAPI.updateCollection.bind(CollectionAPI);
export const deleteCollection = CollectionAPI.deleteCollection.bind(CollectionAPI);

export default CollectionAPI;
