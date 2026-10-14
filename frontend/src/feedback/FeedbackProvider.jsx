import React, { useCallback, useMemo, useRef, useState } from 'react';
import { FiCheckCircle, FiAlertCircle } from 'react-icons/fi';
import Modal from '../components/Modal';
import { FeedbackContext } from './context';
import './Feedback.css';

const TOAST_DURATION = 3500;

const FeedbackProvider = ({ children }) => {
  const [toasts, setToasts] = useState([]);
  const [confirmState, setConfirmState] = useState(null);
  const nextId = useRef(0);

  const toast = useCallback((message, type = 'success') => {
    const id = nextId.current++;
    setToasts((prev) => [...prev, { id, message, type }]);
    setTimeout(() => setToasts((prev) => prev.filter((t) => t.id !== id)), TOAST_DURATION);
  }, []);

  const confirm = useCallback(
    (options) => new Promise((resolve) => setConfirmState({ ...options, resolve })),
    []
  );

  const closeConfirm = (result) => {
    confirmState?.resolve(result);
    setConfirmState(null);
  };

  const value = useMemo(() => ({ toast, confirm }), [toast, confirm]);

  return (
    <FeedbackContext.Provider value={value}>
      {children}

      <Modal open={Boolean(confirmState)} onClose={() => closeConfirm(false)} title={confirmState?.title} size="sm">
        {confirmState?.message && <p>{confirmState.message}</p>}
        <div className="modal-actions">
          <button type="button" className="btn btn--secondary" onClick={() => closeConfirm(false)}>
            Abbrechen
          </button>
          <button
            type="button"
            className={`btn ${confirmState?.danger ? 'btn--danger' : 'btn--primary'}`}
            onClick={() => closeConfirm(true)}
            autoFocus
          >
            {confirmState?.confirmLabel || 'OK'}
          </button>
        </div>
      </Modal>

      <div className="toast-region" role="status" aria-live="polite">
        {toasts.map((t) => (
          <div key={t.id} className={`toast toast--${t.type}`}>
            {t.type === 'error' ? <FiAlertCircle size={18} /> : <FiCheckCircle size={18} />}
            <span>{t.message}</span>
          </div>
        ))}
      </div>
    </FeedbackContext.Provider>
  );
};

export default FeedbackProvider;
