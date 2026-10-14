import React, { useState, useEffect } from 'react';
import EditionAPI from '../services/editionapi';
import './Admin.css';

const Editions = () => {
  const [editions, setEditions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const loadEditions = async () => {
      try {
        const data = await EditionAPI.getAllEditions();
        setEditions(data || []);
      } catch (err) {
        setError(err.message || 'Die Editionen konnten nicht geladen werden.');
      } finally {
        setLoading(false);
      }
    };
    loadEditions();
  }, []);

  if (loading) return <div className="skeleton admin-skeleton" aria-busy="true" aria-label="Wird geladen" />;
  if (error) return <div className="alert alert--error" role="alert">{error}</div>;
  if (editions.length === 0) return <p className="empty-state">Keine Editionen vorhanden.</p>;

  return (
    <ul className="admin-list">
      {editions.map((e) => (
        <li key={e.id} className="admin-list-item">
          <div>
            <span className="admin-list-name">{e.title || e.name}</span>
            {e.description && <span className="admin-list-meta">{e.description}</span>}
          </div>
        </li>
      ))}
    </ul>
  );
};

export default Editions;
