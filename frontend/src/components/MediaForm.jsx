import React, { useState, useEffect } from 'react';
import { MEDIA_TYPES, MEDIA_TYPE_LABELS, CONDITIONS, CONDITION_LABELS, Media } from '../types';
import TVDBAPI from '../services/tvdb';
import TMDBAPI from '../services/tmdb';
import './MediaForm.css';

export default function MediaForm({ media = null, onSubmit, onCancel }) {
  const [formData, setFormData] = useState(
    media || new Media()
  );


  const [searchType, setSearchType] = useState('movie'); // 'movie' | 'series'
  const [searchQuery, setSearchQuery] = useState('');
  const [activeSource, setActiveSource] = useState('tmdb'); // 'tmdb' | 'tvdb'
  const [results, setResults] = useState({ tmdb: [], tvdb: [] });
  const [errors, setErrors] = useState({ tmdb: '', tvdb: '' });
  const [searchLoading, setSearchLoading] = useState(false);

  useEffect(() => {
    if (media) {
      setFormData(media);
    }
  }, [media]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: name === 'year' ? parseInt(value) : value,
    });
  };

  const handleSearch = async (e) => {
    e.preventDefault();
    if (!searchQuery.trim()) return;

    setSearchLoading(true);
    setErrors({ tmdb: '', tvdb: '' });
    setResults({ tmdb: [], tvdb: [] });

    const fetchTmdb = searchType === 'series'
      ? TMDBAPI.searchSeries(searchQuery)
      : TMDBAPI.searchMovies(searchQuery);
    const fetchTvdb = searchType === 'series'
      ? TVDBAPI.searchSeries(searchQuery)
      : TVDBAPI.searchMovies(searchQuery);

    const [tmdbResult, tvdbResult] = await Promise.allSettled([fetchTmdb, fetchTvdb]);

    setResults({
      tmdb: tmdbResult.status === 'fulfilled' ? tmdbResult.value : [],
      tvdb: tvdbResult.status === 'fulfilled' ? tvdbResult.value : [],
    });
    setErrors({
      tmdb: tmdbResult.status === 'rejected' ? (tmdbResult.reason.message || 'Fehler bei TMDB-Suche') : '',
      tvdb: tvdbResult.status === 'rejected' ? (tvdbResult.reason.message || 'Fehler bei TVDB-Suche') : '',
    });

    setSearchLoading(false);
  };

  const handleResultSelect = (result, source) => {
    setFormData({
      ...formData,
      title: result.name || formData.title,
      description: result.overview || formData.description,
      year: result.year || formData.year,
      image_url: result.image_url || formData.image_url,
      tmdb_id: source === 'tmdb' ? result.id : formData.tmdb_id,
      tvdb_id: source === 'tvdb' ? result.tvdb_id : formData.tvdb_id,
    });
    setResults({ tmdb: [], tvdb: [] });
    setSearchQuery('');
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    if (!formData.title.trim()) {
      alert('Titel ist erforderlich!');
      return;
    }

    onSubmit(formData);
  };

     return (
      <form className="media-form" onSubmit={handleSubmit}>
        <div className="media-search-block">
          <label>Online-Suche (optional)</label>
          <div className="media-search-row">
            <select value={searchType} onChange={(e) => setSearchType(e.target.value)}>
              <option value="movie">Film</option>
              <option value="series">Serie</option>
            </select>
            <input
              type="text"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              placeholder="Titel suchen..."
            />
            <button type="button" onClick={handleSearch} disabled={searchLoading}>
              {searchLoading ? 'Suche...' : 'Suchen'}
            </button>
          </div>

          {(results.tmdb.length > 0 || results.tvdb.length > 0) && (
            <div className="media-search-tabs">
              <button
                type="button"
                className={activeSource === 'tmdb' ? 'active' : ''}
                onClick={() => setActiveSource('tmdb')}
              >
                TMDB ({results.tmdb.length})
              </button>
              <button
                type="button"
                className={activeSource === 'tvdb' ? 'active' : ''}
                onClick={() => setActiveSource('tvdb')}
              >
                TVDB ({results.tvdb.length})
              </button>
            </div>
          )}

          {errors[activeSource] && <div className="error-message">{errors[activeSource]}</div>}

          {results[activeSource].length > 0 && (
            <div className="media-search-results">
              {results[activeSource].map((r) => (
                <div
                  key={`${activeSource}-${r.id || r.tvdb_id}`}
                  className="media-search-result-item"
                  onClick={() => handleResultSelect(r, activeSource)}
                >
                  {r.image_url && <img src={r.image_url} alt={r.name} />}
                  <div className="media-search-result-info">
                    <strong>{r.name}</strong> {r.year ? `(${r.year})` : ''}
                    {r.overview && <p>{r.overview.slice(0, 120)}{r.overview.length > 120 ? '…' : ''}</p>}
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>

        <div className="form-group">
          <label htmlFor="title">Titel *</label>
          <input
            type="text"
            id="title"
            name="title"
            value={formData.title}
            onChange={handleChange}
            required
            placeholder="z.B. The Matrix"
          />
        </div>

        <div className="form-group">
          <label htmlFor="media_type">Medientyp *</label>
          <select
            id="media_type"
            name="media_type"
            value={formData.media_type}
            onChange={handleChange}
          >
            {Object.entries(MEDIA_TYPES).map(([key, value]) => (
              <option key={value} value={value}>
                {MEDIA_TYPE_LABELS[value]}
              </option>
            ))}
          </select>
        </div>

        <div className="form-row">
          <div className="form-group">
            <label htmlFor="artist">Künstler</label>
            <input
              type="text"
              id="artist"
              name="artist"
              value={formData.artist}
              onChange={handleChange}
              placeholder="z.B. Pink Floyd"
            />
          </div>

          <div className="form-group">
            <label htmlFor="director">Regisseur</label>
            <input
              type="text"
              id="director"
              name="director"
              value={formData.director}
              onChange={handleChange}
              placeholder="z.B. Christopher Nolan"
            />
          </div>
        </div>

        <div className="form-row">
          <div className="form-group">
            <label htmlFor="year">Jahr</label>
            <input
              type="number"
              id="year"
              name="year"
              value={formData.year}
              onChange={handleChange}
              min="1900"
              max={new Date().getFullYear()}
            />
          </div>

          <div className="form-group">
            <label htmlFor="genre">Genre</label>
            <input
              type="text"
              id="genre"
              name="genre"
              value={formData.genre}
              onChange={handleChange}
              placeholder="z.B. Science Fiction"
            />
          </div>
        </div>

        <div className="form-row">
          <div className="form-group">
            <label htmlFor="condition">Zustand</label>
            <select
              id="condition"
              name="condition"
              value={formData.condition}
              onChange={handleChange}
            >
              {Object.entries(CONDITIONS).map(([key, value]) => (
                <option key={value} value={value}>
                  {CONDITION_LABELS[value]}
                </option>
              ))}
            </select>
          </div>

          <div className="form-group">
            <label htmlFor="location">Lagerort</label>
            <input
              type="text"
              id="location"
              name="location"
              value={formData.location}
              onChange={handleChange}
              placeholder="z.B. Regal 1, Boden 2"
            />
          </div>
        </div>

        <div className="form-group">
          <label htmlFor="description">Beschreibung</label>
          <textarea
            id="description"
            name="description"
            value={formData.description}
            onChange={handleChange}
            placeholder="Zusätzliche Informationen..."
            rows="4"
          />
        </div>

        <div className="form-buttons">
          <button type="submit" className="btn-submit">
            {media ? 'Aktualisieren' : 'Hinzufügen'}
          </button>
          <button type="button" className="btn-cancel" onClick={onCancel}>
            Abbrechen
          </button>
        </div>
      </form>
    );
}