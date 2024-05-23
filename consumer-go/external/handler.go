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

type Handler interface {
	PostParcel(data *Request) ([]byte, error)
}

type ParcelHandler struct {
	client *http.Client
	apiUrl string
}

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

func NewParcelHandler(client *http.Client, apiUrl string) *ParcelHandler {
	if client == nil {
		client = &http.Client{}
	}
	return &ParcelHandler{
		client: client,
		apiUrl: apiUrl,
	}
}

func (handler *ParcelHandler) PostParcel(reqParcels *Request) ([]byte, error) {
	data := getJsonReq(reqParcels)
	base, _ := url.Parse(handler.apiUrl)
	jsonReq, _ := json.Marshal(data)
	log.Println("Sending post with body: ", data)
	response, err := handler.client.Post(base.String(), "application/json", bytes.NewBuffer(jsonReq))
	if err == nil {
		responseJson, _ := io.ReadAll(response.Body)
		log.Println("Response data: ", string(responseJson))
		return responseJson, nil
	}
	log.Println("Failed fetching data")
	return nil, err
}

func getJsonReq(parcelReq *Request) *ParcelRequest {
	return &ParcelRequest{
		Language:  "en",
		ApiKey:    os.Getenv("parcelsApiToken"),
		Shipments: []Request{*parcelReq},
	}
}
