package persist

import (
	mystore "types/mariadb"
	"types/pb"
)

// CreateEvent CreateEvent Persist
func (maria *Mariadb) CreateEvent(in pb.Event) error {
	var _event mystore.GudpEvent
	_event.EventID = in.EventId
	_event.EventType = in.EventType
	_event.AggregateID = in.AggregateId
	_event.AggregateType = in.AggregateType
	_event.EventData = in.EventData
	_event.ExData = in.ExData
	_event.Channel = in.Channel
	_event.Stream = in.Stream
	return maria.db.Create(&_event).Error
}
