import React, { useState } from 'react';
import UserAPI from '../services/userapi';
import { useFeedback } from '../feedback/context';
import './Auth.css';

const Auth = ({ onLoginSuccess }) => {
  const { toast } = useFeedback();
  const [isLogin, setIsLogin] = useState(true);
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [inviteCode, setInviteCode] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const switchMode = () => {
    setIsLogin(!isLogin);
    setError('');
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      if (isLogin) {
        await UserAPI.loginUser(email, password);
        onLoginSuccess();
      } else {
        await UserAPI.registerUser(username, email, password, inviteCode);
        setIsLogin(true);
        setUsername('');
        setPassword('');
        setInviteCode('');
        toast('Registrierung erfolgreich – du kannst dich jetzt anmelden.');
      }
    } catch (err) {
      setError(err.message || 'Ein Fehler ist aufgetreten.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="auth-page">
      <div className="auth-card card">
        <div className="auth-brand">FiMuVer</div>
        <h1>{isLogin ? 'Anmelden' : 'Konto erstellen'}</h1>
        <p className="auth-subtitle">
          {isLogin ? 'Willkommen zurück bei deiner Mediensammlung.' : 'Starte deine eigene Mediensammlung.'}
        </p>

        <form onSubmit={handleSubmit}>
          {!isLogin && (
            <div className="field">
              <label htmlFor="username">Benutzername</label>
              <input
                className="input"
                type="text"
                id="username"
                autoComplete="username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                required
              />
            </div>
          )}
          <div className="field">
            <label htmlFor="email">E-Mail</label>
            <input
              className="input"
              type="email"
              id="email"
              autoComplete="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
          <div className="field">
            <label htmlFor="password">Passwort</label>
            <input
              className="input"
              type="password"
              id="password"
              autoComplete={isLogin ? 'current-password' : 'new-password'}
              minLength={isLogin ? undefined : 12}
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
            {!isLogin && <span className="field-hint">Mindestens 12 Zeichen</span>}
          </div>
          {!isLogin && (
            <div className="field">
              <label htmlFor="inviteCode">
                Einladungscode <span className="field-hint">(optional)</span>
              </label>
              <input
                className="input auth-code-input"
                type="text"
                id="inviteCode"
                value={inviteCode}
                onChange={(e) => setInviteCode(e.target.value.toUpperCase())}
                placeholder="z.B. AB12CD34"
                maxLength={8}
              />
            </div>
          )}

          {error && <p className="alert alert--error" role="alert">{error}</p>}

          <button type="submit" className="btn btn--primary btn--lg btn--block" disabled={loading}>
            {loading ? 'Bitte warten…' : isLogin ? 'Anmelden' : 'Registrieren'}
          </button>
        </form>

        <div className="auth-switch">
          <span>{isLogin ? 'Noch kein Konto?' : 'Bereits ein Konto?'}</span>
          <button type="button" className="auth-switch-link" onClick={switchMode}>
            {isLogin ? 'Registrieren' : 'Zur Anmeldung'}
          </button>
        </div>
      </div>
    </div>
  );
};

export default Auth;
