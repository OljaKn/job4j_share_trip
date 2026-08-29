package domain

import (
	"time"

	"github.com/google/uuid"
)

type OutboxEvent struct {
	Id          uuid.UUID
	EventName   EventName
	AggregateId uuid.UUID
	Payload     string
	CreatedAt   time.Time
}

type EventName string

const (
	TripPublished EventName = "trip_published"
	TripCreated   EventName = "trip_created"
	TripUpdated   EventName = "trip_updated"
	TripCancelled EventName = "trip_cancelled"
	TripDeleted   EventName = "trip_deleted"
	TripComleted  EventName = "trip_completed"
)

type TripPublishedPayload struct {
	TripID string `json:"trip_id"`
}
