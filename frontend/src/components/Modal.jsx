import React, { useEffect, useId, useRef } from 'react';
import { FiX } from 'react-icons/fi';
import './Modal.css';

/**
 * Nutzt das native <dialog>-Element: Escape schließt, der Fokus bleibt im Dialog
 * und kehrt danach automatisch zum auslösenden Element zurück.
 */
const Modal = ({ open, onClose, title, children, size = 'md' }) => {
  const dialogRef = useRef(null);
  const titleId = useId();

  useEffect(() => {
    const dialog = dialogRef.current;
    if (!dialog) return;
    if (open && !dialog.open) dialog.showModal();
    if (!open && dialog.open) dialog.close();
  }, [open]);

  const handleCancel = (e) => {
    e.preventDefault();
    onClose();
  };

  const handleBackdropClick = (e) => {
    if (e.target === dialogRef.current) onClose();
  };

  return (
    <dialog
      ref={dialogRef}
      className={`modal modal--${size}`}
      aria-labelledby={titleId}
      onCancel={handleCancel}
      onClick={handleBackdropClick}
    >
      {open && (
        <div className="modal-body">
          <div className="modal-head">
            <h2 id={titleId}>{title}</h2>
            <button type="button" className="btn btn--ghost btn--icon" onClick={onClose} aria-label="Schließen">
              <FiX size={20} />
            </button>
          </div>
          {children}
        </div>
      )}
    </dialog>
  );
};

export default Modal;
