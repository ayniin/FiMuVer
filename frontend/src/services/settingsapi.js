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

class SettingsAPI {
      // ========== SETTINGS ENDPOINTS ==========

  // Alle Settings abrufen
  static async getAllSettings() {
    try {
      const response = await fetch(`${API_BASE_URL}/settings`, {
        headers: getHeaders(),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      return data.data || [];
    } catch (error) {
      console.error('Fehler beim Abrufen der Settings:', error);
      throw error;
    }
  }

  // Setting nach Name abrufen
  static async getSettingByName(name) {
    try {
      const response = await fetch(`${API_BASE_URL}/settings/${name}`, {
        headers: getHeaders(),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      return data.data;
    } catch (error) {
      console.error(`Fehler beim Abrufen des Settings ${name}:`, error);
      throw error;
    }
  }

  // Setting aktualisieren
  static async updateSetting(name, settingData) {
    try {
      const response = await fetch(`${API_BASE_URL}/settings/${name}`, {
        method: 'PUT',
        headers: getHeaders(),
        body: JSON.stringify(settingData),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      return data;
    } catch (error) {
      console.error(`Fehler beim Aktualisieren des Settings ${name}:`, error);
      throw error;
    }
  }

  // Setting löschen
  static async deleteSetting(id) {
    try {
      const response = await fetch(`${API_BASE_URL}/settings/${id}`, {
        method: 'DELETE',
        headers: getHeaders(),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      return data;
    } catch (error) {
      console.error(`Fehler beim Löschen des Settings ${id}:`, error);
      throw error;
    }
  }
}


// Exportiere die Funktionen als benannte Exports für die Kompatibilität mit Auth.jsx
export const getAllSettings = SettingsAPI.getAllSettings;
export const getSettingByName = SettingsAPI.getSettingByName;
export const updateSetting = SettingsAPI.updateSetting;
export const deleteSetting = SettingsAPI.deleteSetting;

export default SettingsAPI;