package main

import (
	"fmt"
	"time"

	centralhub "github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/central_hub"
	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/outlet"
)

func main() {
	// Create a central hub instance
	centralHub := centralhub.GetHubInstance("Central_Hub1", "Location1")

	// Create an outlet
	numberOutlets := 4
	outlets := make([]*outlet.Outlet, numberOutlets)
	for i := 'A'; i <= 'D'; i++ {
		location := fmt.Sprintf("Location %d", i)
		outletID := fmt.Sprintf("Outlet %d", i)
		outlets[i] = outlet.NewOutlet(outletID, location, 5, holidayMaps[i])
		fmt.Println("Created outlet:", outlets[i].GetOutletID())
	}

	// Sleep to allow for background processes to run
	time.Sleep(10 * time.Second) // TODO: Keep main function running

	// Assuming there might be additional functionality to print or check the status
	fmt.Printf("Central Hub: %+v\n", centralHub)
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
