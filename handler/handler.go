package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	websitesStatus = map[string]string{}
)

var WebsitesStatus map[string]int

func AddWebsites(w http.ResponseWriter, r *http.Request) {
	var website []string
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &website)
	if err != nil {
		log.Printf("Error unmarshalling response")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
	resp, _ := json.Marshal("Websites added")
	w.Write(resp)
	log.Print("Websites stored")
	go Status(website)
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("website")
	if q != "" {
		fmt.Println("Hello=", q, websitesStatus)
		status := make(map[string]string)
		status[q] = websitesStatus[q]
		resp, err := json.Marshal(status)
		if err != nil {
			fmt.Errorf("Error marshalling status of %v", q)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(resp)
		log.Print("Status returned...")
	} else {
		resp, err := json.Marshal(websitesStatus)
		if err != nil {
			fmt.Errorf("Error marshalling status of websites")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(resp)
		log.Print("Status returned...")
	}
}

func Status(website []string) {
	for {
		for _, val := range website {
			_, err := http.Get("http://" + val)
			if err != nil {
				websitesStatus[val] = "DOWN"
			} else {
				websitesStatus[val] = "UP"
			}
		}
		time.Sleep(time.Minute)
	}
}
