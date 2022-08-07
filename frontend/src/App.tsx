import React, { useEffect, useState } from "react";
import "./Reset.css";
import "./Global.css";
import "./App.css";
import { ChatWindow } from "./components/ChatWindow";
import { UserWindow } from "./components/UserWindow";
import { useSetRecoilState } from "recoil";
import { wsState } from "./atoms";
function App() {
  const [errorMessage, setErrorMessage] = useState("");
  const setWs = useSetRecoilState(wsState);

  useEffect(() => {
    let conn: WebSocket | null;
    if (window["WebSocket"]) {
      conn = new WebSocket("ws://127.0.0.1:8000/ws");
      console.log("Creating a new websocket connection");
      conn.onclose = function (evt) {
        console.error("Lost connection to server.");
        setErrorMessage("Lost connection to server.");
      };
      conn.onmessage = function (evt) {
        const messages = evt.data.split("\n");
        console.log("Received messages: " + messages);
      };
      setWs(conn);
      setErrorMessage("");
    } else {
      console.error("This browser does not support websockets");
      setErrorMessage("This browser does not support websockets");
    }

    return () => {
      conn?.close();
      setWs(null);
    };
  }, [setWs]);

  return (
    <>
      {errorMessage && <div>{errorMessage}</div>}
      <div className="App">
        <div className="ChatWindowContainer">
          <ChatWindow />
        </div>
        <div className="UserWindowContainer">
          <UserWindow />
        </div>
      </div>
    </>
  );
}

export default App;
