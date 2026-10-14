import React from 'react';
import { Link, NavLink } from 'react-router-dom';
import { FiLogOut } from 'react-icons/fi';
import './Header.css';

const Header = ({ user, onLogout }) => (
  <header className="app-header">
    <div className="app-header-inner">
      <Link to="/" className="app-brand">FiMuVer</Link>

      <nav className="app-nav" aria-label="Hauptnavigation">
        <NavLink to="/" end>Sammlungen</NavLink>
        {user?.is_admin && <NavLink to="/admin">Admin</NavLink>}
      </nav>

      <div className="app-user">
        {user?.username && <span className="app-username">{user.username}</span>}
        <button type="button" onClick={onLogout} className="app-logout">
          <FiLogOut size={16} />
          <span>Abmelden</span>
        </button>
      </div>
    </div>
  </header>
);

export default Header;
