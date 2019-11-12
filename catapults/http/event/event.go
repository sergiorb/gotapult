package event

import (
	"github.com/sergiorb/gotapult/events"
	"time"
)

type HttpEvent struct {

	Target	events.Target
	Path	string
	Query	map[string]string
	Headers	map[string]string
	Body	interface{}
	Timestamp	time.Time
	Err			interface{}
}
