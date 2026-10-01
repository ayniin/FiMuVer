import React, { useState } from 'react';
import { MEDIA_TYPE_LABELS, CONDITION_LABELS } from '../types';
import CollectionAPI from '../services/collection';
import MediaForm from '../components/MediaForm';

const ItemDetailPage = ({ item, onNavigateBack }) => {
  const [currentItem, setCurrentItem] = useState(item);
  const [editing, setEditing] = useState(false);
  const [error, setError] = useState('');

  const handleUpdate = async (formData) => {
    try {
      const updated = await CollectionAPI.updateItem(currentItem.collection_id, currentItem.id, formData);
      setCurrentItem(updated);
      setEditing(false);
    } catch (err) {
      setError(err.message || 'Fehler beim Aktualisieren des Eintrags');
    }
  };

  const handleDelete = async () => {
    if (!window.confirm('Diesen Eintrag wirklich löschen?')) return;
    try {
      await CollectionAPI.deleteItem(currentItem.collection_id, currentItem.id);
      onNavigateBack();
    } catch (err) {
      setError(err.message || 'Fehler beim Löschen des Eintrags');
    }
  };

  if (editing) {
    return (
      <div className="item-detail-page">
        <div className="item-detail-header">
          <button className="btn-back" onClick={() => setEditing(false)}>
            ← Abbrechen
          </button>
          <h1>{currentItem.title} bearbeiten</h1>
        </div>
        {error && <div className="error-message">{error}</div>}
        <MediaForm media={currentItem} onSubmit={handleUpdate} onCancel={() => setEditing(false)} />
      </div>
    );
  }

  return (
    <div className="item-detail-page">
      <div className="item-detail-header">
        <button className="btn-back" onClick={onNavigateBack}>
          ← Zurück
        </button>
        <h1>{currentItem.title}</h1>
        <div className="item-detail-actions">
          <button className="btn-edit" onClick={() => setEditing(true)}>Bearbeiten</button>
          <button className="btn-delete" onClick={handleDelete}>Löschen</button>
        </div>
      </div>

      {error && <div className="error-message">{error}</div>}

      <div className="item-detail-body">
        {currentItem.media_type && <p><strong>Typ:</strong> {MEDIA_TYPE_LABELS[currentItem.media_type]}</p>}
        {currentItem.artist && <p><strong>Künstler:</strong> {currentItem.artist}</p>}
        {currentItem.director && <p><strong>Regisseur:</strong> {currentItem.director}</p>}
        {currentItem.year && <p><strong>Jahr:</strong> {currentItem.year}</p>}
        {currentItem.genre && <p><strong>Genre:</strong> {currentItem.genre}</p>}
        {currentItem.condition && <p><strong>Zustand:</strong> {CONDITION_LABELS[currentItem.condition]}</p>}
        {currentItem.location && <p><strong>Lagerort:</strong> {currentItem.location}</p>}
        {currentItem.description && <p><strong>Beschreibung:</strong> {currentItem.description}</p>}
      </div>
    </div>
  );
};

export default ItemDetailPage;
