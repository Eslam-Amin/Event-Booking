package models

import "time"

type Event struct {
	ID int
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string	`binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int 
	CreatedAt time.Time
}

var events = []Event{} 

func (event *Event) Save(){
	// later: add it to a database 
	events = append(events, *event)
}

func GetAllEvents() []Event{
	return events
}

func New() *Event{

	return &Event{}
}