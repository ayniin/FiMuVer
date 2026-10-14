import React, { useState, useEffect } from 'react';
import { FiTrash2 } from 'react-icons/fi';
import SettingsAPI from '../services/settingsapi';
import Editions from './Editions';
import InviteCodes from '../components/InviteCodes';
import { useFeedback } from '../feedback/context';
import './Admin.css';

const TABS = {
  settings: 'Einstellungen',
  invites: 'Einladungscodes',
  editions: 'Editionen',
};

const Settings = () => {
  const { toast, confirm } = useFeedback();
  const [settings, setSettings] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const loadSettings = async () => {
      try {
        const data = await SettingsAPI.getAllSettings();
        setSettings(data);
      } catch (err) {
        setError(err.message || 'Die Einstellungen konnten nicht geladen werden.');
      } finally {
        setLoading(false);
      }
    };
    loadSettings();
  }, []);

  const handleToggle = async (setting) => {
    const newValue = !setting.value;
    try {
      await SettingsAPI.updateSetting(setting.name, { value: newValue });
      setSettings((prev) => prev.map((s) => (s.name === setting.name ? { ...s, value: newValue } : s)));
      toast(`„${setting.name}" ${newValue ? 'aktiviert' : 'deaktiviert'}`);
    } catch (err) {
      toast(err.message || 'Die Einstellung konnte nicht geändert werden.', 'error');
    }
  };

  const handleDelete = async (setting) => {
    const ok = await confirm({
      title: 'Einstellung löschen?',
      message: `„${setting.name}" wird dauerhaft entfernt.`,
      confirmLabel: 'Löschen',
      danger: true,
    });
    if (!ok) return;

    try {
      await SettingsAPI.deleteSetting(setting.id);
      setSettings((prev) => prev.filter((s) => s.id !== setting.id));
      toast('Einstellung gelöscht');
    } catch (err) {
      toast(err.message || 'Die Einstellung konnte nicht gelöscht werden.', 'error');
    }
  };

  if (loading) return <div className="skeleton admin-skeleton" aria-busy="true" aria-label="Wird geladen" />;
  if (error) return <div className="alert alert--error" role="alert">{error}</div>;
  if (settings.length === 0) return <p className="empty-state">Keine Einstellungen vorhanden.</p>;

  return (
    <ul className="admin-list">
      {settings.map((setting) => (
        <li key={setting.id} className="admin-list-item">
          <span className="admin-list-name" id={`setting-${setting.id}`}>{setting.name}</span>
          <div className="admin-list-actions">
            <button
              type="button"
              role="switch"
              className="switch"
              aria-checked={Boolean(setting.value)}
              aria-labelledby={`setting-${setting.id}`}
              onClick={() => handleToggle(setting)}
            >
              <span className="switch-thumb" />
            </button>
            <button
              type="button"
              className="btn btn--ghost btn--icon btn--danger-hover"
              onClick={() => handleDelete(setting)}
              aria-label={`„${setting.name}" löschen`}
              title="Löschen"
            >
              <FiTrash2 size={17} />
            </button>
          </div>
        </li>
      ))}
    </ul>
  );
};

const Admin = () => {
  const [activeTab, setActiveTab] = useState('settings');

  return (
    <>
      <div className="page-head">
        <div>
          <h1>Administration</h1>
          <p>Einstellungen, Einladungen und Stammdaten verwalten.</p>
        </div>
      </div>

      <div className="tabs" role="tablist">
        {Object.entries(TABS).map(([key, label]) => (
          <button
            key={key}
            type="button"
            role="tab"
            className="tab"
            id={`tab-${key}`}
            aria-selected={activeTab === key}
            aria-controls={`panel-${key}`}
            onClick={() => setActiveTab(key)}
          >
            {label}
          </button>
        ))}
      </div>

      <section className="card panel" role="tabpanel" id={`panel-${activeTab}`} aria-labelledby={`tab-${activeTab}`}>
        {activeTab === 'settings' && <Settings />}
        {activeTab === 'invites' && <InviteCodes />}
        {activeTab === 'editions' && <Editions />}
      </section>
    </>
  );
};

export default Admin;
