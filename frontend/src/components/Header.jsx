import React from 'react';
import './Header.css';

const Header = ({ user, onLogout }) => {
  return (
    <header className="landing-header">
      <div className="header-content">
        <h1>FiMuVer</h1>
        <div className="header-buttons">
          <button onClick={onLogout} className="logout-btn">
            Logout
          </button>
          <button 
            onClick={() => alert('Admin-Funktionalität hier implementieren')}
            className="admin-btn"
          >
            Admin 
          </button>
        </div>
      </div>
    </header>
  );
};

export default Header;
