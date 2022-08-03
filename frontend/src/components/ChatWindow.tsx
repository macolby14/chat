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
        Message Input
    </div>
}