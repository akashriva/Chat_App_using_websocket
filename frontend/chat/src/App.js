import React, { Component } from 'react';
import Header from './Components/Header';
import ChatInput from './Components/Chatinput/Chatinput'; // Corrected path
import './App.css';
import { connect, sendMsg } from './api';
import ChatHistory from './Components/ChatHistory/ChatHistory'; // Corrected path

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            chatHistory: []
        };

        // Bind the send method to this context
        this.send = this.send.bind(this);
    }

    componentDidMount() {
        connect((msg) => {
            console.log("New Message from user ");
            this.setState(prevState => ({
                chatHistory: [...prevState.chatHistory, msg]
            }));
            console.log(this.state);
        });
    }

    send(event) {
        console.log('Event triggered:', event); // Log the event to check if it's being triggered
        if (event.key === 'Enter') {
            console.log('Enter key pressed'); // Log to confirm Enter key detection
            sendMsg(event.target.value);
            event.target.value = '';
        }
    }

    render() {
        return (
            <div className="App">
                <Header />
                <ChatHistory chatHistory={this.state.chatHistory} />
                <ChatInput send={this.send} />
            </div>
        );
    }
}

export default App;
