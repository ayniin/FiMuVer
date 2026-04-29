import React, { useState, useEffect } from 'react';
import Auth from './pages/Auth';
import Landing from './pages/Landing';
import { getCurrentUser, logout } from './services/api';
import './App.css';

function App() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const currentUser = getCurrentUser();
    if (currentUser) {
      setUser(currentUser);
    }
    setLoading(false);
  }, []);

  const handleLoginSuccess = () => {
    setUser(getCurrentUser());
  };

  const handleLogout = () => {
    logout();
    setUser(null);
  };

  if (loading) {
    return <div>Wird geladen...</div>;
  }

  return (
    <div className="App">
      {user ? (
        <Landing 
          user={user} 
          onLogout={handleLogout}
        />
      ) : (
        <Auth onLoginSuccess={handleLoginSuccess} />
      )}
    </div>
  );
}

export default App;
