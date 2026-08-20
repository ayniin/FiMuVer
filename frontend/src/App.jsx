import React, { useState, useEffect } from 'react';
import Auth from './pages/Auth';
import Landing from './pages/Landing';
import Admin from './pages/Admin';
import { getCurrentUser, logout } from './services/userapi';
import './App.css';
import CollectionPage from './pages/CollectionPage';
import ItemDetailPage from './pages/ItemDetailPage';

function App() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [currentPage, setCurrentPage] = useState('landing');
  const [currentCollection, setCurrentCollection] = useState(null);
  const [currentItem, setCurrentItem] = useState(null);

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

  const handleNavigateToCollection = (collection) => {
    setCurrentCollection(collection);
    setCurrentPage('collection');
  };

  const handleNavogatioToItem = (item) =>{

    setCurrentItem(item);
    setCurrentPage('item');
  }

  const handleNavigateBackToCollection = () => {
    setCurrentItem(null);
    setCurrentPage('collection');
  } 
  if (loading) {
    return <div>Wird geladen...</div>;
  }

  return (
    <div className="App">
      {!user ? (
        <Auth onLoginSuccess={handleLoginSuccess} />
      ) : currentPage === 'admin' ? (
        <Admin 
          onLogout={handleLogout}
          onNavigateBack={handleNavigateBack}
        />
      ) : currentPage === 'collection' ? (
        <CollectionPage
          collection={currentCollection}
          onNavigateBack={handleNavigateBack}
          onNavigateToItem={handleNavogatioToItem}
          user={user}
          onLogout={handleLogout}
        />
      ) : currentPage === 'item' ? (
        <ItemDetailPage
          item={currentItem}
          onNavigateBack={handleNavigateBackToCollection}
        />
      ) : (
        <Landing 
          user={user} 
          onLogout={handleLogout}
          onNavigateToAdmin={handleNavigateToAdmin}
          onNavigateToCollection={handleNavigateToCollection}
        />
      )}
    </div>
  );
}

export default App;
