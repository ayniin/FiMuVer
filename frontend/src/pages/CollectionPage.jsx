import React, { useState, useEffect } from 'react';
import CollectionAPI from '../services/collection';
import MediaCard from '../components/MediaCard';
import Header from '../components/Header';

const CollectionPage = ({ collection, onNavigateToItem, onNavigateBack, user, onLogout }) => {
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    loadItems();
  }, []);

  const loadItems = async () => {
    try {
      setLoading(true);
      setError('');
      const data = await CollectionAPI.getCollectionById(collection.id);
      setItems(data.items || []);
    } catch (err) {
      setError(err.message || 'Fehler beim Laden der Einträge');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="collection-page">
      <Header user={user} onLogout={onLogout} />

      <main className="collection-main">
        <div className="collection-page-header">
          <button className="btn-back" onClick={onNavigateBack}>
            ← Zurück
          </button>
          <div>
            <h1>{collection.name}</h1>
            {collection.description && <p>{collection.description}</p>}
          </div>
        </div>

        {error && <div className="error-message">{error}</div>}
        {loading && <div className="loading">Wird geladen...</div>}

        {!loading && items.length === 0 && (
          <div className="empty-state">
            <p>Diese Collection enthält noch keine Einträge.</p>
          </div>
        )}

        {!loading && items.length > 0 && (
          <div className="media-grid">
            {items.map((item) => (
              <MediaCard
                key={item.id}
                media={item}
                onEdit={() => onNavigateToItem(item)}
                onDelete={() => {}}
              />
            ))}
          </div>
        )}
      </main>
    </div>
  );
};

export default CollectionPage;
