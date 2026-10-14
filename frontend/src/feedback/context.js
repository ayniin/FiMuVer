import { createContext, useContext } from 'react';

export const FeedbackContext = createContext(null);

/**
 * Zugriff auf Toasts und Bestätigungsdialog:
 *   const { toast, confirm } = useFeedback();
 *   toast('Gespeichert');
 *   toast('Fehler', 'error');
 *   if (await confirm({ title: 'Löschen?', message: '…', confirmLabel: 'Löschen', danger: true })) { … }
 */
export const useFeedback = () => useContext(FeedbackContext);
