package models

type IncomingBuses []struct {
	VehicleID       string `json:"vehicleId"`
	LineName        string `json:"lineName"`
	DestinationName string `json:"destinationName"`
	ExpectedArrival string `json:"expectedArrival"`
}
