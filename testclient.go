package main

import (
	"log"
	"net/http"

	centralhub "github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/central_hub"
)

func main() {
	centralHub := centralhub.GetHubInstanceDefault()
	http.HandleFunc("/centralhub", centralHub.HandleWebSocket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
