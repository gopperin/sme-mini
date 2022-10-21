package persist

import (
	mystore "github.com/gopperin/sme-mini/types/mariadb"
	"github.com/gopperin/sme-mini/types/proto"
)

// CreateEvent CreateEvent Persist
func (maria *Mariadb) CreateEvent(in proto.Event) error {
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
