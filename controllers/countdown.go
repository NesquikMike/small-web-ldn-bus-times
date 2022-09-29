package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/nesquikmike/small-web-ldn-bus-times/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"html/template"
)

type Controller struct {
	tpl *template.Template
}

func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

func (c Controller) Countdown(w http.ResponseWriter, req *http.Request) {
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

	incomingBusesToStop := models.IncomingBuses{}
	err = json.Unmarshal(body, &incomingBusesToStop)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(incomingBusesToStop)

	c.tpl.ExecuteTemplate(w, "countdown.gohtml", incomingBusesToStop)
}

func (c Controller) Index(w http.ResponseWriter, req *http.Request) {
	c.tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
