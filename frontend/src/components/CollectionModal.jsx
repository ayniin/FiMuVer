import React, { useState } from 'react';
import Modal from './Modal';
import CollectionAPI from '../services/collection';

/**
 * Dialog zum Anlegen (ohne `collection`) oder Bearbeiten (mit `collection`) einer Sammlung.
 * Wird bei jedem Öffnen neu gemountet (siehe `key` beim Aufrufer), damit die Felder frisch befüllt sind.
 */
const CollectionModal = ({ open, collection = null, onClose, onSaved }) => {
  const isEdit = Boolean(collection);
  const [name, setName] = useState(collection?.name || '');
  const [description, setDescription] = useState(collection?.description || '');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!name.trim()) {
      setError('Bitte gib einen Namen ein.');
      return;
    }

    setLoading(true);
    setError('');
    const payload = { name: name.trim(), description: description.trim() };

    try {
      const saved = isEdit
        ? await CollectionAPI.updateCollection(collection.id, payload)
        : await CollectionAPI.createCollection(payload);
      onSaved({ ...collection, ...payload, ...saved });
      onClose();
    } catch (err) {
      setError(err.message || 'Die Sammlung konnte nicht gespeichert werden.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Modal open={open} onClose={onClose} title={isEdit ? 'Sammlung bearbeiten' : 'Neue Sammlung'}>
      <form onSubmit={handleSubmit}>
        <div className="field">
          <label htmlFor="collection-name">Name</label>
          <input
            className="input"
            type="text"
            id="collection-name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="z.B. Meine Lieblingsfilme"
            disabled={loading}
            required
            autoFocus
          />
        </div>

        <div className="field">
          <label htmlFor="collection-description">
            Beschreibung <span className="field-hint">(optional)</span>
          </label>
          <textarea
            className="input"
            id="collection-description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            placeholder="z.B. Die besten Filme aller Zeiten…"
            rows="3"
            disabled={loading}
          />
        </div>

        {error && <div className="alert alert--error" role="alert">{error}</div>}

        <div className="modal-actions">
          <button type="button" className="btn btn--secondary" onClick={onClose} disabled={loading}>
            Abbrechen
          </button>
          <button type="submit" className="btn btn--primary" disabled={loading}>
            {loading ? 'Wird gespeichert…' : isEdit ? 'Speichern' : 'Sammlung erstellen'}
          </button>
        </div>
      </form>
    </Modal>
  );
};

export default CollectionModal;
