import React from 'react';
import './Header.css';

const Header = ({ user, onLogout, onNavigateToAdmin }) => {
  const handleAdminClick = () => {
    if (onNavigateToAdmin) {
      onNavigateToAdmin();
    }
  };
  
  return (
    <header className="landing-header">
      <div className="header-content">
        <h1>FiMuVer</h1>
        <div className="header-buttons">
          {user?.is_admin && onNavigateToAdmin && (
            <button 
              onClick={handleAdminClick}
              className="admin-btn"
              aria-label="Admin-Panel öffnen"
            >
              Admin 
            </button>
          )}
          <button onClick={onLogout} className="logout-btn" aria-label="Abmelden">
            Logout
          </button>
        </div>
      </div>
    </header>
  );
};

export default Header;
