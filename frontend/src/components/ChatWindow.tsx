import React, { useState } from 'react'
import './ChatWindow.css'

export function ChatWindow() {
    return <div className="ChatWindow">
        <MessageHistory />
        <MessageInput />
    </div>
}

export function MessageHistory() {
    return <div className="MessageHistory">
        Message History
    </div>
}


export function MessageInput() {
    const [input, setInput] = useState("");

    function handleKeyDown(e: React.KeyboardEvent<HTMLInputElement>) {
        if(e.code !== 'Enter'){ return; }
        setInput("");
        console.log(`TODO: Sent message ${input}`);
    }

    return <div className="MessageInput">
        <input className="Input" type="text" placeholder="Chat Message" id="ChatInput" 
        onKeyDown={handleKeyDown}
        value={input}
        onChange = {(e) => setInput(e.target.value)} />
    </div>
}