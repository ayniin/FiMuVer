import React, { useState, useEffect } from 'react';
import './Admin.css';
import { getCurrentUser } from '../services/userapi';
import SettingsAPI from '../services/settingsapi';
import Header from '../components/Header';

const Admin = ({ user, onLogout, onNavigateBack }) => {
  const currentUser = getCurrentUser();
  const [settings, setSettings] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [successMessage, setSuccessMessage] = useState('');

  // Prüfe ob Benutzer Admin ist
  useEffect(() => {
    if (!currentUser?.is_admin) {
      setError('Du hast keine Berechtigung für diese Seite');
      setTimeout(() => onNavigateBack(), 2000);
      return;
    }
    
    loadSettings();
  }, []);

  const loadSettings = async () => {
    try {
      setLoading(true);
      const data = await SettingsAPI.getAllSettings();
      setSettings(data);
      setError('');
    } catch (err) {
      setError(err.message || 'Fehler beim Laden der Settings');
    } finally {
      setLoading(false);
    }
  };

  const handleToggleSetting = async (setting) => {
    try {
      setSuccessMessage('');
      const newValue = !setting.value;
      
      await SettingsAPI.updateSetting(setting.name, { value: newValue });
      
      // Update lokale State
      setSettings(settings.map(s => 
        s.name === setting.name ? { ...s, value: newValue } : s
      ));
      
      setSuccessMessage(`Setting '${setting.name}' aktualisiert`);
      setTimeout(() => setSuccessMessage(''), 3000);
    } catch (err) {
      setError(err.message || 'Fehler beim Aktualisieren des Settings');
    }
  };

  const handleDeleteSetting = async (setting) => {
    if (window.confirm(`Setting '${setting.name}' wirklich löschen?`)) {
      try {
        setSuccessMessage('');
        await SettingsAPI.deleteSetting(setting.id);
        setSettings(settings.filter(s => s.id !== setting.id));
        setSuccessMessage(`Setting '${setting.name}' gelöscht`);
        setTimeout(() => setSuccessMessage(''), 3000);
      } catch (err) {
        setError(err.message || 'Fehler beim Löschen des Settings');
      }
    }
  };

  return (
    <div className="admin-container">
      <Header user={currentUser} onLogout={onLogout} />

      <main className="admin-main">
        <div className="admin-header">
          <h1>Admin-Panel</h1>
          <button className="back-btn" onClick={onNavigateBack}>
            ← Zurück
          </button>
        </div>

        {error && <div className="error-message">{error}</div>}
        {successMessage && <div className="success-message">{successMessage}</div>}

        <div className="settings-section">
          <h2>Einstellungen</h2>
          
          {loading && <div className="loading">Wird geladen...</div>}

          {!loading && settings.length === 0 && (
            <div className="empty-state">
              <p>Keine Einstellungen vorhanden</p>
            </div>
          )}

          {!loading && settings.length > 0 && (
            <div className="settings-list">
              {settings.map((setting) => (
                <div key={setting.id} className="settings-item">
                  <div className="setting-info">
                    <div className="setting-name">{setting.name}</div>
                    <div className="setting-value">
                      Status: {setting.value ? '✅ Aktiviert' : '❌ Deaktiviert'}
                    </div>
                  </div>
                  <div className="setting-actions">
                    <button 
                      className="toggle-btn"
                      onClick={() => handleToggleSetting(setting)}
                    >
                      {setting.value ? 'Ausschalten' : 'Einschalten'}
                    </button>
                    <button 
                      className="delete-btn"
                      onClick={() => handleDeleteSetting(setting)}
                    >
                      Löschen
                    </button>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </main>

      <footer className="admin-footer">
        <p>&copy; 2026 FiMuVer - Admin Panel</p>
      </footer>
    </div>
  );
};

export default Admin;
