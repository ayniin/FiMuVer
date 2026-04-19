import React, { useState } from 'react';
import { MEDIA_TYPES, MEDIA_TYPE_LABELS } from '../types';
import './FilterBar.css';

export default function FilterBar({ onFilter, onSearch }) {
  const [searchQuery, setSearchQuery] = useState('');

  const handleSearch = (e) => {
    const query = e.target.value;
    setSearchQuery(query);
    onSearch(query);
  };

  const handleTypeFilter = (type) => {
    onFilter(type);
  };

  return (
    <div className="filter-bar">
      <div className="search-box">
        <input
          type="text"
          placeholder="Nach Titel, Künstler oder Regisseur suchen..."
          value={searchQuery}
          onChange={handleSearch}
          className="search-input"
        />
      </div>

      <div className="filter-buttons">
        <button
          className="filter-btn active"
          onClick={() => handleTypeFilter(null)}
        >
          Alle
        </button>
        {Object.entries(MEDIA_TYPES).map(([key, value]) => (
          <button
            key={value}
            className="filter-btn"
            onClick={() => handleTypeFilter(value)}
          >
            {MEDIA_TYPE_LABELS[value]}
          </button>
        ))}
      </div>
    </div>
  );
}

