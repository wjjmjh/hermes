package logic

import "time"

/*
 Parameters for all functions relating to logic package
*/

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Update channel name
type UpdateName_ struct {
	UpdatedName string
}

// Parameters for CreateChannel function
type CreateChannel_ struct {
	ChannelName string
}

// Parameters for getUsers function
type GetUsersParams_ struct {
	ReturnType string // [userId, username
}

// Conn default:
// maxMessageSize: 1000
// pong: 60 * time.Second
// ping: (pong * 9) / 10 < pong
// maxWriteWaitTime 10 * time.Second
type Conn struct {
	// Maximum message size allowed from peer.
	maxMessageSize int64
	// PingPong: Two processes send packets of information back and forth a number of times.
	pong time.Duration
	ping time.Duration
	// Maximum waiting time when writing to peers.
	maxWriteWaitTime time.Duration
}

/*
	server_logic
*/
type FindChannelParams struct {
	name *string
	id   *string
}
