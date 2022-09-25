package main

import (
	"html/template"
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

var tpl *template.Template

type incomingBuses []struct {
	VehicleID       string `json:"vehicleId"`
	LineName        string `json:"lineName"`
	DestinationName string `json:"destinationName"`
	ExpectedArrival string `json:"expectedArrival"`
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/countdown", countdown)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func countdown(w http.ResponseWriter, req *http.Request) {
	stopCode := req.FormValue("stop-code")

	tflStopPointApiUrl := fmt.Sprintf("https://api.tfl.gov.uk/StopPoint/%s/arrivals", stopCode)

	stopPointClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	stopPointReq, err := http.NewRequest(http.MethodGet, tflStopPointApiUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	stopPointRes, err := stopPointClient.Do(stopPointReq)
	if err != nil {
		log.Fatal(err)
	}

	if stopPointRes.Body != nil {
		defer stopPointRes.Body.Close()
	}

	body, err := ioutil.ReadAll(stopPointRes.Body)
	if err != nil {
		log.Fatal(err)
	}

	incomingBusesToStop := incomingBuses{}
	err = json.Unmarshal(body, &incomingBusesToStop)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(incomingBusesToStop)

	tpl.ExecuteTemplate(w, "countdown.gohtml", incomingBusesToStop)
}
