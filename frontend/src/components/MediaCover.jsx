import React, { useState } from 'react';
import { FiFilm, FiMusic } from 'react-icons/fi';
import { isMusicType } from '../types';

/** Cover-Bild mit Platzhalter, falls keine oder eine kaputte Bild-URL vorliegt. */
const MediaCover = ({ media, className = '' }) => {
  const [failed, setFailed] = useState(false);
  const Icon = isMusicType(media.media_type) ? FiMusic : FiFilm;

  return (
    <div className={`media-cover ${className}`}>
      {media.image_url && !failed ? (
        <img src={media.image_url} alt="" loading="lazy" onError={() => setFailed(true)} />
      ) : (
        <Icon size={40} aria-hidden="true" />
      )}
    </div>
  );
};

export default MediaCover;
