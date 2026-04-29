import React from 'react';
import './Header.css';

const Header = ({ user, onLogout, onNavigateToAdmin }) => {
  // Debug-Output
  console.log('Header user prop:', user);
  console.log('is_admin check:', user?.is_admin);
  console.log('onNavigateToAdmin function:', onNavigateToAdmin);
  
  const handleAdminClick = () => {
    console.log('Admin button clicked!');
    if (onNavigateToAdmin) {
      console.log('Calling onNavigateToAdmin...');
      onNavigateToAdmin();
    } else {
      console.log('ERROR: onNavigateToAdmin is undefined!');
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
            >
              Admin 
            </button>
          )}
          <button onClick={onLogout} className="logout-btn">
            Logout
          </button>
        </div>
      </div>
    </header>
  );
};

export default Header;
