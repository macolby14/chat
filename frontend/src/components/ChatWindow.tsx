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
    return <div className="MessageInput">
        <input className="Input" type="text" placeholder="Chat Message" id="ChatInput" />
    </div>
}