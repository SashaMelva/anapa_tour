package autenficationmodel

import "time"

type User struct {
	Id           uint32
	FirstName    string
	MiddelName   string
	LastName     string
	DateBerthday time.Time
}
