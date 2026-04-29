import React, { useState, useEffect } from 'react';
import Auth from './pages/Auth';
import Landing from './pages/Landing';
import Admin from './pages/Admin';
import { getCurrentUser, logout } from './services/userapi';
import './App.css';

function App() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [currentPage, setCurrentPage] = useState('landing');

  useEffect(() => {
    const currentUser = getCurrentUser();
    if (currentUser) {
      setUser(currentUser);
    }
    setLoading(false);
  }, []);

  const handleLoginSuccess = () => {
    setUser(getCurrentUser());
    setCurrentPage('landing');
  };

  const handleLogout = () => {
    logout();
    setUser(null);
    setCurrentPage('landing');
  };

  const handleNavigateToAdmin = () => {
    if (user?.is_admin) {
      setCurrentPage('admin');
    }
  };

  const handleNavigateBack = () => {
    setCurrentPage('landing');
  };

  if (loading) {
    return <div>Wird geladen...</div>;
  }

  return (
    <div className="App">
      {!user ? (
        <Auth onLoginSuccess={handleLoginSuccess} />
      ) : currentPage === 'admin' ? (
        <Admin 
          user={user} 
          onLogout={handleLogout}
          onNavigateBack={handleNavigateBack}
        />
      ) : (
        <Landing 
          user={user} 
          onLogout={handleLogout}
          onNavigateToAdmin={handleNavigateToAdmin}
        />
      )}
    </div>
  );
}

export default App;
