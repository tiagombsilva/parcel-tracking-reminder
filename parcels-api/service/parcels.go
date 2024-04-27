package service

import (
	"encoding/json"
	"log"
	"parcelsApi/external"
)

type ParcelService struct {
	handler external.Handler
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
}

func NewParcelService(handler external.Handler) *ParcelService {
	return &ParcelService{handler}
}

func (ps *ParcelService) GetParcel(parcelReq *external.Request) *Tracking {
	jsonRes, err := ps.handler.PostParcel(parcelReq)
	if err != nil {
		log.Panic("failed to fetch Parcel")
		return nil
	}
	return getParcelFromJson(jsonRes)
}

func getParcelFromJson(response []byte) *Tracking {
	var parcelJsons Tracking
	json.Unmarshal(response, &parcelJsons)
	return &parcelJsons
}
