import React from 'react';
import './Reset.css'
import './Global.css'
import './App.css';
import { ChatWindow } from "./components/ChatWindow"
import { UserWindow } from './components/UserWindow';

function App() {
  return (
    <div className="App">
      <div className="ChatWindowContainer">
        <ChatWindow />
      </div>
      <div className="UserWindowContainer">
        <UserWindow />
      </div>
    </div>
  );
}

export default App;
