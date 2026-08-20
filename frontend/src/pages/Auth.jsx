import React, { useState } from 'react';
import UserAPI from '../services/userapi';
import './Auth.css';

const Auth = ({ onLoginSuccess }) => {
  const [isLogin, setIsLogin] = useState(true);
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [inviteCode, setInviteCode] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

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
        setEmail('');
        setPassword('');
        setInviteCode('');
      }
    } catch (err) {
      setError(err.message || 'Ein Fehler ist aufgetreten.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="auth-container">
      <div className="auth-form">
        <h2>{isLogin ? 'Login' : 'Registrieren'}</h2>
        <form onSubmit={handleSubmit}>
          {!isLogin && (
            <div className="form-group">
              <label htmlFor="username">Benutzername</label>
              <input
                type="text"
                id="username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                required
              />
            </div>
          )}
          <div className="form-group">
            <label htmlFor="email">E-Mail</label>
            <input
              type="email"
              id="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="password">Passwort</label>
            <input
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          {!isLogin && (
            <div className="form-group">
              <label htmlFor="inviteCode">Invite Code <span className="optional-label">(optional)</span></label>
              <input
                type="text"
                id="inviteCode"
                value={inviteCode}
                onChange={(e) => setInviteCode(e.target.value.toUpperCase())}
                placeholder="z.B. AB12CD34"
                maxLength={8}
              />
            </div>
          )}
          {error && <p className="error-message">{error}</p>}
          <button type="submit" disabled={loading}>
            {loading ? 'Wird geladen...' : isLogin ? 'Login' : 'Registrieren'}
          </button>
        </form>
        <div className="auth-switch">
          <span>{isLogin ? 'Noch kein Konto?' : 'Bereits ein Konto?'}</span>
          <button className="auth-switch-link" onClick={() => setIsLogin(!isLogin)}>
            {isLogin ? 'Registrieren' : 'Zum Login'}
          </button>
        </div>
      </div>
    </div>
  );
};

export default Auth;
