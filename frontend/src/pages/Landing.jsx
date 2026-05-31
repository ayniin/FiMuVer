import React, { useState, useEffect } from 'react';
import { FiPlus, FiEdit2, FiTrash2, FiArrowRight, FiPackage } from 'react-icons/fi';
import './Landing.css';
import { getCurrentUser } from '../services/userapi';
import CollectionAPI from '../services/collection';
import Header from '../components/Header';
import CreateCollectionModal from '../components/CreateCollectionModal';

const hasAuthToken = () => {
  try {
    return Boolean(sessionStorage.getItem('auth_token'));
  } catch {
    return false;
  }
};

const Landing = ({ user, onLogout, onNavigateToAdmin }) => {
  const currentUser = getCurrentUser();
  const [collections, setCollections] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [showModal, setShowModal] = useState(false);
  const [successMessage, setSuccessMessage] = useState('');

  useEffect(() => {
    if (hasAuthToken()) {
      loadCollections();
      return;
    }

    setCollections([]);
    setLoading(false);
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

  const handleEditCollection = async (id) => {
    try {
      // TODO: Hardcoded Name
        await CollectionAPI.updateCollection(id, { name: 'Neuer Name' });
    } catch (err) {
        setError(err.message || 'Fehler beim Aktualisieren der Collection');
    }
  };

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
            aria-label="Neue Collection erstellen"
          >
            <FiPlus size={20} style={{ marginRight: '8px', verticalAlign: 'middle' }} />
            Neue Collection
          </button>
        </div>

        {error && <div className="error-message">{error}</div>}
        {successMessage && <div className="success-message">{successMessage}</div>}

        {loading && <div className="loading">Wird geladen...</div>}

        {!loading && collections.length === 0 && (
          <div className="empty-state">
            <div className="empty-icon">
              <FiPackage size={64} />
            </div>
            <h2>Noch keine Sammlungen</h2>
            <p>Erstelle deine erste Collection um Medien zu verwalten</p>
            <button 
              className="btn-create-collection btn-large"
              onClick={() => setShowModal(true)}
              aria-label="Erste Collection erstellen"
            >
              <FiPlus size={20} style={{ marginRight: '8px', verticalAlign: 'middle' }} />
              Erste Collection erstellen
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
                    className="btn-edit"
                    onClick={() => handleEditCollection(collection.id)}
                    title="Editieren"
                    aria-label="Collection bearbeiten"
                  >
                    <FiEdit2 size={18} />
                  </button>
                  <button 
                    className="btn-delete"
                    onClick={() => handleDeleteCollection(collection.id, collection.name)}
                    title="Löschen"
                    aria-label="Collection löschen"
                  >
                    <FiTrash2 size={18} />
                  </button>
                </div>
                <p className="collection-description">
                  {collection.description || 'Keine Beschreibung'}
                </p>
                <div className="collection-footer">
                  <span className="collection-items">
                    {collection.items?.length || 0} Einträge
                  </span>
                  <button className="btn-open" aria-label="Collection öffnen">
                    Öffnen <FiArrowRight size={16} style={{ marginLeft: '6px', verticalAlign: 'middle' }} />
                  </button>
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
