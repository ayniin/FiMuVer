import React, { useState } from 'react';
import { FiSearch } from 'react-icons/fi';
import { MEDIA_TYPES, MEDIA_TYPE_LABELS, CONDITIONS, CONDITION_LABELS, Media, isMusicType } from '../types';
import TVDBAPI from '../services/tvdb';
import TMDBAPI from '../services/tmdb';
import DiscogsAPI from '../services/discogs';
import MediaCover from './MediaCover';
import './MediaForm.css';

const EMPTY_RESULTS = { tmdb: [], tvdb: [], discogs: [] };
const EMPTY_ERRORS = { tmdb: '', tvdb: '', discogs: '' };
const SOURCE_LABELS = { tmdb: 'TMDB', tvdb: 'TVDB', discogs: 'Discogs' };

// Discogs liefert den Titel als „Künstler - Album"
const fromDiscogs = (r) => {
  const [artist, ...rest] = (r.title || '').split(' - ');
  return {
    key: r.id,
    name: rest.length ? rest.join(' - ') : artist,
    artist: rest.length ? artist : '',
    year: parseInt(r.year, 10) || null,
    overview: [r.genre?.join(', '), r.format?.join(', '), r.country].filter(Boolean).join(' · '),
    image_url: r.cover_image,
  };
};

