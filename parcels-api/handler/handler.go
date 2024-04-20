package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type State struct {
	Date    string `json: "date"`
	Carrier string `json: "carrier"`
	Status  string `json: "status"`
}

type Shipments struct {
	OriginCode      string `json: "originCode"`
	DestinationCode string `json: "destinationCode"`
	Status          string `json: "status"`
	TrackingId      string `json: "trackingId"`
	LastState       State  `json: "lastState"`
}

type Parcel struct {
	UUID      string      `json: "uuid"`
	Shipments []Shipments `json: "shipments"`
}

func GetParcel(uuid string) *Parcel {
	base, err := url.Parse("https://parcelsapp.com/api/v3/shipments/tracking")
	if err != nil {
		return nil
	}
	params := url.Values{}
	params.Add("uuid", uuid)
	params.Add("apiKey", os.Getenv("parcelsApiToken"))
	base.RawQuery = params.Encode()
	response, err := http.Get(base.String())
	if err != nil {
		log.Panic(err.Error())
		return nil
	}
	return getParcelFromJson(response)
}

func getParcelFromJson(response *http.Response) *Parcel {
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var parcelJsons Parcel
	json.Unmarshal(responseData, &parcelJsons)
	return &parcelJsons
}
