package domain

import (
	"time"

	"github.com/google/uuid"
)

type Trip struct {
	Id            uuid.UUID
	DriverId      uuid.UUID
	FromPoint     string
	ToPoint       string
	DepartureTime time.Time
	Seats         int
	Status        StatusTrip
	CreatedAt     time.Time
}
