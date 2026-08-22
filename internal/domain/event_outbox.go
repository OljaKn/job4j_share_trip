package domain

import (
	"time"

	"github.com/google/uuid"
)

type OutboxEvent struct {
	Id          uuid.UUID
	EventName   string
	AggregateId uuid.UUID
	Payload     string
	CreatedAt   time.Time
}
