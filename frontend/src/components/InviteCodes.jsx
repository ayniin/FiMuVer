import React, { useState, useEffect, useCallback } from 'react';
import InviteAPI from '../services/inviteapi';
import './InviteCodes.css';

const InviteCodes = () => {
    const [codes, setCodes] = useState([]);
    const [loading, setLoading] = useState(false);
    const [generating, setGenerating] = useState(false);
    const [error, setError] = useState('');
    const [copied, setCopied] = useState(null);

    const loadCodes = useCallback(async () => {
        setLoading(true);
        try {
            const data = await InviteAPI.listCodes();
            setCodes(data || []);
        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    }, []);

    useEffect(() => {
        loadCodes();
    }, [loadCodes]);

    const handleGenerate = async () => {
        setGenerating(true);
        setError('');
        try {
            await InviteAPI.generateCode();
            await loadCodes();
        } catch (err) {
            setError(err.message);
        } finally {
            setGenerating(false);
        }
    };

    const handleDelete = async (id) => {
        if (!window.confirm('Invite Code wirklich löschen?')) return;
        try {
            await InviteAPI.deleteCode(id);
            setCodes((prev) => prev.filter((c) => c.id !== id));
        } catch (err) {
            setError(err.message);
        }
    };

    const handleCopy = (code) => {
        navigator.clipboard.writeText(code);
        setCopied(code);
        setTimeout(() => setCopied(null), 2000);
    };

    const formatDate = (dateStr) => {
        if (!dateStr) return '—';
        return new Date(dateStr).toLocaleDateString('de-DE');
    };

    return (
        <div className="invite-codes">
            <div className="invite-codes-header">
                <h3>Invite Codes</h3>
                <button
                    className="invite-generate-btn"
                    onClick={handleGenerate}
                    disabled={generating}
                >
                    {generating ? 'Wird erstellt...' : '+ Code generieren'}
                </button>
            </div>

            {error && <p className="invite-error">{error}</p>}

            {loading ? (
                <p className="invite-loading">Lade Codes...</p>
            ) : codes.length === 0 ? (
                <p className="invite-empty">Noch keine Invite Codes vorhanden.</p>
            ) : (
                <table className="invite-table">
                    <thead>
                        <tr>
                            <th>Code</th>
                            <th>Verwendungen</th>
                            <th>Läuft ab</th>
                            <th>Erstellt</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {codes.map((c) => (
                            <tr key={c.id} className={c.used_by_user_id ? 'invite-used' : ''}>
                                <td>
                                    <span className="invite-code-text">{c.code}</span>
                                    <button
                                        className="invite-copy-btn"
                                        onClick={() => handleCopy(c.code)}
                                        title="Kopieren"
                                    >
                                        {copied === c.code ? '✓' : '⎘'}
                                    </button>
                                </td>
                                <td>
                                    {c.current_uses} / {c.max_uses === 0 ? '∞' : c.max_uses}
                                </td>
                                <td>{formatDate(c.expires_at)}</td>
                                <td>{formatDate(c.created_at)}</td>
                                <td>
                                    <button
                                        className="invite-delete-btn"
                                        onClick={() => handleDelete(c.id)}
                                        title="Löschen"
                                    >
                                        ✕
                                    </button>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            )}
        </div>
    );
};

export default InviteCodes;
