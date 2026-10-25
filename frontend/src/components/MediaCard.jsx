import React from 'react';
import { MEDIA_TYPE_LABELS, CONDITION_LABELS } from '../types';
import './MediaCard.css';

export default function MediaCard({ media, onOpen, onEdit, onDelete }) {
  return (
    <div className="media-card">
      {media.image_url && (
        <img className="media-card-poster" src={media.image_url} alt={media.title} onClick={onOpen} />
      )}
      <div className="media-card-header">
        <h3 onClick={onOpen} style={{ cursor: onOpen ? 'pointer' : undefined }}>{media.title}</h3>
        <span className="media-type-badge">{MEDIA_TYPE_LABELS[media.media_type]}</span>
      </div>

      <div className="media-card-body">
        {media.edition && <p><strong>Edition:</strong> {media.edition}</p>}
        {media.artist && <p><strong>Künstler:</strong> {media.artist}</p>}
        {media.director && <p><strong>Regisseur:</strong> {media.director}</p>}
        {media.year && <p><strong>Jahr:</strong> {media.year}</p>}
        {media.genre && <p><strong>Genre:</strong> {media.genre}</p>}
        {media.condition && <p><strong>Zustand:</strong> {CONDITION_LABELS[media.condition]}</p>}
        {media.location && <p><strong>Lagerort:</strong> {media.location}</p>}
        {media.description && <p><strong>Beschreibung:</strong> {media.description}</p>}
      </div>

      <div className="media-card-footer">
        <button className="btn-edit" onClick={() => onEdit(media)}>Bearbeiten</button>
        <button className="btn-delete" onClick={() => onDelete(media.id)}>Löschen</button>
      </div>
    </div>
  );
}

