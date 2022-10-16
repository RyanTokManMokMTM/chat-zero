package instant_messaging_system

import "sync"

//Define a Channel Map for managing each connection
//ChannelMap
type ChannelMap interface {
	Add(channel Channel)
	Remove(id string)
	Get(id string) (Channel, bool)
	All() []Channel
}

type ChannelsMapImp struct {
	channels *sync.Map //safety map
}

func NewChannelsMap(num int) ChannelMap {
	return &ChannelsMapImp{
		channels: new(sync.Map),
	}
}

func (ch *ChannelsMapImp) Add(channel Channel)           {}
func (ch *ChannelsMapImp) Remove(id string)              {}
func (ch *ChannelsMapImp) Get(id string) (Channel, bool) {}
func (ch *ChannelsMapImp) All() []Channel {
	return nil
}
