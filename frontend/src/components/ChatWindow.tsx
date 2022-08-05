import React, { useState } from "react";
import { useRecoilValue } from "recoil";
import { wsState } from "../atoms";
import "./ChatWindow.css";

export function ChatWindow() {
  return (
    <div className="ChatWindow">
      <MessageHistory />
      <MessageInput />
    </div>
  );
}

export function MessageHistory() {
  return <div className="MessageHistory">Message History</div>;
}

export function MessageInput() {
  const [input, setInput] = useState("");
  const ws = useRecoilValue(wsState);

  function handleKeyDown(e: React.KeyboardEvent<HTMLInputElement>) {
    if (e.code !== "Enter") {
      return;
    }
    const inputToSend = input;
    setInput("");
    if (!ws) {
      console.error("WS is not set and user is trying to send a mesage");
      return;
    }
    console.log(`Sending message: ${inputToSend}`);
    ws.send(inputToSend);
  }

  return (
    <div className="MessageInput">
      <input
        className="Input"
        type="text"
        placeholder="Chat Message"
        id="ChatInput"
        onKeyDown={handleKeyDown}
        value={input}
        onChange={(e) => setInput(e.target.value)}
      />
    </div>
  );
}
