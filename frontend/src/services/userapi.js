const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';

class UserAPI {
    static async registerUser(username, email, password) {
        if (!username || !email || !password) {
            throw new Error('Alle Felder müssen ausgefüllt werden.');
        }

        if (password.length < 12) {
            throw new Error('Das Passwort muss mindestens 12 Zeichen lang sein.');
        }

        const body = JSON.stringify({ username, email, password });

        try {
            const response = await fetch(`${API_BASE_URL}/users/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: body,
            });

            // Wenn erfolgreich aber leer (z.B. 201 oder 204)
            if (response.ok && response.headers.get('Content-Length') === '0') {
                return true;
            }

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.error || `HTTP error! status: ${response.status}`);
            }
            const data = await response.json();

            console.log('Registrierungsantwort:', data);

            this.saveUser(data.data);

            return data;
        } catch (error) {
            console.error('Fehler bei der Registrierung:', error);
            throw error;
        }
    }

    static async loginUser(email, password) {
        try {
            const response = await fetch(`${API_BASE_URL}/users/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email, password }),
            });
            
            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.error || `HTTP error! status: ${response.status}`);
            }   
            const data = await response.json();
            
            console.log('Login-Antwort:', data);

            this.saveUser(data.data);
            
            return data;
        } catch (error) {
            console.error('Fehler beim Login:', error);
            throw error;
        }
    }

    static async saveUser(data) {
        localStorage.setItem('user', JSON.stringify(data));
    }
}

export const getCurrentUser = () => {
  const userStr = localStorage.getItem('user');
  return userStr ? JSON.parse(userStr) : null;
};

export const logout = () => {
    localStorage.removeItem('user');
}

// Exportiere die Funktionen als benannte Exports für die Kompatibilität mit Auth.jsx
export const registerUser = UserAPI.registerUser;
export const loginUser = UserAPI.loginUser;

export default UserAPI;