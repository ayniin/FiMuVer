import React, { useState, useEffect } from 'react';
import { FiCheck, FiCopy, FiPlus, FiTrash2 } from 'react-icons/fi';
import InviteAPI from '../services/inviteapi';
import { useFeedback } from '../feedback/context';
import './InviteCodes.css';

const formatDate = (dateStr) => (dateStr ? new Date(dateStr).toLocaleDateString('de-DE') : '—');

const InviteCodes = () => {
    const { toast, confirm } = useFeedback();
    const [codes, setCodes] = useState([]);
    const [loading, setLoading] = useState(true);
    const [generating, setGenerating] = useState(false);
    const [error, setError] = useState('');
    const [copied, setCopied] = useState(null);

    useEffect(() => {
        InviteAPI.listCodes()
            .then((data) => setCodes(data || []))
            .catch((err) => setError(err.message || 'Die Einladungscodes konnten nicht geladen werden.'))
            .finally(() => setLoading(false));
    }, []);

    const handleGenerate = async () => {
        setGenerating(true);
        try {
            await InviteAPI.generateCode();
            setCodes((await InviteAPI.listCodes()) || []);
            toast('Neuer Einladungscode erstellt');
        } catch (err) {
            toast(err.message || 'Der Code konnte nicht erstellt werden.', 'error');
        } finally {
            setGenerating(false);
        }
    };

    const handleDelete = async (code) => {
        const ok = await confirm({
            title: 'Einladungscode löschen?',
            message: `Der Code ${code.code} kann danach nicht mehr verwendet werden.`,
            confirmLabel: 'Löschen',
            danger: true,
        });
        if (!ok) return;

        try {
            await InviteAPI.deleteCode(code.id);
            setCodes((prev) => prev.filter((c) => c.id !== code.id));
            toast('Einladungscode gelöscht');
        } catch (err) {
            toast(err.message || 'Der Code konnte nicht gelöscht werden.', 'error');
        }
    };

    const handleCopy = async (code) => {
        try {
            await navigator.clipboard.writeText(code);
            setCopied(code);
            setTimeout(() => setCopied(null), 2000);
        } catch {
            toast('Kopieren nicht möglich – bitte manuell markieren.', 'error');
        }
    };

    return (
        <div className="invite-codes">
            <div className="panel-head">
                <p className="invite-intro">Neue Nutzer können sich mit einem Einladungscode registrieren.</p>
                <button className="btn btn--primary" onClick={handleGenerate} disabled={generating}>
                    <FiPlus size={17} />
                    {generating ? 'Wird erstellt…' : 'Code generieren'}
                </button>
            </div>

            {error && <div className="alert alert--error" role="alert">{error}</div>}

            {loading ? (
                <div className="skeleton admin-skeleton" aria-busy="true" aria-label="Wird geladen" />
            ) : codes.length === 0 ? (
                !error && <p className="empty-state">Noch keine Einladungscodes vorhanden.</p>
            ) : (
                <div className="invite-table-wrap">
                    <table className="invite-table">
                        <thead>
                            <tr>
                                <th scope="col">Code</th>
                                <th scope="col">Verwendet</th>
                                <th scope="col">Läuft ab</th>
                                <th scope="col">Erstellt</th>
                                <th scope="col"><span className="visually-hidden">Aktionen</span></th>
                            </tr>
                        </thead>
                        <tbody>
                            {codes.map((c) => (
                                <tr key={c.id} className={c.used_by_user_id ? 'invite-used' : ''}>
                                    <td>
                                        <span className="invite-code">
                                            <code>{c.code}</code>
                                            <button
                                                type="button"
                                                className="btn btn--ghost btn--icon"
                                                onClick={() => handleCopy(c.code)}
                                                aria-label={copied === c.code ? 'Kopiert' : `Code ${c.code} kopieren`}
                                                title="Kopieren"
                                            >
                                                {copied === c.code ? <FiCheck size={15} /> : <FiCopy size={15} />}
                                            </button>
                                        </span>
                                    </td>
                                    <td>{c.current_uses} / {c.max_uses === 0 ? '∞' : c.max_uses}</td>
                                    <td>{formatDate(c.expires_at)}</td>
                                    <td>{formatDate(c.created_at)}</td>
                                    <td className="invite-actions">
                                        <button
                                            type="button"
                                            className="btn btn--ghost btn--icon btn--danger-hover"
                                            onClick={() => handleDelete(c)}
                                            aria-label={`Code ${c.code} löschen`}
                                            title="Löschen"
                                        >
                                            <FiTrash2 size={16} />
                                        </button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}
        </div>
    );
};

export default InviteCodes;
