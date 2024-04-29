package external

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type parcelHandler struct{}

type ParcelRequest struct {
	Shipments []Request `json:"shipments"`
	Language  string    `json:"language"`
	ApiKey    string    `json:"apiKey"`
}

type Request struct {
	TrackingId  string `json:"trackingId"`
	DestCountry string `json:"destinationCountry"`
	Zipcode     string `json:"zipcode"`
}

func NewParcelHandler() *parcelHandler {
	return &parcelHandler{}
}

func (handler *parcelHandler) PostParcel(reqParcels *Request) ([]byte, error) {
	data := getJsonReq(reqParcels)
	base, _ := url.Parse("https://parcelsapp.com/api/v3/shipments/tracking")
	jsonReq, _ := json.Marshal(data)
	log.Println("Sending post with body: ", data)
	response, err := http.Post(base.String(), "application/json", bytes.NewBuffer(jsonReq))
	if err == nil {
		responseJson, _ := io.ReadAll(response.Body)
		log.Println("response data: ", string(responseJson))
		return responseJson, nil
	}
	log.Println("failed fetching data")
	return nil, err
}

func getJsonReq(parcelReq *Request) *ParcelRequest {
	return &ParcelRequest{
		Language:  "en",
		ApiKey:    os.Getenv("parcelsApiToken"),
		Shipments: []Request{*parcelReq},
	}
}