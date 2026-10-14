import React from 'react';
import { Link } from 'react-router-dom';
import { FiTrash2 } from 'react-icons/fi';
import { MEDIA_TYPE_LABELS } from '../types';
import MediaCover from './MediaCover';
import './MediaCard.css';

export default function MediaCard({ media, to, onDelete }) {
  const creator = media.artist || media.director;
  const meta = [creator, media.year || null].filter(Boolean).join(' · ');

  return (
    <article className="card media-card">
      <MediaCover media={media} className="media-card-cover" />

      <div className="media-card-body">
        {media.media_type && <span className="badge">{MEDIA_TYPE_LABELS[media.media_type]}</span>}
        <h3>
          <Link to={to} className="media-card-link">{media.title}</Link>
        </h3>
        {meta && <p className="media-card-meta">{meta}</p>}
      </div>

      <button
        className="btn btn--icon media-card-delete"
        onClick={onDelete}
        aria-label={`„${media.title}" löschen`}
        title="Löschen"
      >
        <FiTrash2 size={16} />
      </button>
    </article>
  );
}
