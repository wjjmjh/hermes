package logic

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Channel struct {
	channelID   *string
	channelName *string
	users       map[*User]bool
	threads     map[*Thread]bool
	register    chan *User
	unregister  chan *User
	broadcast   chan *Message
	Private     bool `json:"private"`
}

// Create channel method -> Used by channel_manager.go
// Created here to allow Channel to be immutable
func CreateChannel(channelName string, private bool) *Channel {

	// Initialise fields
	channelID := uuid.New().String()
	users := make(map[*User]bool)
	threads := make(map[*Thread]bool)
	register := make(chan *User)
	unregister := make(chan *User)
	broadcast := make(chan *Message)

	return &Channel{&channelID,
		&channelName,
		users,
		threads,
		register,
		unregister,
		broadcast,
		private}
}

func (channel *Channel) Run() {
	fmt.Println("Channel %s Running", channel.channelName)
	for {
		select {

		// If content exists in channel.register chan, pull it out
		case user := <-channel.register:
			channel.registerUser(user)

		case user := <-channel.unregister:
			channel.unregisterUser(user)

		case message := <-channel.broadcast:
			channel.broadcastToUsers(MessageMarshal(*message))
		}
	}
}

/*
	Methods to get channel fields
*/

func (channel *Channel) GetID() *string {
	return channel.channelID
}

func (channel *Channel) GetName() *string {
	return channel.channelName
}

func (channel *Channel) GetAllUsers() map[*User]bool {
	return channel.users

}

func (channel *Channel) getThreads() map[*Thread]bool {
	return channel.threads
}

// Description:    Gets all users in a specific Channel.
// Input:          getUsersParams struct (logicParameters.go).
// Returns:        List of pointers to user username or userID and error.
func (channel *Channel) GetUsers(p GetUsersParams_) ([]*string, error) {
	var users []*string
	var errorMsg error = nil

	// Loop through and append to return array all users satisfying users: True
	for User, value := range channel.users {
		if value {
			if p.ReturnType == "username" {
				users = append(users, User.GetUsername())
			} else if p.ReturnType == "userId" {
				id := User.GetID()
				users = append(users, &id)
			} else {
				errorMsg = errors.New(fmt.Sprintf("Invalid getChannelUsers() input parameter: %s", p.ReturnType))
				users = nil
				break
			}
		}
	}
	return users, errorMsg
}

/*
	Channel modification methods
*/
func (channel *Channel) UpdateName(p UpdateName_) {
	channel.channelName = &p.UpdatedName
}

/*
	Channel Threads Methods
*/

// Create a new thread within a channel. Function must be called from an instantiated channel.
func (channel *Channel) CreateThread() *Thread {
	threadID := uuid.New().String()
	users := make(map[*User]bool)
	return &Thread{&threadID, users, channel}
}

// Adds a user to a room
func (channel *Channel) registerUser(user *User) {
	// Send join message to room
	message := &Message{Action: JoinChannelAction,
		Message: fmt.Sprintf("Welcome %s to the %s!", *user.username, *channel.channelName)}
	channel.broadcastToUsers(MessageMarshal(*message))

	// Register user
	channel.users[user] = true

	// Notify channel members that someone joined
	channel.notifyUserJoined(user)
}

// Removes user from a room
func (channel *Channel) unregisterUser(user *User) {
	// Send leave message to room
	message := &Message{Action: "User Left",
		Message: fmt.Sprintf("%s left the channel", *user.username)}
	channel.broadcastToUsers(MessageMarshal(*message))

	// Remove from room
	if _, ok := channel.users[user]; ok {
		delete(channel.users, user)
	}

}

func (channel *Channel) broadcastToUsers(message []byte) {
	for user := range channel.users {
		user.dataBuffer <- message
	}
}

// Notifies the room that the user with username x joined.
func (channel *Channel) notifyUserJoined(user *User) {
	const welcomeMessage = "%s joined the room"
	message := &Message{
		Action:  SendMessageAction,
		Target:  channel,
		Message: fmt.Sprintf(welcomeMessage, *user.GetUsername()),
	}

	// Send to all the users of the channel.
	channel.broadcastToUsers(MessageMarshal(*message))
}
