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

		"Winter Sale": &Event{
			EventDescription: "The annual winter sale attracts numerous shoppers. Expect increased sales in warm beverages like Hot Chocolate and Coffee.",
			EventDate:        time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
		},
		"New Year Fitness Rush": &Event{
			EventDescription: "With the New Year, there's a surge in fitness enthusiasts. Stock up on healthy snacks and energy drinks to cater to this demographic.",
			EventDate:        time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC),
		},
		"Local Art Exhibit": &Event{
			EventDescription: "A local art exhibit is drawing in a cultured crowd. Anticipate a higher demand for gourmet coffees and pastries.",
			EventDate:        time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC),
		},
		"Snow Festival": &Event{
			EventDescription: "The annual snow festival will increase foot traffic. Stock additional winter gear and hot beverages.",
			EventDate:        time.Date(2023, 1, 20, 0, 0, 0, 0, time.UTC),
		},
		"Book Reading Club": &Event{
			EventDescription: "A book club meeting at the outlet. Prepare extra seating and offer discounts on tea and coffee.",
			EventDate:        time.Date(2023, 1, 25, 0, 0, 0, 0, time.UTC),
		},
		"Local Musician Performance": &Event{
			EventDescription: "A popular local musician will perform nearby. Expect increased evening sales.",
			EventDate:        time.Date(2023, 1, 30, 0, 0, 0, 0, time.UTC),
		},
		"Healthy Eating Workshop": &Event{
			EventDescription: "A workshop on healthy eating is scheduled, which will likely increase interest in organic products and health foods.",
			EventDate:        time.Date(2023, 1, 7, 0, 0, 0, 0, time.UTC),
		},
		"Local Sports Tournament": &Event{
			EventDescription: "A local sports tournament is expected to draw in a large crowd. Stock up on sports drinks, snacks, and quick meals.",
			EventDate:        time.Date(2023, 1, 12, 0, 0, 0, 0, time.UTC),
		},
		"Community Charity Event": &Event{
			EventDescription: "A charity event in the community will attract a lot of families. Plan for an increased demand in household essentials and snacks.",
			EventDate:        time.Date(2023, 1, 18, 0, 0, 0, 0, time.UTC),
		},
		"Local Film Festival": &Event{
			EventDescription: "The local film festival is expected to bring in movie enthusiasts. Enhance stock of popcorn, soft drinks, and movie-themed merchandise.",
			EventDate:        time.Date(2023, 1, 22, 0, 0, 0, 0, time.UTC),
		},
	},
	//EVENTS for outlet2
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

		"Winter Jazz Nights": &Event{
			EventDescription: "Jazz concerts throughout the city will likely bring a sophisticated crowd. Consider stocking up on fine wines and gourmet cheeses.",
			EventDate:        time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC),
		},
		"Culinary Cooking Classes": &Event{
			EventDescription: "Cooking classes focusing on international cuisine will attract food enthusiasts. Increase the inventory of exotic spices and specialty ingredients.",
			EventDate:        time.Date(2023, 1, 6, 0, 0, 0, 0, time.UTC),
		},
		"Local Theater Festival": &Event{
			EventDescription: "The local theater festival is set to draw in a large audience. Stock up on cultural magazines, snacks, and beverages.",
			EventDate:        time.Date(2023, 1, 11, 0, 0, 0, 0, time.UTC),
		},
		"Children's Book Fair": &Event{
			EventDescription: "A children's book fair is expected this month. Prepare for an increased demand in children's books and educational games.",
			EventDate:        time.Date(2023, 1, 14, 0, 0, 0, 0, time.UTC),
		},
		"Winter Wellness Expo": &Event{
			EventDescription: "An expo on wellness and health will take place. Anticipate a higher demand for vitamins, supplements, and organic products.",
			EventDate:        time.Date(2023, 1, 17, 0, 0, 0, 0, time.UTC),
		},
		"Photography Exhibition": &Event{
			EventDescription: "A photography exhibition will attract enthusiasts. Consider selling camera accessories and photo printing services.",
			EventDate:        time.Date(2023, 1, 21, 0, 0, 0, 0, time.UTC),
		},
		"Urban Gardening Workshop": &Event{
			EventDescription: "A workshop on urban gardening is planned. Stock up on gardening tools, seeds, and plant care books.",
			EventDate:        time.Date(2023, 1, 24, 0, 0, 0, 0, time.UTC),
		},
		"Local Craft Market": &Event{
			EventDescription: "A craft market showcasing local artisans will likely increase the sales of craft supplies and local artisan products.",
			EventDate:        time.Date(2023, 1, 27, 0, 0, 0, 0, time.UTC),
		},
		"Tech Gadget Fair": &Event{
			EventDescription: "A tech gadget fair is scheduled, attracting tech enthusiasts. Expect an increased interest in electronic gadgets and accessories.",
			EventDate:        time.Date(2023, 1, 29, 0, 0, 0, 0, time.UTC),
		},
		"Sustainable Living Seminar": &Event{
			EventDescription: "This seminar on sustainable living will draw environmentally conscious customers. Stock eco-friendly products and reusable items.",
			EventDate:        time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC),
		},
	},
	//EVENTS for outlet3
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
		"Vintage Film Screenings": &Event{
			EventDescription: "Screenings of classic films are scheduled throughout the month. Anticipate increased sales in nostalgic snacks and memorabilia.",
			EventDate:        time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC),
		},
		"Indoor Farmers Market": &Event{
			EventDescription: "The indoor farmers market will attract a health-conscious crowd. Stock up on fresh produce and organic food items.",
			EventDate:        time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC),
		},
		"Local Poet's Society Meetup": &Event{
			EventDescription: "A gathering of the local poet's society is expected. Offer a special on coffee and pastries to attract the literary crowd.",
			EventDate:        time.Date(2023, 1, 13, 0, 0, 0, 0, time.UTC),
		},
		"Board Game Tournament": &Event{
			EventDescription: "A board game tournament is drawing players of all ages. Consider stocking up on popular board games and snacks.",
			EventDate:        time.Date(2023, 1, 19, 0, 0, 0, 0, time.UTC),
		},
		"DIY Home Decor Workshop": &Event{
			EventDescription: "A workshop on DIY home decor is set to attract creative individuals. Prepare supplies like paint, brushes, and decorative items.",
			EventDate:        time.Date(2023, 1, 23, 0, 0, 0, 0, time.UTC),
		},
		"Health & Fitness Expo": &Event{
			EventDescription: "An expo focusing on health and fitness will bring in a crowd interested in sports gear and healthy eating options.",
			EventDate:        time.Date(2023, 1, 28, 0, 0, 0, 0, time.UTC),
		},
	},
	//EVENTS for outlet4
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

		"Art & Craft Fair": &Event{
			EventDescription: "An art and craft fair showcasing local talents is expected to attract a diverse crowd. Stock up on crafting materials and local artisan products.",
			EventDate:        time.Date(2023, 1, 9, 0, 0, 0, 0, time.UTC),
		},
		"Eco-Living Seminar": &Event{
			EventDescription: "A seminar on eco-friendly living practices will draw environmentally conscious consumers. Consider offering discounts on sustainable products.",
			EventDate:        time.Date(2023, 1, 16, 0, 0, 0, 0, time.UTC),
		},
		"Gourmet Cooking Show": &Event{
			EventDescription: "A gourmet cooking show is set to feature local chefs. Anticipate increased interest in high-quality kitchenware and gourmet ingredients.",
			EventDate:        time.Date(2023, 1, 26, 0, 0, 0, 0, time.UTC),
		},
		"Yoga and Wellness Retreat": &Event{
			EventDescription: "A yoga and wellness retreat is being organized nearby. Stock up on yoga mats, fitness wear, and health-focused snacks.",
			EventDate:        time.Date(2023, 1, 29, 0, 0, 0, 0, time.UTC),
		},
	},
}
