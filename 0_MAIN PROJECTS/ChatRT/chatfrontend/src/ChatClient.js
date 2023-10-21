import React, { Component } from 'react';
import './ChatClient.css';

class ChatClient extends Component {
  constructor() {
    super();
    this.state = {
      username: '',
      message: '',
      messages: [],
    };

    this.socket = new WebSocket('ws://localhost:8080');

    this.socket.onopen = () => {
      console.log('WebSocket connection is open.');
    };

    this.socket.onmessage = (event) => {
      const message = event.data;
      this.setState((prevState) => ({
        messages: [...prevState.messages, message],
      }));
    };

    this.socket.onclose = (event) => {
      if (event.wasClean) {
        console.log('WebSocket connection closed cleanly, code=' + event.code + ', reason=' + event.reason);
      } else {
        console.error('WebSocket connection abruptly closed.');
      }
    };

    this.socket.onerror = (error) => {
      console.error('WebSocket error: ' + error.message);
    };
  }

  handleUsernameChange = (event) => {
    this.setState({ username: event.target.value });
  };

  handleMessageChange = (event) => {
    this.setState({ message: event.target.value });
  };

  handleSendMessage = () => {
    if (this.socket.readyState === WebSocket.OPEN) {
      const { username, message } = this.state;
      const combinedMessage = `${username}: ${message}`;
      this.socket.send(combinedMessage);
      this.setState({ message: '' });
    } else {
      console.error('WebSocket connection is not open.');
    }
  }

  render() {
    return (
      <div className="chat-client">
        <div className="chat-window">
          <div className="message-list">
            {this.state.messages.map((message, index) => (
              <div key={index} className="message">
                {message}
              </div>
            ))}
          </div>
          <div className="message-input">
            <input
              type="text"
              placeholder="Enter your username"
              value={this.state.username}
              onChange={this.handleUsernameChange}
            />
            <input
              type="text"
              placeholder="Type a message"
              value={this.state.message}
              onChange={this.handleMessageChange}
            />
            <button onClick={this.handleSendMessage}>Send</button>
          </div>
        </div>
      </div>
    );
  }
}

export default ChatClient;
