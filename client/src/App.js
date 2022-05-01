import React from 'react';
import PingComponent from './PingComponent';
import SamplePage from './components/SamplePage';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <p>
          Deploy React + Go to Heroku using Docker
        </p>
        <PingComponent/>
        <SamplePage/>

      </header>
    </div>
  );
}

export default App;
