package logic

import (
	"errors"
	"fmt"
)

type Thread struct {
	threadID *string
	users    map[*User]bool
	channel  *Channel
}

func (thread *Thread) GetID() *string {
	return thread.threadID
}

func (thread *Thread) GetParentChannel() *Channel {
	return thread.channel
}

func (thread *Thread) GetUsers() map[*User]bool {
	return thread.users
}

// Description:    Gets all users in a specific Thread.
// Input:          getUsersParams struct (logicParameters.go).
// Returns:        List of pointers to user username or userID and error.
func (thread *Thread) GetAllUsers(p GetUsersParams_) ([]*string, error) {
	var users []*string
	var errorMsg error = nil

	// Loop through and append to return array all users satisfying users: True
	for User, value := range thread.users {
		if value {
			if p.ReturnType == "username" {
				users = append(users, User.GetUsername())
			} else if p.ReturnType == "userId" {
				id := User.GetID()
				users = append(users, &id)
			} else {
				errorMsg = errors.New(fmt.Sprintf("Invalid getThreadUsers() input parameter: %s", p.ReturnType))
				users = nil
				break
			}
		}

	}
	return users, errorMsg
}
