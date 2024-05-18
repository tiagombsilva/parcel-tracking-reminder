package external

import (
	"encoding/json"
	"log"
)

type ParcelService struct {
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

func NewParcelService(handler Handler) *ParcelService {
	return &ParcelService{handler}
}

func (ps *ParcelService) GetParcel(parcelReq *Request) *Tracking {
	log.Printf("Fetching parcel from external parcelsapp")
	jsonRes, err := ps.handler.PostParcel(parcelReq)
	if err != nil {
		log.Panic("failed to fetch Parcel")
		return nil
	}
	return getParcelFromJson(jsonRes)
}

func (ps *ParcelService) GetLatestParcelState(parcelReq *Request) *State {
	response := ps.GetParcel(parcelReq)
	return &response.Shipments[len(response.Shipments)-1].LastState
}

func getParcelFromJson(response []byte) *Tracking {
	var parcelJsons Tracking
	json.Unmarshal(response, &parcelJsons)
	return &parcelJsons
}
