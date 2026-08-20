import { apiGet, apiPut, apiDelete } from './apiClient';

/**
 * SICHERHEIT: Alle API-Requests nutzen den sicheren apiClient mit:
 * - HttpOnly Cookie Auth (credentials: 'include')
 * - CSRF-Token Protection
 * - Zentrales Error Handling
 */

class SettingsAPI {
  /**
   * Alle Settings abrufen
   */
  static async getAllSettings() {
    try {
      const response = await apiGet('/settings');
      return response?.data || [];
    } catch (error) {
      console.error('Fehler beim Abrufen der Settings:', error);
      throw error;
    }
  }

  /**
   * Setting nach Name abrufen
   */
  static async getSettingByName(name) {
    try {
      const response = await apiGet(`/settings/${name}`);
      return response?.data;
    } catch (error) {
      console.error(`Fehler beim Abrufen des Settings ${name}:`, error);
      throw error;
    }
  }

  /**
   * Setting aktualisieren
   */
  static async updateSetting(name, settingData) {
    try {
      const response = await apiPut(`/settings/${name}`, settingData);
      return response;
    } catch (error) {
      console.error(`Fehler beim Aktualisieren des Settings ${name}:`, error);
      throw error;
    }
  }

  /**
   * Setting löschen
   */
  static async deleteSetting(id) {
    try {
      const response = await apiDelete(`/settings/${id}`);
      return response;
    } catch (error) {
      console.error(`Fehler beim Löschen des Settings ${id}:`, error);
      throw error;
    }
  }
}

export const getAllSettings = SettingsAPI.getAllSettings.bind(SettingsAPI);
export const getSettingByName = SettingsAPI.getSettingByName.bind(SettingsAPI);
export const updateSetting = SettingsAPI.updateSetting.bind(SettingsAPI);
export const deleteSetting = SettingsAPI.deleteSetting.bind(SettingsAPI);

export default SettingsAPI;