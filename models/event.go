package models

import "time"

type Event struct {
	ID int
	Name string
	Description string
	Location string
	DateTime time.Time
	UserID int 
	CreatedAt time.Time
}

var events = []Event{} 

func New() *Event{

	return &Event{}
}