export default function MediaForm({ media = null, onSubmit, onCancel }) {
  const [formData, setFormData] = useState(() => ({
    ...new Media(),
    ...media,
    year: media?.year || '',
  }));
  const [submitting, setSubmitting] = useState(false);
  const [submitError, setSubmitError] = useState('');

  const isMusic = isMusicType(formData.media_type);
  const sources = isMusic ? ['discogs'] : ['tmdb', 'tvdb'];

  const [searchType, setSearchType] = useState('movie'); // 'movie' | 'series'
  const [searchQuery, setSearchQuery] = useState('');
  const [activeSource, setActiveSource] = useState(sources[0]);
  const [results, setResults] = useState(EMPTY_RESULTS);
  const [errors, setErrors] = useState(EMPTY_ERRORS);
  const [searchLoading, setSearchLoading] = useState(false);
  const [searched, setSearched] = useState(false);

  const resetSearch = () => {
    setResults(EMPTY_RESULTS);
    setErrors(EMPTY_ERRORS);
    setSearched(false);
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });

    if (name === 'media_type' && isMusicType(value) !== isMusic) {
      resetSearch();
      setActiveSource(isMusicType(value) ? 'discogs' : 'tmdb');
    }
  };

  const runSearch = async () => {
    const query = searchQuery.trim();
    if (!query) return;

    setSearchLoading(true);
    resetSearch();

    const requests = isMusic
      ? { discogs: DiscogsAPI.searchReleases(query).then((list) => list.map(fromDiscogs)) }
      : {
          tmdb: searchType === 'series' ? TMDBAPI.searchSeries(query) : TMDBAPI.searchMovies(query),
          tvdb: searchType === 'series' ? TVDBAPI.searchSeries(query) : TVDBAPI.searchMovies(query),
        };

    const names = Object.keys(requests);
    const settled = await Promise.allSettled(Object.values(requests));
    const nextResults = { ...EMPTY_RESULTS };
    const nextErrors = { ...EMPTY_ERRORS };
    settled.forEach((res, i) => {
      const source = names[i];
      if (res.status === 'fulfilled') nextResults[source] = res.value || [];
      else nextErrors[source] = res.reason?.message || `Fehler bei der ${SOURCE_LABELS[source]}-Suche`;
    });

    setResults(nextResults);
    setErrors(nextErrors);
    setSearched(true);
    setSearchLoading(false);
    // Zur ersten Quelle mit Treffern springen
    setActiveSource(names.find((n) => nextResults[n].length > 0) || names[0]);
  };

  const handleSearchKeyDown = (e) => {
    // Enter im Suchfeld soll suchen, nicht das ganze Formular absenden
    if (e.key === 'Enter') {
      e.preventDefault();
      runSearch();
    }
  };

  const handleResultSelect = (result, source) => {
    setFormData({
      ...formData,
      title: result.name || formData.title,
      artist: result.artist || formData.artist,
      description: result.overview && source !== 'discogs' ? result.overview : formData.description,
      year: result.year || formData.year,
      image_url: result.image_url || formData.image_url,
      tmdb_id: source === 'tmdb' ? result.id : formData.tmdb_id,
      tvdb_id: source === 'tvdb' ? result.tvdb_id : formData.tvdb_id,
    });
    resetSearch();
    setSearchQuery('');
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setSubmitting(true);
    setSubmitError('');
    try {
      await onSubmit({
        ...formData,
        title: formData.title.trim(),
        year: parseInt(formData.year, 10) || 0,
        // Nur das zum Medientyp passende Feld behalten
        artist: isMusic ? formData.artist : '',
        director: isMusic ? '' : formData.director,
      });
    } catch (err) {
      setSubmitError(err.message || 'Der Eintrag konnte nicht gespeichert werden.');
      setSubmitting(false);
    }
  };

  const activeResults = results[activeSource] || [];
  const totalResults = sources.reduce((sum, s) => sum + results[s].length, 0);

  return (
    <form className="media-form" onSubmit={handleSubmit}>
      <div className="field">
        <span className="field-label" id="media-type-label">Medientyp</span>
        <div className="media-type-picker" role="radiogroup" aria-labelledby="media-type-label">
          {Object.values(MEDIA_TYPES).map((value) => (
            <button
              key={value}
              type="button"
              role="radio"
              aria-checked={formData.media_type === value}
              className="chip"
              onClick={() => handleChange({ target: { name: 'media_type', value } })}
            >
              {MEDIA_TYPE_LABELS[value]}
            </button>
          ))}
        </div>
      </div>

      <section className="media-search" aria-label="Online-Suche">
        <label htmlFor="media-search-input">
          Online-Suche <span className="field-hint">
            (optional, über {sources.map((s) => SOURCE_LABELS[s]).join(' und ')})
          </span>
        </label>
        <div className="media-search-row">
          {!isMusic && (
            <select
              className="input"
              value={searchType}
              onChange={(e) => setSearchType(e.target.value)}
              aria-label="Art der Suche"
            >
              <option value="movie">Film</option>
              <option value="series">Serie</option>
            </select>
          )}
          <input
            className="input"
            id="media-search-input"
            type="search"
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            onKeyDown={handleSearchKeyDown}
            placeholder={isMusic ? 'Album oder Künstler suchen…' : 'Titel suchen…'}
          />
          <button
            type="button"
            className="btn btn--secondary"
            onClick={runSearch}
            disabled={searchLoading || !searchQuery.trim()}
          >
            <FiSearch size={16} />
            {searchLoading ? 'Suche…' : 'Suchen'}
          </button>
        </div>

        {searched && sources.length > 1 && (
          <div className="tabs media-search-tabs" role="tablist">
            {sources.map((source) => (
              <button
                key={source}
                type="button"
                role="tab"
                className="tab"
                aria-selected={activeSource === source}
                onClick={() => setActiveSource(source)}
              >
                {SOURCE_LABELS[source]} ({results[source].length})
              </button>
            ))}
          </div>
        )}

        {errors[activeSource] && <div className="alert alert--error">{errors[activeSource]}</div>}

        {searched && totalResults === 0 && !errors[activeSource] && (
          <p className="media-search-empty">Keine Treffer gefunden.</p>
        )}

        {activeResults.length > 0 && (
          <ul className="media-search-results">
            {activeResults.map((r) => (
              <li key={`${activeSource}-${r.key ?? r.id ?? r.tvdb_id}`}>
                <button type="button" className="media-search-result" onClick={() => handleResultSelect(r, activeSource)}>
                  <MediaCover media={{ ...r, media_type: formData.media_type }} className="media-search-cover" />
                  <span className="media-search-result-info">
                    <strong>{r.name}</strong>
                    {(r.artist || r.year) && (
                      <span className="media-search-result-meta">
                        {[r.artist, r.year].filter(Boolean).join(' · ')}
                      </span>
                    )}
                    {r.overview && (
                      <span className="media-search-result-text">
                        {r.overview.length > 140 ? `${r.overview.slice(0, 140)}…` : r.overview}
                      </span>
                    )}
                  </span>
                </button>
              </li>
            ))}
          </ul>
        )}
      </section>

      <div className="media-form-main">
        <MediaCover media={formData} className="media-form-cover" />

        <div className="media-form-fields">
          <div className="field">
            <label htmlFor="title">Titel</label>
            <input
              className="input"
              type="text"
              id="title"
              name="title"
              value={formData.title}
              onChange={handleChange}
              required
              placeholder={isMusic ? 'z.B. The Dark Side of the Moon' : 'z.B. The Matrix'}
            />
          </div>

          <div className="field-row">
            {isMusic ? (
              <div className="field">
                <label htmlFor="artist">Künstler</label>
                <input
                  className="input"
                  type="text"
                  id="artist"
                  name="artist"
                  value={formData.artist}
                  onChange={handleChange}
                  placeholder="z.B. Pink Floyd"
                />
              </div>
            ) : (
              <div className="field">
                <label htmlFor="director">Regisseur</label>
                <input
                  className="input"
                  type="text"
                  id="director"
                  name="director"
                  value={formData.director}
                  onChange={handleChange}
                  placeholder="z.B. Christopher Nolan"
                />
              </div>
            )}

            <div className="field">
              <label htmlFor="year">Jahr</label>
              <input
                className="input"
                type="number"
                id="year"
                name="year"
                value={formData.year}
                onChange={handleChange}
                min="1900"
                max={new Date().getFullYear() + 1}
              />
            </div>
          </div>

          <div className="field-row">
            <div className="field">
              <label htmlFor="genre">Genre</label>
              <input
                className="input"
                type="text"
                id="genre"
                name="genre"
                value={formData.genre}
                onChange={handleChange}
                placeholder={isMusic ? 'z.B. Progressive Rock' : 'z.B. Science Fiction'}
              />
            </div>

            <div className="field">
              <label htmlFor="condition">Zustand</label>
              <select
                className="input"
                id="condition"
                name="condition"
                value={formData.condition}
                onChange={handleChange}
              >
                {Object.values(CONDITIONS).map((value) => (
                  <option key={value} value={value}>
                    {CONDITION_LABELS[value]}
                  </option>
                ))}
              </select>
            </div>
          </div>

          <div className="field">
            <label htmlFor="location">Lagerort</label>
            <input
              className="input"
              type="text"
              id="location"
              name="location"
              value={formData.location}
              onChange={handleChange}
              placeholder="z.B. Regal 1, Boden 2"
            />
          </div>
        </div>
      </div>

      <div className="field">
        <label htmlFor="description">Beschreibung</label>
        <textarea
          className="input"
          id="description"
          name="description"
          value={formData.description}
          onChange={handleChange}
          placeholder="Zusätzliche Informationen…"
          rows="4"
        />
      </div>

      {submitError && <div className="alert alert--error" role="alert">{submitError}</div>}

      <div className="modal-actions">
        <button type="button" className="btn btn--secondary" onClick={onCancel} disabled={submitting}>
          Abbrechen
        </button>
        <button type="submit" className="btn btn--primary" disabled={submitting}>
          {submitting ? 'Wird gespeichert…' : media ? 'Änderungen speichern' : 'Hinzufügen'}
        </button>
      </div>
    </form>
  );
}
