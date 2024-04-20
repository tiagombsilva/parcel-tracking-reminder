package service

import (
	"encoding/json"
	"log"
	"parcelsApi/handler"
)

type ParcelService struct {
	handler handler.Handler
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

type Parcel struct {
	UUID      string      `json:"uuid"`
	Shipments []Shipments `json:"shipments"`
}

func NewParcelService(handler handler.Handler) *ParcelService {
	return &ParcelService{handler}
}

func (ps *ParcelService) GetParcel(parcelReq *handler.Request) *Parcel {
	jsonRes, err := ps.handler.PostParcel(parcelReq)
	if err != nil {
		log.Panic("failed to fetch Parcel")
		return nil
	}
	return getParcelFromJson(jsonRes)
}

func getParcelFromJson(response []byte) *Parcel {
	var parcelJsons Parcel
	json.Unmarshal(response, &parcelJsons)
	return &parcelJsons
}
