// Medientypen
export const MEDIA_TYPES = {
  BLURAY: 'bluray',
  DVD: 'dvd',
  VINYL: 'vinyl',
  TAPE: 'tape'
};

export const MEDIA_TYPE_LABELS = {
  bluray: 'Bluray',
  dvd: 'DVD',
  vinyl: 'Vinyl',
  tape: 'Tape'
};

export const CONDITIONS = {
  MINT: 'mint',
  GOOD: 'good',
  FAIR: 'fair',
  POOR: 'poor'
};

export const CONDITION_LABELS = {
  mint: 'Mint',
  good: 'Gut',
  fair: 'Akzeptabel',
  poor: 'Schlecht'
};

// API Response Types
export class Media {
  id = null;
  title = '';
  description = '';
  media_type = MEDIA_TYPES.BLURAY;
  artist = '';
  director = '';
  year = new Date().getFullYear();
  genre = '';
  condition = CONDITIONS.GOOD;
  location = '';
  notes = {};
  created_at = new Date();
  updated_at = new Date();

  constructor(data = {}) {
    Object.assign(this, data);
  }
}

export class ApiResponse {
  data = null;
  error = null;
  message = '';

  constructor(data = {}) {
    Object.assign(this, data);
  }
}

