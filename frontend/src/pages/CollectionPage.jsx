import React, { useState, useEffect, useMemo } from 'react';
import { Link, useParams } from 'react-router-dom';
import { FiArrowLeft, FiEdit2, FiInbox, FiPlus } from 'react-icons/fi';
import CollectionAPI from '../services/collection';
import MediaCard from '../components/MediaCard';
import MediaForm from '../components/MediaForm';
import Modal from '../components/Modal';
import CollectionModal from '../components/CollectionModal';
import FilterBar from '../components/FilterBar';
import { useFeedback } from '../feedback/context';

const compareItems = {
  title: (a, b) => a.title.localeCompare(b.title, 'de'),
  year_desc: (a, b) => (b.year || 0) - (a.year || 0),
  year_asc: (a, b) => (a.year || 0) - (b.year || 0),
  created_desc: (a, b) => new Date(b.created_at) - new Date(a.created_at),
};

const CollectionPage = () => {
  const { collectionId } = useParams();
  const { toast, confirm } = useFeedback();
  const [collection, setCollection] = useState(null);
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [showAddForm, setShowAddForm] = useState(false);
  const [showEditCollection, setShowEditCollection] = useState(false);

  const [search, setSearch] = useState('');
  const [typeFilter, setTypeFilter] = useState(null);
  const [sort, setSort] = useState('title');

  useEffect(() => {
    const loadCollection = async () => {
      setLoading(true);
      setError('');
      try {
        const data = await CollectionAPI.getCollectionById(collectionId);
        setCollection(data);
        setItems(data?.items || []);
      } catch (err) {
        setError(err.message || 'Die Sammlung konnte nicht geladen werden.');
      } finally {
        setLoading(false);
      }
    };
    loadCollection();
  }, [collectionId]);

  const visibleItems = useMemo(() => {
    const query = search.trim().toLowerCase();
    return items
      .filter((item) => !typeFilter || item.media_type === typeFilter)
      .filter((item) => !query || [item.title, item.artist, item.director]
        .some((value) => value?.toLowerCase().includes(query)))
      .sort(compareItems[sort]);
  }, [items, search, typeFilter, sort]);

  const handleAddItem = async (formData) => {
    const newItem = await CollectionAPI.addItem(collectionId, formData);
    setItems([...items, newItem]);
    setShowAddForm(false);
    toast(`„${newItem.title}" hinzugefügt`);
  };

  const handleDeleteItem = async (item) => {
    const ok = await confirm({
      title: 'Eintrag löschen?',
      message: `„${item.title}" wird dauerhaft aus der Sammlung entfernt.`,
      confirmLabel: 'Löschen',
      danger: true,
    });
    if (!ok) return;

    try {
      await CollectionAPI.deleteItem(collectionId, item.id);
      setItems(items.filter((i) => i.id !== item.id));
      toast('Eintrag gelöscht');
    } catch (err) {
      toast(err.message || 'Der Eintrag konnte nicht gelöscht werden.', 'error');
    }
  };

  const hasFilter = search.trim() || typeFilter;

  return (
    <>
      <Link to="/" className="breadcrumb">
        <FiArrowLeft size={15} /> Alle Sammlungen
      </Link>

      {error && <div className="alert alert--error" role="alert">{error}</div>}

      {loading && (
        <>
          <div className="skeleton skeleton--title" />
          <div className="media-grid" aria-busy="true" aria-label="Wird geladen">
            {[0, 1, 2, 3, 4].map((i) => <div key={i} className="skeleton skeleton--media" />)}
          </div>
        </>
      )}

      {!loading && collection && (
        <>
          <div className="page-head">
            <div>
              <h1>{collection.name}</h1>
              {collection.description && <p>{collection.description}</p>}
            </div>
            <div className="page-head-actions">
              <button className="btn btn--secondary" onClick={() => setShowEditCollection(true)}>
                <FiEdit2 size={16} />
                Bearbeiten
              </button>
              <button className="btn btn--primary" onClick={() => setShowAddForm(true)}>
                <FiPlus size={18} />
                Eintrag hinzufügen
              </button>
            </div>
          </div>

          {items.length > 0 && (
            <FilterBar
              search={search}
              onSearch={setSearch}
              type={typeFilter}
              onType={setTypeFilter}
              sort={sort}
              onSort={setSort}
            />
          )}

          {items.length === 0 && (
            <div className="card empty-state">
              <FiInbox size={56} />
              <h2>Noch leer</h2>
              <p>Diese Sammlung enthält noch keine Einträge.</p>
              <button className="btn btn--primary btn--lg" onClick={() => setShowAddForm(true)}>
                <FiPlus size={18} />
                Ersten Eintrag hinzufügen
              </button>
            </div>
          )}

          {items.length > 0 && visibleItems.length === 0 && hasFilter && (
            <div className="empty-state">
              <p>Keine Einträge passen zu deiner Suche.</p>
              <button className="btn btn--secondary" onClick={() => { setSearch(''); setTypeFilter(null); }}>
                Filter zurücksetzen
              </button>
            </div>
          )}

          {visibleItems.length > 0 && (
            <div className="media-grid">
              {visibleItems.map((item) => (
                <MediaCard
                  key={item.id}
                  media={item}
                  to={`/collections/${collectionId}/items/${item.id}`}
                  onDelete={() => handleDeleteItem(item)}
                />
              ))}
            </div>
          )}
        </>
      )}

      <Modal open={showAddForm} onClose={() => setShowAddForm(false)} title="Eintrag hinzufügen" size="lg">
        <MediaForm onSubmit={handleAddItem} onCancel={() => setShowAddForm(false)} />
      </Modal>

      {showEditCollection && (
        <CollectionModal
          open
          collection={collection}
          onClose={() => setShowEditCollection(false)}
          onSaved={(saved) => {
            setCollection(saved);
            toast('Sammlung gespeichert');
          }}
        />
      )}
    </>
  );
};

export default CollectionPage;
