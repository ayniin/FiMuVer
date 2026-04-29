import React from 'react';
import './Landing.css';

const Landing = ({ user, onLogout }) => {
  return (
    <div className="landing-container">
      <header className="landing-header">
        <div className="header-content">
          <h1>FiMuVer</h1>
          <button onClick={onLogout} className="logout-btn">
            Logout
          </button>
          <button onClick={() => alert('Admin-Funktionalität hier implementieren')}>
            Admin 
          </button>
        </div>
      </header>

      <main className="landing-main">
        <div className="welcome-section">
          <h2>Willkommen, {user?.username || 'Benutzer'}!</h2>
          <p>Verwalte deine Mediensammlung ganz einfach.</p>
        </div>

        <div className="features-section">
          <div className="feature-card">
            <h3>🎬 Deine Sammlung</h3>
            <p>Verwalte deine Filme, Musik und andere Medien an einem Ort.</p>
          </div>

          <div className="feature-card">
            <h3>📱 Einfache Verwaltung</h3>
            <p>Füge neue Medien hinzu, bearbeite und organisiere deine Sammlung.</p>
          </div>

          <div className="feature-card">
            <h3>🎯 Übersicht</h3>
            <p>Behalte alle deine Medien im Überblick mit einer übersichtlichen Oberfläche.</p>
          </div>
        </div>
      </main>

      <footer className="landing-footer">
        <p>&copy; 2026 FiMuVer - Medienverwaltung leicht gemacht</p>
      </footer>
    </div>
  );
};

export default Landing;
