import { apiPost, apiGet, apiDelete } from './apiClient';

/**
 * SICHERHEIT: Token wird NICHT im localStorage gespeichert!
 * - Token wird vom Backend in HttpOnly Cookie gespeichert
 * - Frontend sendet Cookie automatisch in jedem Request
 * - localStorage speichert nur öffentliche User-Metadaten
 */

class UserAPI {
  /**
   * Benutzer registrieren
   * Backend setzt HttpOnly Cookie mit Auth-Token
   */
  static async registerUser(username, email, password, inviteCode = '') {
    if (!username || !email || !password) {
      throw new Error('Alle Felder müssen ausgefüllt werden.');
    }

    if (password.length < 12) {
      throw new Error('Das Passwort muss mindestens 12 Zeichen lang sein.');
    }

    const payload = { username, email, password };
    if (inviteCode) {
      payload.invite_code = inviteCode;
    }

    try {
      const response = await apiPost('/users/register', payload);

      // Speichere nur User-Metadaten (KEIN TOKEN!)
      if (response?.data) {
        this.saveUserMetadata(response.data);
      }

      return response;
    } catch (error) {
      console.error('Fehler bei der Registrierung:', error);
      throw error;
    }
  }

  /**
   * Benutzer einloggen
   * Backend setzt HttpOnly Cookie mit Auth-Token
   */
  static async loginUser(email, password) {
    try {
      const response = await apiPost('/users/login', {
        email,
        password,
      });

      // Speichere nur User-Metadaten (KEIN TOKEN!)
      if (response?.data) {
        this.saveUserMetadata(response.data);
      }

      return response;
    } catch (error) {
      console.error('Fehler beim Login:', error);
      throw error;
    }
  }

  /**
   * Speichere öffentliche User-Metadaten und den JWT-Token für API-Calls.
   */
  static saveUserMetadata(data) {
    const metadata = {
      id: data.id,
      username: data.username,
      email: data.email,
      is_admin: data.is_admin,
    };
    localStorage.setItem('user_metadata', JSON.stringify(metadata));

    if (data.token) {
      sessionStorage.setItem('auth_token', data.token);
    }
  }

  /**
   * Logout - Cookie wird vom Backend gelöscht
   */
  static async logout() {
    try {
      await apiPost('/users/logout', {});
    } catch (error) {
      console.error('Fehler beim Logout:', error);
    } finally {
      // LocalStorage löschen
      localStorage.removeItem('user_metadata');
      sessionStorage.removeItem('auth_token');
    }
  }
}

/**
 * Gib aktuelle User-Metadaten zurück (KEIN TOKEN)
 */
export const getCurrentUser = () => {
  const userStr = localStorage.getItem('user_metadata');
  return userStr ? JSON.parse(userStr) : null;
};

/**
 * Logout
 */
export const logout = () => {
  return UserAPI.logout();
};

// Exportiere die Funktionen als benannte Exports für Kompatibilität
export const registerUser = UserAPI.registerUser.bind(UserAPI);
export const loginUser = UserAPI.loginUser.bind(UserAPI);

export default UserAPI;