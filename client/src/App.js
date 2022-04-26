import React from 'react';
import logo from './logo.svg';
import './App.css';
import PingComponent2 from './PingComponent2';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Deploy React + Go to Heroku using Docker
        </p>
        <PingComponent2/>

      </header>
    </div>
  );
}

export default App;
