package mariadb

import (
	"github.com/jinzhu/gorm"
)

// GudpEvent GudpEvent
type GudpEvent struct {
	gorm.Model
	GudpEventBase
}

// GudpEventBase GudpEventBase
type GudpEventBase struct {
	EventID       int64  `gorm:"index:idx_event_id" form:"event_id" json:"event_id"`              // event_id
	EventType     string `gorm:"size:100;default:''" form:"event_type" json:"event_type"`         // event_type
	AggregateID   int64  `gorm:"index:idx_aggregate_id" form:"aggregate_id" json:"aggregate_id"`  // aggregate_id
	AggregateType string `gorm:"size:100;default:''" form:"aggregate_type" json:"aggregate_type"` // aggregate_type
	EventData     string `gorm:"size:1000;default:''" form:"event_data" json:"event_data"`        // event_data
	ExData        string `gorm:"size:1000;default:''" form:"ex_data" json:"ex_data"`              // ex_data
	Channel       string `gorm:"size:100;default:''" form:"channel" json:"channel"`               // channel
	Stream        string `gorm:"size:100;default:''" form:"stream" json:"stream"`                 // stream
}
