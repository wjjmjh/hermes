package managers

import (
	"wjjmjh/hermes/managers/logic"
)

// Create new Channel using CreateChanel_ parameters
func (cm *ChannelManager) CreateChannel(p CreateChannel_, private bool) *logic.Channel {
	channel := logic.CreateChannel(p.ChannelName, private)
	go channel.Run()
	cm.channels[channel] = true
	return channel
}
