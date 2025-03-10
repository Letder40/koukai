import React, { useState, useEffect } from 'react';

export interface UserData {
    documentId: string;
    name: string;
}

export interface Message {
    documentId: string;
    body: string;
    sent_by: UserData;
}

function PublicServer() {
    const [messages, setMessages] = useState<Message[]>([]);
    const [inputValue, setInputValue] = useState('');

    useEffect(() => {
        const eventSource = new EventSource('/api/listen/server/public');

        eventSource.onmessage = (event: MessageEvent) => {
            try {
                const message: Message = JSON.parse(event.data);
                console.log('Received message:', message);
                setMessages((prevMessages) => [...prevMessages, message]);
            } catch (error) {
                console.error('Error parsing SSE message:', error);
            }
        };

        eventSource.onerror = () => {
            console.error('SSE connection error');
            eventSource.close();
        };

        return () => {
            eventSource.close();
        };
    }, []); // Empty dependency array ensures this runs only on mount

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        const body = inputValue.trim();
        if (!body) return;

        try {
            const response = await fetch('/api/write/server/public', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ body })
            });

            if (response.ok) {
                setInputValue('');
            } else {
                console.error('Failed to send message:', response.statusText);
            }
        } catch (error) {
            console.error('Error sending message:', error);
        }
    };

    return (
        <div>
            <div id="messages">
                {messages.map((message) => (
                    <p key={message.documentId}>
                        {`${message.sent_by.name}: ${message.body}`}
                    </p>
                ))}
            </div>
            <form id="messageForm" onSubmit={handleSubmit}>
                <input
                    id="messageInput"
                    type="text"
                    value={inputValue}
                    onChange={(e) => setInputValue(e.target.value)}
                    placeholder="Type a message"
                />
                <button type="submit">Send</button>
            </form>
        </div>
    );
}

export default function PublciServerZone() {
    return (
        <div>
            <h1>Public Server Chat</h1>
            <PublicServer />
        </div>
    );
}

