package managers

import (
	"github.com/gorilla/websocket"
	"wjjmjh/hermes/managers/logic"
)

// Parameters for CreateChannel function
type CreateChannel_ struct {
	ChannelName string
}

type CreateUser_ struct {
	UserName string
	conn     *websocket.Conn
	wsServer *logic.WsServer
}
