import React, { useState } from 'react';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';
import Auth from './pages/Auth';
import Landing from './pages/Landing';
import Admin from './pages/Admin';
import CollectionPage from './pages/CollectionPage';
import ItemDetailPage from './pages/ItemDetailPage';
import Layout from './components/Layout';
import FeedbackProvider from './feedback/FeedbackProvider';
import { getCurrentUser, logout } from './services/userapi';

function App() {
  const [user, setUser] = useState(getCurrentUser);

  const handleLoginSuccess = () => {
    setUser(getCurrentUser());
  };

  const handleLogout = () => {
    logout();
    setUser(null);
  };

  return (
    <FeedbackProvider>
      <BrowserRouter>
        {!user ? (
          <Routes>
            <Route path="*" element={<Auth onLoginSuccess={handleLoginSuccess} />} />
          </Routes>
        ) : (
          <Routes>
            <Route element={<Layout user={user} onLogout={handleLogout} />}>
              <Route index element={<Landing />} />
              <Route path="collections/:collectionId" element={<CollectionPage />} />
              <Route path="collections/:collectionId/items/:itemId" element={<ItemDetailPage />} />
              <Route
                path="admin"
                element={user.is_admin ? <Admin /> : <Navigate to="/" replace />}
              />
              <Route path="*" element={<Navigate to="/" replace />} />
            </Route>
          </Routes>
        )}
      </BrowserRouter>
    </FeedbackProvider>
  );
}

export default App;
