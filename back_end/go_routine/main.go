package main

import (
	"time"

	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/outlet"
)

func main() {
	// Test the initialization of the outlet
	outlet.INIT()
	time.Sleep(6000 * time.Second)
}
