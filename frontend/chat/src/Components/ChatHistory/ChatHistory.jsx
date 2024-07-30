import React,{Comment, Component} from "react";
import Message from "../Message";
import './ChatHistory.scss'


class ChatHistory extends Component{
    render(){
        const message = this.props.chatHistory.map(msg =><Message key={msg.timeStamp} message ={msg.data}/>)
        console.log(message)
        return(
            <div className="ChatHistory">
                <h2>Chat History</h2>
                {message}
            </div>
        )
    }
}

export default ChatHistory;