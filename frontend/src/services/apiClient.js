const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';

const getAuthToken = () => {
  try {
    return sessionStorage.getItem('auth_token');
  } catch {
    return null;
  }
};

/**
 * CSRF Token aus Cookie auslesen
 */
const getCsrfToken = () => {
  const name = 'X-CSRF-Token';
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(';').shift();
  return null;
};

/**
 * Fetch-Wrapper mit:
 * - HttpOnly Cookie Auto-Versand (credentials: 'include')
 * - CSRF-Token Protection
 * - Error Handling
 * - JSON Parsing
 */
export const apiRequest = async (endpoint, options = {}) => {
  const {
    method = 'GET',
    body = null,
    headers = {},
    ...fetchOptions
  } = options;

  // Headers vorbereiten
  const finalHeaders = {
    'Content-Type': 'application/json',
    ...headers,
  };

  const authToken = getAuthToken();
  if (authToken && !finalHeaders.Authorization && !finalHeaders.authorization) {
    finalHeaders.Authorization = `Bearer ${authToken}`;
  }

  // CSRF-Token nur für Mutation-Requests (POST, PUT, DELETE, PATCH)
  if (['POST', 'PUT', 'DELETE', 'PATCH'].includes(method.toUpperCase())) {
    const csrfToken = getCsrfToken();
    if (csrfToken) {
      finalHeaders['X-CSRF-Token'] = csrfToken;
    }
  }

  const url = endpoint.startsWith('http')
    ? endpoint
    : `${API_BASE_URL}${endpoint}`;

  try {
    const response = await fetch(url, {
      method,
      headers: finalHeaders,
      body: body ? JSON.stringify(body) : null,
      // WICHTIG: Credentials inkludieren für Cookie-Auth
      credentials: 'include',
      ...fetchOptions,
    });

    // Error Handling
    if (!response.ok) {
      let errorMessage = `HTTP ${response.status}`;
      
      try {
        const errorData = await response.json();
        errorMessage = errorData.error || errorData.message || errorMessage;
      } catch {
        // Response war nicht JSON
      }

      const error = new Error(errorMessage);
      error.status = response.status;
      throw error;
    }

    // Leere Response (204 No Content)
    if (response.status === 204) {
      return null;
    }

    // JSON Response
    try {
      const data = await response.json();
      return data;
    } catch {
      // Response war nicht JSON (z.B. 200 OK mit Empty Body)
      return null;
    }
  } catch (error) {
    console.error(`API Request failed: ${method} ${url}`, error);
    throw error;
  }
};

/**
 * Convenience Methods
 */
export const apiGet = (endpoint, options = {}) =>
  apiRequest(endpoint, { method: 'GET', ...options });

export const apiPost = (endpoint, body, options = {}) =>
  apiRequest(endpoint, { method: 'POST', body, ...options });

export const apiPut = (endpoint, body, options = {}) =>
  apiRequest(endpoint, { method: 'PUT', body, ...options });

export const apiDelete = (endpoint, options = {}) =>
  apiRequest(endpoint, { method: 'DELETE', ...options });

export const apiPatch = (endpoint, body, options = {}) =>
  apiRequest(endpoint, { method: 'PATCH', body, ...options });
