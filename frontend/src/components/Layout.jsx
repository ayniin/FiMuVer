import React from 'react';
import { Outlet } from 'react-router-dom';
import Header from './Header';

const Layout = ({ user, onLogout }) => (
  <div className="page">
    <Header user={user} onLogout={onLogout} />
    <main className="page-main">
      <Outlet />
    </main>
    <footer className="page-footer">
      &copy; {new Date().getFullYear()} FiMuVer – Medienverwaltung leicht gemacht
    </footer>
  </div>
);

export default Layout;
