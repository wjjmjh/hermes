package managers

import (
	"wjjmjh/hermes/managers/logic"
)

// Create new User using CreateUser_ parameters
func CreateUser(p CreateUser_) *logic.User {
	return logic.CreateUser(p.UserName, p.conn, p.wsServer)
}
