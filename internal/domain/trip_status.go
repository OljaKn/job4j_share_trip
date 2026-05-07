package domain

type StatusTrip string

const (
	Draft     StatusTrip = "draft"
	Public    StatusTrip = "published"
	Canceled  StatusTrip = "canceled"
	Completed StatusTrip = "completed"
)
