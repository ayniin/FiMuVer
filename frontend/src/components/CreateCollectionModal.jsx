import React, { useState } from 'react';
import './CreateCollectionModal.css';
import CollectionAPI from '../services/collection';

const CreateCollectionModal = ({ isOpen, onClose, onCollectionCreated }) => {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    if (!name.trim()) {
      setError('Name ist erforderlich');
      return;
    }

    setLoading(true);
    setError('');

    try {
      const newCollection = await CollectionAPI.createCollection({
        name: name.trim(),
        description: description.trim(),
      });
      
      // Reset form
      setName('');
      setDescription('');
      
      // Callback an Landing.jsx
      onCollectionCreated(newCollection);
      onClose();
    } catch (err) {
      setError(err.message || 'Fehler beim Erstellen der Collection');
    } finally {
      setLoading(false);
    }
  };

  if (!isOpen) return null;

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <div className="modal-header">
          <h2>Neue Collection erstellen</h2>
          <button className="modal-close" onClick={onClose}>✕</button>
        </div>

        <form onSubmit={handleSubmit} className="modal-form">
          <div className="form-group">
            <label htmlFor="name">Name *</label>
            <input
              type="text"
              id="name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="z.B. Meine Lieblingsfilme"
              disabled={loading}
              autoFocus
            />
          </div>

          <div className="form-group">
            <label htmlFor="description">Beschreibung (optional)</label>
            <textarea
              id="description"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              placeholder="z.B. Die besten Filme aller Zeiten..."
              rows="4"
              disabled={loading}
            />
          </div>

          {error && <div className="error-message">{error}</div>}

          <div className="modal-actions">
            <button 
              type="button" 
              className="btn-cancel" 
              onClick={onClose}
              disabled={loading}
            >
              Abbrechen
            </button>
            <button 
              type="submit" 
              className="btn-create"
              disabled={loading}
            >
              {loading ? 'Wird erstellt...' : 'Collection erstellen'}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default CreateCollectionModal;
