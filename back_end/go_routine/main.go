package main

import (
	"time"

	"github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/outlet"
)

func main() {
	outlet.INIT()
	time.Sleep(6000 * time.Second) // TODO: Keep main function running
}
