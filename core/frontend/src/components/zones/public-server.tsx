import React, { useState, useEffect } from 'react';

export interface UserData {
    documentId: string;
    username: string;
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

    const username = localStorage.getItem("username");
    return (
        <div className='w-full h-full relative'>
            <div id="messages" className='flex flex-col gap-5'>
                {messages.map((message) => {
                    if (message.sent_by.username === username) {
                        return (
                            <div key={message.documentId} className='bg-dblueDeep text-white text-wrap flex flex-col w-80 rounded-md ml-10 px-10 py-2'>
                                <div className='bold flex gap-2 items-center'>
                                    <svg width="21px" height="21px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <path d="M8 11H8.01M12 11H12.01M16 11H16.01M21 20L17.6757 18.3378C17.4237 18.2118 17.2977 18.1488 17.1656 18.1044C17.0484 18.065 16.9277 18.0365 16.8052 18.0193C16.6672 18 16.5263 18 16.2446 18H6.2C5.07989 18 4.51984 18 4.09202 17.782C3.71569 17.5903 3.40973 17.2843 3.21799 16.908C3 16.4802 3 15.9201 3 14.8V7.2C3 6.07989 3 5.51984 3.21799 5.09202C3.40973 4.71569 3.71569 4.40973 4.09202 4.21799C4.51984 4 5.0799 4 6.2 4H17.8C18.9201 4 19.4802 4 19.908 4.21799C20.2843 4.40973 20.5903 4.71569 20.782 5.09202C21 5.51984 21 6.0799 21 7.2V20Z" stroke="#ffffff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path> </g></svg>
                                    {message.sent_by.username}
                                </div>
                                <div>{message.body}</div>
                            </div>
                        )
                    } else {
                        return (
                            <div key={message.documentId} className='bg-dblue text-white text-wrap flex flex-col w-80 rounded-md ml-10 px-10 py-2'>
                                <div className='bold flex gap-2 items-center'>
                                    <svg width="21px" height="21px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <path d="M8 11H8.01M12 11H12.01M16 11H16.01M21 20L17.6757 18.3378C17.4237 18.2118 17.2977 18.1488 17.1656 18.1044C17.0484 18.065 16.9277 18.0365 16.8052 18.0193C16.6672 18 16.5263 18 16.2446 18H6.2C5.07989 18 4.51984 18 4.09202 17.782C3.71569 17.5903 3.40973 17.2843 3.21799 16.908C3 16.4802 3 15.9201 3 14.8V7.2C3 6.07989 3 5.51984 3.21799 5.09202C3.40973 4.71569 3.71569 4.40973 4.09202 4.21799C4.51984 4 5.0799 4 6.2 4H17.8C18.9201 4 19.4802 4 19.908 4.21799C20.2843 4.40973 20.5903 4.71569 20.782 5.09202C21 5.51984 21 6.0799 21 7.2V20Z" stroke="#ffffff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path> </g></svg>
                                    {message.sent_by.username}
                                </div>
                                <div>{message.body}</div>
                            </div>
                        )
                    }
                })}
            </div>
            <form id="messageForm" onSubmit={handleSubmit} className='fixed bottom-0 w-10/12 flex justify-center'>
                <input 
                    className='bg-dblack-Stronger px-10 py-2 w-6/12 text-white rounded-md mb-5'
                    id="messageInput"
                    type="text"
                    value={inputValue}
                    onChange={(e) => setInputValue(e.target.value)}
                    placeholder="Type a message"
                />
                <button type="submit" className='hidden'></button>
            </form>
        </div>
    );
}

export default function PublciServerZone() {
    return (
        <div>
            <h1 className='w-full h-20 bg-dblack-Strong text-white text-3xl bold flex justify-center items-center mb-10'>Public Server Chat</h1>
            <PublicServer />
        </div>
    );
}

