import React, { useState, useEffect } from 'react';
import { MEDIA_TYPES, MEDIA_TYPE_LABELS, CONDITIONS, CONDITION_LABELS, Media } from '../types';
import TVDBAPI from '../services/tvdb';
import EditionAPI from '../services/editionapi';
import './MediaForm.css';

export default function MediaForm({ media = null, onSubmit, onCancel }) {
  const [formData, setFormData] = useState(
    media || new Media()
  );


  const [tvdbType, setTvdbType] = useState('movie'); // 'movie' | 'series'
  const [tvdbQuery, setTvdbQuery] = useState('');
  const [tvdbResults, setTvdbResults] = useState([]);
  const [tvdbLoading, setTvdbLoading] = useState(false);
  const [tvdbError, setTvdbError] = useState('');
  const [editions, setEditions] = useState([]);

  useEffect(() => {
    EditionAPI.getAllEditions().then(setEditions).catch(() => setEditions([]));
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: name === 'year' ? parseInt(value) : value,
    });
  };

  const handleTvdbSearch = async (e) => {
    e.preventDefault();
    if (!tvdbQuery.trim()) return;
    
    setTvdbLoading(true);
    setTvdbError('');
    setTvdbResults([]);

    try{
      const results = tvdbType === 'series' ? await TVDBAPI.searchSeries(tvdbQuery) : await TVDBAPI.searchMovies(tvdbQuery);
      setTvdbResults(results);
    } catch(error) {
      setTvdbError(error.message || 'Fehler bei der Suche');
    }finally {
      setTvdbLoading(false);
    }
  };

  const handleTvdbSelect = async (result) => {
    setFormData((prev) => ({
      ...prev,
      title: result.name || prev.title,
      description: result.overview || prev.description,
      year: result.year || prev.year,
      image_url: result.image_url || prev.image_url,
      tvdb_id: result.tvdb_id || prev.tvdb_id,
    }));
    setTvdbResults([]);
    setTvdbQuery('');

    // Filme werden über die IMDb-ID identifiziert, TMDB liefert sie nur über einen Extra-Call
    if (result.tmdb_id) {
      try {
        const imdbId = await TVDBAPI.getImdbId(result.tmdb_id);
        setFormData((prev) => ({ ...prev, imdb_id: imdbId }));
        if (!imdbId) setTvdbError('TMDB kennt für diesen Film keine IMDb-ID');
      } catch (error) {
        setTvdbError(error.message || 'IMDb-ID konnte nicht geladen werden');
      }
    }
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
        <div className="tvdb-search-block">
          <label>Online-Suche (optional)</label>
          <div className="tvdb-search-row">
            <select value={tvdbType} onChange={(e) => setTvdbType(e.target.value)}>
              <option value="movie">Film</option>
              <option value="series">Serie</option>
            </select>
            <input
              type="text"
              value={tvdbQuery}
              onChange={(e) => setTvdbQuery(e.target.value)}
              placeholder="Titel suchen..."
            />
            <button type="button" onClick={handleTvdbSearch} disabled={tvdbLoading}>
              {tvdbLoading ? 'Suche...' : 'Suchen'}
            </button>
          </div>

          {tvdbError && <div className="error-message">{tvdbError}</div>}

          {tvdbResults.length > 0 && (
            <div className="tvdb-results">
              {tvdbResults.map((r) => (
                <div
                  key={r.tmdb_id || r.tvdb_id}
                  className="tvdb-result-item"
                  onClick={() => handleTvdbSelect(r)}
                >
                  {r.image_url && <img src={r.image_url} alt={r.name} />}
                  <div className="tvdb-result-info">
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

        <div className="form-group">
          <label htmlFor="imdb_id">IMDb-ID</label>
          <input
            type="text"
            id="imdb_id"
            name="imdb_id"
            value={formData.imdb_id || ''}
            onChange={handleChange}
            placeholder="z.B. tt0133093"
          />
        </div>

        <div className="form-group">
          <label htmlFor="edition">Edition</label>
          <select
            id="edition"
            name="edition"
            value={formData.edition || ''}
            onChange={handleChange}
          >
            <option value="">—</option>
            {editions.map((e) => (
              <option key={e.id} value={e.name}>{e.name}</option>
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