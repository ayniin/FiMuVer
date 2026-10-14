import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { FiPlus, FiEdit2, FiTrash2, FiArrowRight, FiPackage } from 'react-icons/fi';
import CollectionAPI from '../services/collection';
import CollectionModal from '../components/CollectionModal';
import { useFeedback } from '../feedback/context';
import './Landing.css';

const Landing = () => {
  const { toast, confirm } = useFeedback();
  const [collections, setCollections] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  // null = geschlossen, {} = neu anlegen, {collection} = bearbeiten
  const [modal, setModal] = useState(null);

  useEffect(() => {
    const loadCollections = async () => {
      try {
        const data = await CollectionAPI.getAllCollectionsForUser();
        setCollections(data || []);
      } catch (err) {
        setError(err.message || 'Die Sammlungen konnten nicht geladen werden.');
      } finally {
        setLoading(false);
      }
    };
    loadCollections();
  }, []);

  const handleSaved = (saved) => {
    const exists = collections.some((c) => c.id === saved.id);
    setCollections(exists
      ? collections.map((c) => (c.id === saved.id ? saved : c))
      : [...collections, saved]);
    toast(exists ? 'Sammlung gespeichert' : `Sammlung „${saved.name}" erstellt`);
  };

  const handleDelete = async (collection) => {
    const ok = await confirm({
      title: 'Sammlung löschen?',
      message: `„${collection.name}" und alle enthaltenen Einträge werden dauerhaft gelöscht.`,
      confirmLabel: 'Löschen',
      danger: true,
    });
    if (!ok) return;

    try {
      await CollectionAPI.deleteCollection(collection.id);
      setCollections(collections.filter((c) => c.id !== collection.id));
      toast('Sammlung gelöscht');
    } catch (err) {
      toast(err.message || 'Die Sammlung konnte nicht gelöscht werden.', 'error');
    }
  };

  return (
    <>
      <div className="page-head">
        <div>
          <h1>Deine Sammlungen</h1>
          <p>Verwalte deine Filme, Serien und Musik an einem Ort.</p>
        </div>
        {collections.length > 0 && (
          <button className="btn btn--primary" onClick={() => setModal({})}>
            <FiPlus size={18} />
            Neue Sammlung
          </button>
        )}
      </div>

      {error && <div className="alert alert--error" role="alert">{error}</div>}

      {loading && (
        <div className="collections-grid" aria-busy="true" aria-label="Wird geladen">
          {[0, 1, 2].map((i) => <div key={i} className="skeleton collection-skeleton" />)}
        </div>
      )}

      {!loading && !error && collections.length === 0 && (
        <div className="card empty-state">
          <FiPackage size={56} />
          <h2>Noch keine Sammlungen</h2>
          <p>Erstelle deine erste Sammlung, um Medien zu verwalten.</p>
          <button className="btn btn--primary btn--lg" onClick={() => setModal({})}>
            <FiPlus size={18} />
            Erste Sammlung erstellen
          </button>
        </div>
      )}

      {!loading && collections.length > 0 && (
        <div className="collections-grid">
          {collections.map((collection) => (
            <article key={collection.id} className="card collection-card">
              <div className="collection-card-head">
                <h3>
                  <Link to={`/collections/${collection.id}`} className="collection-card-link">
                    {collection.name}
                  </Link>
                </h3>
                <div className="collection-card-actions">
                  <button
                    className="btn btn--ghost btn--icon"
                    onClick={() => setModal({ collection })}
                    aria-label={`„${collection.name}" bearbeiten`}
                    title="Bearbeiten"
                  >
                    <FiEdit2 size={17} />
                  </button>
                  <button
                    className="btn btn--ghost btn--icon btn--danger-hover"
                    onClick={() => handleDelete(collection)}
                    aria-label={`„${collection.name}" löschen`}
                    title="Löschen"
                  >
                    <FiTrash2 size={17} />
                  </button>
                </div>
              </div>
              <p className="collection-card-description">
                {collection.description || 'Keine Beschreibung'}
              </p>
              <div className="collection-card-foot">
                <span>
                  {collection.items?.length || 0} {collection.items?.length === 1 ? 'Eintrag' : 'Einträge'}
                </span>
                <span className="collection-card-open" aria-hidden="true">
                  Öffnen <FiArrowRight size={15} />
                </span>
              </div>
            </article>
          ))}
        </div>
      )}

      {modal && (
        <CollectionModal
          open
          collection={modal.collection}
          onClose={() => setModal(null)}
          onSaved={handleSaved}
        />
      )}
    </>
  );
};

export default Landing;
