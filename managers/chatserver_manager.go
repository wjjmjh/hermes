package managers

import (
	"flag"
	"fmt"
	"net/http"
	"wjjmjh/hermes/managers/logic"
	"wjjmjh/hermes/pkg/setting"
)

/*
Define Fundamental Types
*/

// Struct that bundles and acts as a controller between the users and channels
type ChatServerManager struct {
	UserManager    *UserManager
	ChannelManager *ChannelManager
	wsServer       *logic.WsServer
}

// Handles all business logic relating to a User
type UserManager struct {
	userManagerId *string
	users         map[*logic.User]bool
}

// Handles all business logic relating to a Channel
type ChannelManager struct {
	channelManagerID *string
	channels         map[*logic.Channel]bool
}

// Initialises managers and defines object design pattern`
func InitialiseManager() *ChatServerManager {

	controller := new(ChatServerManager)

	// Initialise the websocketServer
	server := logic.NewWsServer()
	controller.wsServer = server

	// Initialise child structs
	um := new(UserManager)
	cm := new(ChannelManager)

	controller.UserManager = um
	controller.ChannelManager = cm

	return controller

}

// RunWsServer starts the websocket server, and beings listening on the port
// specified in config. On client connection/upgrade request, it will attempt
// to establish a websocket handshake.
func (chatManager *ChatServerManager) RunWsServer() {
	var addr = flag.String("addr", setting.WsServerSetting.Port, "http service address")

	// Start websocket register listener
	go chatManager.wsServer.Run()

	// Start websocket read/write pump listening
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chatManager.wsServer.ServeWs(w, r)
	})

	// Port listening
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		fmt.Println("ListenAndServe Error: ", err)
	}
}
