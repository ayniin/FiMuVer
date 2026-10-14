import React, { useEffect, useState } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import { FiArrowLeft, FiEdit2, FiTrash2 } from 'react-icons/fi';
import { MEDIA_TYPE_LABELS, CONDITION_LABELS, isMusicType } from '../types';
import CollectionAPI from '../services/collection';
import MediaForm from '../components/MediaForm';
import MediaCover from '../components/MediaCover';
import Modal from '../components/Modal';
import { useFeedback } from '../feedback/context';
import './ItemDetailPage.css';

const ItemDetailPage = () => {
  const { collectionId, itemId } = useParams();
  const navigate = useNavigate();
  const { toast, confirm } = useFeedback();
  const [collection, setCollection] = useState(null);
  const [item, setItem] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [editing, setEditing] = useState(false);

  // Es gibt keinen Einzel-Endpunkt für Items, daher über die Sammlung laden
  useEffect(() => {
    const loadItem = async () => {
      setLoading(true);
      setError('');
      try {
        const data = await CollectionAPI.getCollectionById(collectionId);
        const found = data?.items?.find((i) => String(i.id) === itemId);
        setCollection(data);
        setItem(found || null);
        if (!found) setError('Dieser Eintrag wurde nicht gefunden.');
      } catch (err) {
        setError(err.message || 'Der Eintrag konnte nicht geladen werden.');
      } finally {
        setLoading(false);
      }
    };
    loadItem();
  }, [collectionId, itemId]);

  const backTo = `/collections/${collectionId}`;

  const handleUpdate = async (formData) => {
    const updated = await CollectionAPI.updateItem(collectionId, item.id, formData);
    setItem({ ...item, ...formData, ...updated });
    setEditing(false);
    toast('Änderungen gespeichert');
  };

  const handleDelete = async () => {
    const ok = await confirm({
      title: 'Eintrag löschen?',
      message: `„${item.title}" wird dauerhaft aus der Sammlung entfernt.`,
      confirmLabel: 'Löschen',
      danger: true,
    });
    if (!ok) return;

    try {
      await CollectionAPI.deleteItem(collectionId, item.id);
      toast('Eintrag gelöscht');
      navigate(backTo);
    } catch (err) {
      toast(err.message || 'Der Eintrag konnte nicht gelöscht werden.', 'error');
    }
  };

  const details = item ? [
    ['Typ', MEDIA_TYPE_LABELS[item.media_type]],
    [isMusicType(item.media_type) ? 'Künstler' : 'Regisseur', isMusicType(item.media_type) ? item.artist : item.director],
    ['Jahr', item.year || null],
    ['Genre', item.genre],
    ['Zustand', CONDITION_LABELS[item.condition]],
    ['Lagerort', item.location],
  ].filter(([, value]) => value) : [];

  return (
    <>
      <Link to={backTo} className="breadcrumb">
        <FiArrowLeft size={15} /> {collection?.name || 'Zurück zur Sammlung'}
      </Link>

      {error && <div className="alert alert--error" role="alert">{error}</div>}

      {loading && (
        <div className="item-detail" aria-busy="true" aria-label="Wird geladen">
          <div className="skeleton item-detail-cover" />
          <div>
            <div className="skeleton skeleton--title" />
          </div>
        </div>
      )}

      {!loading && item && (
        <article className="item-detail">
          <MediaCover media={item} className="item-detail-cover" />

          <div className="item-detail-content">
            <div className="item-detail-head">
              <h1>{item.title}</h1>
              <div className="page-head-actions">
                <button className="btn btn--secondary" onClick={() => setEditing(true)}>
                  <FiEdit2 size={16} /> Bearbeiten
                </button>
                <button className="btn btn--ghost btn--danger-hover" onClick={handleDelete}>
                  <FiTrash2 size={16} /> Löschen
                </button>
              </div>
            </div>

            {details.length > 0 && (
              <dl className="item-detail-list card">
                {details.map(([label, value]) => (
                  <div key={label}>
                    <dt>{label}</dt>
                    <dd>{value}</dd>
                  </div>
                ))}
              </dl>
            )}

            {item.description && (
              <section className="item-detail-description">
                <h2>Beschreibung</h2>
                <p>{item.description}</p>
              </section>
            )}
          </div>
        </article>
      )}

      {item && (
        <Modal open={editing} onClose={() => setEditing(false)} title="Eintrag bearbeiten" size="lg">
          <MediaForm media={item} onSubmit={handleUpdate} onCancel={() => setEditing(false)} />
        </Modal>
      )}
    </>
  );
};

export default ItemDetailPage;
