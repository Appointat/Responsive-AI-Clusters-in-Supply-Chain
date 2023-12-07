package event

import (
	"time"
)

type Event struct {
	EventDescription string
	EventDate        time.Time
}

var HolidayMaps = []map[string]Event{
	//EVENTS for outlet1
	{"Christmas": Event{
			EventDescription: "The upcoming Lavender Festival in the region is expected to significantly increase the demand for local specialties. We anticipate a higher demand for Olive Oil and Baguette as tourists prefer local culinary experiences. Preparing additional stock of these items is advised to meet the increased customer flow."
			EventDate:        time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
	},
	},
	//EVENTS for outlet2

	//EVENTS for outlet3

	//EVENTS for outlet4
}
