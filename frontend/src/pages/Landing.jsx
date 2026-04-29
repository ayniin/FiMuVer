import React, { useState, useEffect } from 'react';
import './Landing.css';
import { getCurrentUser } from '../services/userapi';
import CollectionAPI from '../services/collection';
import Header from '../components/Header';
import CreateCollectionModal from '../components/CreateCollectionModal';

const Landing = ({ user, onLogout, onNavigateToAdmin }) => {
  const currentUser = getCurrentUser();
  const [collections, setCollections] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [showModal, setShowModal] = useState(false);
  const [successMessage, setSuccessMessage] = useState('');

  useEffect(() => {
    loadCollections();
  }, []);

  const loadCollections = async () => {
    try {
      setLoading(true);
      setError('');
      const data = await CollectionAPI.getAllCollectionsForUser();
      setCollections(data || []);
    } catch (err) {
      setError(err.message || 'Fehler beim Laden der Collections');
      setCollections([]);
    } finally {
      setLoading(false);
    }
  };

  const handleCollectionCreated = (newCollection) => {
    setCollections([...collections, newCollection]);
    setSuccessMessage(`Collection "${newCollection.name}" erstellt!`);
    setTimeout(() => setSuccessMessage(''), 3000);
  };

  const handleDeleteCollection = async (id, name) => {
    if (window.confirm(`Collection "${name}" wirklich löschen?`)) {
      try {
        await CollectionAPI.deleteCollection(id);
        setCollections(collections.filter(c => c.id !== id));
        setSuccessMessage(`Collection gelöscht!`);
        setTimeout(() => setSuccessMessage(''), 3000);
      } catch (err) {
        setError(err.message || 'Fehler beim Löschen der Collection');
      }
    }
  };

  console.log('Aktueller Benutzer in Landing:', currentUser);

  return (
    <div className="landing-container">
      <Header user={currentUser} onLogout={onLogout} onNavigateToAdmin={onNavigateToAdmin} />

      <main className="landing-main">
        <div className="collections-header">
          <div>
            <h1>Deine Sammlungen</h1>
            <p>Verwalte deine Mediensammlungen</p>
          </div>
          <button 
            className="btn-create-collection"
            onClick={() => setShowModal(true)}
          >
            ➕ Neue Collection
          </button>
        </div>

        {error && <div className="error-message">{error}</div>}
        {successMessage && <div className="success-message">{successMessage}</div>}

        {loading && <div className="loading">Wird geladen...</div>}

        {!loading && collections.length === 0 && (
          <div className="empty-state">
            <div className="empty-icon">📦</div>
            <h2>Noch keine Sammlungen</h2>
            <p>Erstelle deine erste Collection um Medien zu verwalten</p>
            <button 
              className="btn-create-collection btn-large"
              onClick={() => setShowModal(true)}
            >
              ➕ Erste Collection erstellen
            </button>
          </div>
        )}

        {!loading && collections.length > 0 && (
          <div className="collections-grid">
            {collections.map((collection) => (
              <div key={collection.id} className="collection-card">
                <div className="collection-header">
                  <h3>{collection.name}</h3>
                  <button 
                    className="btn-delete"
                    onClick={() => handleDeleteCollection(collection.id, collection.name)}
                    title="Löschen"
                  >
                    🗑️
                  </button>
                </div>
                <p className="collection-description">
                  {collection.description || 'Keine Beschreibung'}
                </p>
                <div className="collection-footer">
                  <span className="collection-items">
                    {collection.items?.length || 0} Einträge
                  </span>
                  <button className="btn-open">Öffnen →</button>
                </div>
              </div>
            ))}
          </div>
        )}
      </main>

      <CreateCollectionModal 
        isOpen={showModal}
        onClose={() => setShowModal(false)}
        onCollectionCreated={handleCollectionCreated}
      />

      <footer className="landing-footer">
        <p>&copy; 2026 FiMuVer - Medienverwaltung leicht gemacht</p>
      </footer>
    </div>
  );
};

export default Landing;
