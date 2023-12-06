package main

import (
	"time"

	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/outlet"
)

func main() {
	outlet.INIT()
	time.Sleep(10 * time.Second) // TODO: Keep main function running
}

// HolidayEvents maps event names to their dates.
var holidayMaps = []map[string]time.Time{
	{
		"New Year's Day":    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		"Valentine's Day":   time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
		"St. Patrick's Day": time.Date(2023, 3, 17, 0, 0, 0, 0, time.UTC),
		// ... (other events) ...
		"Chinese National Day": time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		"Event_11_25":          time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
		"Event_11_26":          time.Date(2023, 11, 26, 0, 0, 0, 0, time.UTC),
	},
	// Copy the same map for the sake of this example
	{
		"New Year's Day":    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		"Valentine's Day":   time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
		"St. Patrick's Day": time.Date(2023, 3, 17, 0, 0, 0, 0, time.UTC),
		// ... (other events) ...
		"Chinese National Day": time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		"Event_11_25":          time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
		"Event_11_26":          time.Date(2023, 11, 26, 0, 0, 0, 0, time.UTC),
	},
	// Repeat the same map 3 more times to simulate different maps
	// In real usage, these would likely be different
	{
		"New Year's Day":    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		"Valentine's Day":   time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
		"St. Patrick's Day": time.Date(2023, 3, 17, 0, 0, 0, 0, time.UTC),
		// ... (other events) ...
		"Chinese National Day": time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		"Event_11_25":          time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
		"Event_11_26":          time.Date(2023, 11, 26, 0, 0, 0, 0, time.UTC),
	},
	{
		"New Year's Day":    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		"Valentine's Day":   time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
		"St. Patrick's Day": time.Date(2023, 3, 17, 0, 0, 0, 0, time.UTC),
		// ... (other events) ...
		"Chinese National Day": time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		"Event_11_25":          time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
		"Event_11_26":          time.Date(2023, 11, 26, 0, 0, 0, 0, time.UTC),
	},
	{
		"New Year's Day":    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		"Valentine's Day":   time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
		"St. Patrick's Day": time.Date(2023, 3, 17, 0, 0, 0, 0, time.UTC),
		// ... (other events) ...
		"Chinese National Day": time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		"Event_11_25":          time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
		"Event_11_26":          time.Date(2023, 11, 26, 0, 0, 0, 0, time.UTC),
	},
}
