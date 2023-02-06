package logic

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"wjjmjh/hermes/pkg/setting"
	"wjjmjh/hermes/pkg/util/connection"
)

// Websocket server data struct
type WsServer struct {

	// Registered users (clients)
	users map[*User]bool

	// Channels associated with server
	channels map[*Channel]bool

	// Incoming user messages
	broadcast chan []byte

	// User register requests
	register chan *User

	// User unregister requests
	unregister chan *User
}

// NewWsServer creates a new websocket server struct and returns it's address.
// Requires no parameters and returns empty channels/maps.
func NewWsServer() *WsServer {
	return &WsServer{
		broadcast:  make(chan []byte),
		register:   make(chan *User),
		unregister: make(chan *User),
		users:      make(map[*User]bool),
		channels:   make(map[*Channel]bool),
	}
}

// broadcastToUsers will send the message/messages stored in databuffer to
// all users currently registered on the server.
func (server *WsServer) broadcastToUsers(message []byte) {
	for user := range server.users {
		user.dataBuffer <- message
	}
}

// FindChannel searches through the servers channel array
// and returns
func (server *WsServer) FindChannel(p FindChannelParams) (*Channel, error) {
	var res *Channel
	for channel := range server.channels {
		if p.name != nil {
			if channel.GetName() == p.name {
				res = channel
			}
		} else if p.id != nil {
			if channel.GetID() == p.name {
				res = channel
			}
		}
	}
	if res != nil {
		return res, nil
	} else {
		return nil, errors.New("Unable to find channel")
	}
}

func (server *WsServer) findChannelByName(channelName string) *Channel {
	var res *Channel
	for channel := range server.channels {
		if *channel.GetName() == channelName {
			res = channel
			break
		}
	}

	return res
}

func (server *WsServer) findChannelByID(ID string) *Channel {
	var res *Channel
	for channel := range server.channels {
		if *channel.GetID() == ID {
			res = channel
			break
		}
	}

	return res
}

func (server *WsServer) findUserByID(ID string) *User {
	var res *User
	for user := range server.users {
		if user.UserId == ID {
			res = user
			break
		}
	}

	return res
}

// Creates a new channel and appends it to the map of
// channels stored in the websocket server.
func (server *WsServer) NewWsChannel(channelName string, private bool) *Channel {
	channel := CreateChannel(channelName, private)
	go channel.Run()
	server.channels[channel] = true
	return channel
}

func (server *WsServer) notifyUserJoined(user *User) {
	message := &Message{
		Action: UserJoinAction,
		Sender: user,
	}

	server.broadcastToUsers(MessageMarshal(*message))
}

func (server *WsServer) notifyUserLeft(user *User) {
	message := &Message{
		Action: UserLeftAction,
		Sender: user,
	}

	server.broadcastToUsers(MessageMarshal(*message))
}

func (server *WsServer) listOnlineClients(user *User) {
	for existingUser := range server.users {
		message := &Message{
			Action: UserJoinAction,
			Sender: existingUser,
		}
		user.dataBuffer <- MessageMarshal(*message)
	}
}

func (server *WsServer) addUser(user *User) {
	server.notifyUserJoined(user)
	server.listOnlineClients(user)
	server.users[user] = true
}

func (server *WsServer) removeUser(user *User) {
	if _, ok := server.users[user]; ok {
		delete(server.users, user)
		server.notifyUserLeft(user)
	}
}

// ServeWs receives a http upgrade request from a client, completes this request
// and establishes the websocket connection. It then opens up concurent read/write
// listener for the user.
func (server *WsServer) ServeWs(w http.ResponseWriter, r *http.Request) {

	name, ok := r.URL.Query()["name"]

	if !ok || len(name[0]) < 1 {
		log.Println("Url Param 'name' is missing")
		return
	}

	wsConnection, err := connection.UpgradeHTTPToWS(w, r)
	if err != nil {
		fmt.Println("Error in establishing websocket connection: ", err)
	}
	user := CreateUser(uuid.NewString(), wsConnection, server)
	log.Printf("[INFO] new client connected")

	go user.CircularWrite(setting.WsServerSetting.Ping, setting.WsServerSetting.MaxWriteWaitTime)
	go user.CircularRead(setting.WsServerSetting.MaxMessageSize, setting.WsServerSetting.Pong)

	server.register <- user
}

// Run the websocket server and listen for register/unregister requests.
// Will run continuously.
func (server *WsServer) Run() {
	fmt.Println("WebSocket Server Initialised and Running")
	for {
		select {
		// Register user
		case user := <-server.register:
			server.addUser(user)
		case user := <-server.unregister:
			server.removeUser(user)
		case message := <-server.broadcast:
			server.broadcastToUsers(message)
		}
	}
}
