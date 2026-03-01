import React, { useState } from 'react';
import Login from './components/Login';
import Register from './components/Register';

function App() {
  const [view, setView] = useState('login');

  // TODO: Add global state management (Context/Redux) for auth status
  // TODO: Add React Router for navigation

  return (
    <div className="App">
      <h1>Simple Auth App</h1>
      <nav>
        <button onClick={() => setView('login')}>Login</button>
        <button onClick={() => setView('register')}>Register</button>
      </nav>
      
      {view === 'login' ? <Login /> : <Register />}
    </div>
  );
}

export default App;
