package logic

import (
	"encoding/json"
	"log"
)

// Message actions
const SendMessageAction = "send-message"
const JoinChannelAction = "join-channel"
const LeaveChannelAction = "leave-channel"
const UserJoinAction = "user-join"
const UserLeftAction = "user-left"
const JoinPrivateChannelAction = "join-private-channel"
const ChannelJoinedAction = "channel-joined"

type Message struct {
	// Message request type
	Action string `json:"action"`

	// Actual Message
	Message string `json:"message"`

	// Message target (room or user)
	Target *Channel `json:"target"`

	// User sending the message
	Sender *User `json:"sender"`
}

func MessageMarshal(msg Message) []byte {
	_json, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	return _json
}

func MessageUnmarshal(msg []byte) *Message {
	var unmarshalledMessage Message
	err := json.Unmarshal(msg, &unmarshalledMessage)
	if err != nil {
		log.Println("Error with unmarshal", err)
		return nil
	} else {
		return &unmarshalledMessage
	}
}
