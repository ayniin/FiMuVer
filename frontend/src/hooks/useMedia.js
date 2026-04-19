import { useState, useEffect } from 'react';
import MediaAPI from '../services/api';

export function useMedia() {
  const [media, setMedia] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchAllMedia = async (mediaType = null) => {
    setLoading(true);
    setError(null);
    try {
      const data = await MediaAPI.getAllMedia(mediaType);
      setMedia(data);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const fetchMediaById = async (id) => {
    setLoading(true);
    setError(null);
    try {
      const data = await MediaAPI.getMediaById(id);
      return data;
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const addMedia = async (newMedia) => {
    setLoading(true);
    setError(null);
    try {
      const data = await MediaAPI.createMedia(newMedia);
      setMedia([data, ...media]);
      return data;
    } catch (err) {
      setError(err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  const updateMedia = async (id, updatedMedia) => {
    setLoading(true);
    setError(null);
    try {
      await MediaAPI.updateMedia(id, updatedMedia);
      setMedia(media.map(m => m.id === id ? { ...m, ...updatedMedia } : m));
    } catch (err) {
      setError(err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  const deleteMedia = async (id) => {
    setLoading(true);
    setError(null);
    try {
      await MediaAPI.deleteMedia(id);
      setMedia(media.filter(m => m.id !== id));
    } catch (err) {
      setError(err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  const searchMedia = async (query) => {
    setLoading(true);
    setError(null);
    try {
      const data = await MediaAPI.searchMedia(query);
      setMedia(data);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return {
    media,
    loading,
    error,
    fetchAllMedia,
    fetchMediaById,
    addMedia,
    updateMedia,
    deleteMedia,
    searchMedia,
  };
}

