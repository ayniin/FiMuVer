import React from 'react';
import { MEDIA_TYPE_LABELS, CONDITION_LABELS } from '../types';

const ItemDetailPage = ({ item, onNavigateBack }) => {
  return (
    <div className="item-detail-page">
      <div className="item-detail-header">
        <button className="btn-back" onClick={onNavigateBack}>
          ← Zurück
        </button>
        <h1>{item.title}</h1>
      </div>

      <div className="item-detail-body">
        {item.image_url && <img className="item-detail-poster" src={item.image_url} alt={item.title} />}
        {item.media_type && <p><strong>Typ:</strong> {MEDIA_TYPE_LABELS[item.media_type]}</p>}
        {item.edition && <p><strong>Edition:</strong> {item.edition}</p>}
        {item.imdb_id && (
          <p><strong>IMDb:</strong> <a href={`https://www.imdb.com/title/${item.imdb_id}/`} target="_blank" rel="noreferrer">{item.imdb_id}</a></p>
        )}
        {item.artist && <p><strong>Künstler:</strong> {item.artist}</p>}
        {item.director && <p><strong>Regisseur:</strong> {item.director}</p>}
        {item.year && <p><strong>Jahr:</strong> {item.year}</p>}
        {item.genre && <p><strong>Genre:</strong> {item.genre}</p>}
        {item.condition && <p><strong>Zustand:</strong> {CONDITION_LABELS[item.condition]}</p>}
        {item.location && <p><strong>Lagerort:</strong> {item.location}</p>}
        {item.description && <p><strong>Beschreibung:</strong> {item.description}</p>}
      </div>
    </div>
  );
};

export default ItemDetailPage;
