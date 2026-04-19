import React, { useState, useEffect } from 'react';
import { useMedia } from '../hooks/useMedia';
import MediaCard from '../components/MediaCard';
import MediaForm from '../components/MediaForm';
import FilterBar from '../components/FilterBar';
import './Dashboard.css';

export default function Dashboard() {
  const { media, loading, error, fetchAllMedia, addMedia, updateMedia, deleteMedia, searchMedia } = useMedia();
  const [showForm, setShowForm] = useState(false);
  const [editingMedia, setEditingMedia] = useState(null);
  const [currentFilter, setCurrentFilter] = useState(null);

  useEffect(() => {
    fetchAllMedia();
  }, []);

  const handleAddClick = () => {
    setEditingMedia(null);
    setShowForm(true);
  };

  const handleFormSubmit = async (formData) => {
    try {
      if (editingMedia) {
        await updateMedia(editingMedia.id, formData);
      } else {
        await addMedia(formData);
      }
      setShowForm(false);
      setEditingMedia(null);
    } catch (err) {
      console.error('Fehler beim Speichern:', err);
    }
  };

  const handleFormCancel = () => {
    setShowForm(false);
    setEditingMedia(null);
  };

  const handleEdit = (mediaItem) => {
    setEditingMedia(mediaItem);
    setShowForm(true);
  };

  const handleDelete = async (id) => {
    if (window.confirm('Dieses Medium wirklich löschen?')) {
      try {
        await deleteMedia(id);
      } catch (err) {
        console.error('Fehler beim Löschen:', err);
      }
    }
  };

  const handleFilter = async (type) => {
    setCurrentFilter(type);
    if (type) {
      await fetchAllMedia(type);
    } else {
      await fetchAllMedia();
    }
  };

  const handleSearch = async (query) => {
    if (query.trim()) {
      await searchMedia(query);
    } else {
      await fetchAllMedia(currentFilter);
    }
  };

  return (
    <div className="dashboard">
      <header className="dashboard-header">
        <h1>🎬 FiMuVer - Deine Medienverwaltung</h1>
        <p>Verwalte deine Blurays, DVDs, Vinyls und Tapes</p>
      </header>

      <div className="dashboard-content">
        <FilterBar onFilter={handleFilter} onSearch={handleSearch} />

        {error && <div className="error-message">Fehler: {error}</div>}

        {!showForm ? (
          <button className="btn-add-media" onClick={handleAddClick}>
            + Neues Medium hinzufügen
          </button>
        ) : null}

        {showForm && (
          <div className="form-container">
            <h2>{editingMedia ? 'Medium bearbeiten' : 'Neues Medium hinzufügen'}</h2>
            <MediaForm
              media={editingMedia}
              onSubmit={handleFormSubmit}
              onCancel={handleFormCancel}
            />
          </div>
        )}

        {loading ? (
          <div className="loading">Lädt...</div>
        ) : media && media.length > 0 ? (
          <div className="media-list">
            <h2>Medien ({media.length})</h2>
            {media.map((item) => (
              <MediaCard
                key={item.id}
                media={item}
                onEdit={handleEdit}
                onDelete={handleDelete}
              />
            ))}
          </div>
        ) : (
          <div className="empty-state">
            <p>Noch keine Medien vorhanden. Füge eines hinzu!</p>
          </div>
        )}
      </div>
    </div>
  );
}

