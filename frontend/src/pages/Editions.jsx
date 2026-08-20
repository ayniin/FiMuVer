import React, { useState, useEffect } from 'react';
import EditionAPI from '../services/editionapi';
import './Admin.css';

const Editions = () => {
  const [editions, setEditions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => { loadEditions(); }, []);

  const loadEditions = async () => {
    try {
      setLoading(true);
      const data = await EditionAPI.getAllEditions();
      setEditions(data || []);
      setError('');
    } catch (err) {
      setError(err.message || 'Fehler beim Laden der Editionen');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="settings-section">
      <h2>Editionen</h2>
      {loading && <div className="loading">Wird geladen...</div>}
      {error && <div className="error-message">{error}</div>}
      {!loading && editions.length === 0 && <div className="empty-state"><p>Keine Editionen vorhanden</p></div>}
      {!loading && editions.length > 0 && (
        <div className="settings-list">
          {editions.map(e => (
            <div key={e.id} className="settings-item">
              <div className="setting-info">
                <div className="setting-name">{e.title || e.name}</div>
                <div className="setting-value">{e.description || ''}</div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default Editions;
