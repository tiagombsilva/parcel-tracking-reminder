package external

import (
	"encoding/json"
	"errors"
	"log"
)

type ParcelService interface {
	GetParcel(parcelReq *Request) (*Tracking, error)
	GetLatestParcelState(parcelReq *Request) (*State, error)
}

type ParcelServiceImpl struct {
	handler Handler
}

type State struct {
	Date    string `json:"date"`
	Carrier string `json:"carrier"`
	Status  string `json:"status"`
}

type Shipments struct {
	OriginCode      string `json:"originCode"`
	DestinationCode string `json:"destinationCode"`
	Status          string `json:"status"`
	TrackingId      string `json:"trackingId"`
	LastState       State  `json:"lastState"`
}

type Tracking struct {
	UUID      string      `json:"uuid"`
	Shipments []Shipments `json:"shipments"`
	Done      bool        `json:"done"`
}

type ResponseError struct {
	Error string `json:"error"`
}

func NewParcelService(handler Handler) *ParcelServiceImpl {
	return &ParcelServiceImpl{handler}
}

func (ps *ParcelServiceImpl) GetParcel(parcelReq *Request) (*Tracking, error) {
	log.Printf("Fetching parcel from external parcelsapp")
	jsonRes, err := ps.handler.PostParcel(parcelReq)
	if err != nil {
		return nil, err
	}
	jsonUnmarshal, err := getParcelFromJson(jsonRes)
	if err != nil {
		return nil, err
	}
	return jsonUnmarshal, nil
}

func (ps *ParcelServiceImpl) GetLatestParcelState(parcelReq *Request) (*State, error) {
	response, err := ps.GetParcel(parcelReq)
	if err != nil {
		return nil, err
	}
	return &response.Shipments[len(response.Shipments)-1].LastState, nil
}

func getError(response []byte) error {
	var responseError ResponseError
	json.Unmarshal(response, &responseError)
	if responseError.Error != "" {
		return errors.New("Error found in API: " + responseError.Error)
	}
	return nil
}

func getParcelFromJson(response []byte) (*Tracking, error) {
	err := getError(response)
	if err != nil {
		return nil, err
	}
	var parcelJsons Tracking
	json.Unmarshal(response, &parcelJsons)
	return &parcelJsons, nil
}
