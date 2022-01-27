package persist

import (
	mystore "types/mariadb"
	"types/pb"
)

// CreateEvent CreateEvent Persist
func (maria *Mariadb) CreateEvent(in pb.Event) error {
	var event mystore.GudpEvent
	event.EventID = in.EventId
	event.EventType = in.EventType
	event.AggregateID = in.AggregateId
	event.AggregateType = in.AggregateType
	event.EventData = in.EventData
	event.ExData = in.ExData
	event.Channel = in.Channel
	event.Stream = in.Stream
	return maria.db.Create(&event).Error
}
