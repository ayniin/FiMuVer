import React from 'react';
import { FiSearch } from 'react-icons/fi';
import { MEDIA_TYPES, MEDIA_TYPE_LABELS } from '../types';
import './FilterBar.css';

const SORT_OPTIONS = {
  title: 'Titel (A–Z)',
  year_desc: 'Jahr (neueste zuerst)',
  year_asc: 'Jahr (älteste zuerst)',
  created_desc: 'Zuletzt hinzugefügt',
};

export default function FilterBar({ search, onSearch, type, onType, sort, onSort }) {
  return (
    <div className="filter-bar">
      <div className="filter-bar-row">
        <label className="filter-search">
          <FiSearch size={17} aria-hidden="true" />
          <span className="visually-hidden">Suchen</span>
          <input
            className="input"
            type="search"
            placeholder="Titel, Künstler oder Regisseur…"
            value={search}
            onChange={(e) => onSearch(e.target.value)}
          />
        </label>

        <label className="filter-sort">
          <span className="visually-hidden">Sortieren</span>
          <select className="input" value={sort} onChange={(e) => onSort(e.target.value)}>
            {Object.entries(SORT_OPTIONS).map(([value, label]) => (
              <option key={value} value={value}>{label}</option>
            ))}
          </select>
        </label>
      </div>

      <div className="filter-chips" role="group" aria-label="Nach Medientyp filtern">
        <button type="button" className="chip" aria-pressed={type === null} onClick={() => onType(null)}>
          Alle
        </button>
        {Object.values(MEDIA_TYPES).map((value) => (
          <button
            key={value}
            type="button"
            className="chip"
            aria-pressed={type === value}
            onClick={() => onType(value)}
          >
            {MEDIA_TYPE_LABELS[value]}
          </button>
        ))}
      </div>
    </div>
  );
}
