package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://127.0.0.1:5000"

func getEvents() {
	resp, err := http.Get(baseURL + "/events")
	if err != nil {
		fmt.Println("Error fetching events:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Events List:")
	fmt.Println(string(body))
}

func getEvent(eventID int) {
	resp, err := http.Get(fmt.Sprintf("%s/events/%d", baseURL, eventID))
	if err != nil {
		fmt.Println("Error fetching event:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		fmt.Println("Event not found")
		return
	}

	var event map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&event)
	if err != nil {
		fmt.Println("Error decoding event:", err)
		return
	}

	fmt.Printf("Event Details: %+v\n", event)
}

func main() {
	getEvents()

	getEvent(1) // 假设您想获取ID为1的事件
}
