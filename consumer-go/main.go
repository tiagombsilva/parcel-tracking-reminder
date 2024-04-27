package main

import (
	"fmt"
	"parcelsApi/external"
	"parcelsApi/service"
)

func main() {
	parcelHandler := external.NewParcelHandler()
	parcelService := service.NewParcelService(parcelHandler)
	parcelReq := &external.Request{
		TrackingId:  "31857080137424",
		DestCountry: "Portugal",
		Zipcode:     "9760-180",
	}
	response := parcelService.GetParcel(parcelReq)
	fmt.Println(response.Shipments[len(response.Shipments)-1].LastState)
}
