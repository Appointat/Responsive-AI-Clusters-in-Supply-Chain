package event

import (
	"time"
)

type Event struct {
	EventDescription string
	EventDate        time.Time
}

var HolidayMaps = []map[string]*Event{
	//EVENTS for outlet1
	{
		"Christmas": &Event{
			EventDescription: "The upcoming Lavender Festival in the region is expected to significantly increase the demand for local specialties. We anticipate a higher demand for Olive Oil and Baguette as tourists prefer local culinary experiences. Preparing additional stock of these items is advised to meet the increased customer flow.",
			EventDate:        time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
		},
		//Add Holloween
		"Halloween": &Event{
			EventDescription: "As Halloween approaches, there is a notable increase in demand for festive items. Black Tea, known for its warm and comforting qualities, becomes particularly popular during this season. Additionally, Baguette, often used in creative Halloween-themed recipes, is expected to see higher sales. Stocking up on these items is recommended to cater to the seasonal demand.",
			EventDate:        time.Date(2023, 10, 31, 0, 0, 0, 0, time.UTC),
		},
		"New Year's Day": &Event{
			EventDescription: "The celebration of New Year's Day leads to an increased demand for festive items. Products like Black Tea and Olive Oil, often used in traditional New Year's meals, are expected to be in higher demand. Stocking these items in advance is recommended.",
			EventDate:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"Labour Day": &Event{
			EventDescription: "Labour Day is typically a time for relaxation and family gatherings. There's a surge in demand for comfort foods like Baguette and Olive Oil for home cooking. Preparing additional stock is advisable.",
			EventDate:        time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
		},
		"Victory in Europe Day": &Event{
			EventDescription: "This day commemorates the end of WWII in Europe. Traditional celebrations might increase the demand for items like Baguette and Manchego Cheese, which are popular at social gatherings.",
			EventDate:        time.Date(2023, 5, 8, 0, 0, 0, 0, time.UTC),
		},
		"Bastille Day": &Event{
			EventDescription: "Bastille Day, a national holiday in France, often involves public festivities and private celebrations. Expect higher sales of items like Olive Oil and Baguette, staples in French culinary celebrations.",
			EventDate:        time.Date(2023, 7, 14, 0, 0, 0, 0, time.UTC),
		},
		"Assumption of Mary": &Event{
			EventDescription: "This religious holiday often involves family gatherings and feasts. Olive Oil and Black Tea, as popular items for such occasions, are likely to see an increase in demand.",
			EventDate:        time.Date(2023, 8, 15, 0, 0, 0, 0, time.UTC),
		},

		//Add accident fire
		"Fire": &Event{
			EventDescription: "There was a fire in the outlet, and the outlet will be closed for 3 days. So the outlet will not be able to receive the replenishment.",
			EventDate:        time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
		},
	},
	//EVENTS for outlet2
	{
		"Christmas": &Event{
			EventDescription: "The upcoming Lavender Festival in the region is expected to significantly increase the demand for local specialties. We anticipate a higher demand for Olive Oil and Baguette as tourists prefer local culinary experiences. Preparing additional stock of these items is advised to meet the increased customer flow.",
			EventDate:        time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
		},
	},
	//EVENTS for outlet3
	{
		"Christmas": &Event{
			EventDescription: "The upcoming Lavender Festival in the region is expected to significantly increase the demand for local specialties. We anticipate a higher demand for Olive Oil and Baguette as tourists prefer local culinary experiences. Preparing additional stock of these items is advised to meet the increased customer flow.",
			EventDate:        time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
		},
	},
	//EVENTS for outlet4
	{
		"Christmas": &Event{
			EventDescription: "The upcoming Lavender Festival in the region is expected to significantly increase the demand for local specialties. We anticipate a higher demand for Olive Oil and Baguette as tourists prefer local culinary experiences. Preparing additional stock of these items is advised to meet the increased customer flow.",
			EventDate:        time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
		},
	},
}
