package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://127.0.0.1:5000"

func createEvent(name, date string) {
	event := map[string]string{
		"name": name,
		"date": date,
	}
	jsonData, err := json.Marshal(event)
	if err != nil {
		fmt.Println("Error encoding event:", err)
		return
	}

	resp, err := http.Post(baseURL+"/events", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating event:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Create Event Response:")
	fmt.Println(string(body))
}

func main() {
	createEvent("Go Conference", "2023-07-01")
}
