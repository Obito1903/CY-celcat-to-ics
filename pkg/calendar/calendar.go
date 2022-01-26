package calendar

import (
	"time"

	"github.com/Obito1903/CY-celcat/pkg/celcat"
)

type Event struct {
	Id         string
	Start      time.Time
	End        time.Time
	AllDay     bool
	Category   string
	Subjects   []string
	Location   []string
	Professors []string
	Notes      string
}

type Calendar struct {
	Name   string
	Events []Event
}

func eventFromCelcat(celcatEvent celcat.CelcatCalEvent) Event {
	var event Event
	event.Id = celcatEvent.Id
	event.Start, _ = time.ParseInLocation("2006-01-02T15:04:05", celcatEvent.Start, time.Local)
	event.End, _ = time.ParseInLocation("2006-01-02T15:04:05", celcatEvent.End, time.Local)
	event.AllDay = celcatEvent.AllDay
	for _, element := range celcatEvent.Elements {
		switch element.Label {
		case "Catégorie":
			event.Category = element.Content
		case "Matière", "Name":
			event.Subjects = append(event.Subjects, element.Content)
		case "Notes":
			event.Notes = element.Content
		default:
			switch element.EntityType {
			case 102:
				event.Location = append(event.Location, element.Content)
			case 101:
				event.Professors = append(event.Professors, element.Content)
			}
		}
	}
	return event
}

func FromCelcat(celcatCalendar []celcat.CelcatCalEvent, name string) Calendar {
	var calendar Calendar
	calendar.Name = name
	for _, event := range celcatCalendar {
		calendar.Events = append(calendar.Events, eventFromCelcat(event))
	}
	return calendar
}
