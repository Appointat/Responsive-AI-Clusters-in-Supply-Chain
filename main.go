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
	outlet1 := outlet.NewOutlet("Outlet1", "Location1", 3)
	outlet1.AddEvent("Test2")

	// Output the created outlet information
	fmt.Println("Created outlet:", outlet1.GetOutletID())

	// Sleep to allow for background processes to run
	time.Sleep(10 * time.Second)

	// Assuming there might be additional functionality to print or check the status
	// This could include printing the current date or events
	fmt.Println("Current Date:", outlet.GetCurrentDate())
	fmt.Printf("Central Hub: %+v\n", centralHub)
}
