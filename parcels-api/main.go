package main

import (
	"fmt"
	"parcelsApi/handler"
	"parcelsApi/service"
)

func main() {
	parcelHandler := handler.NewParcelHandler()
	parcelService := service.NewParcelService(parcelHandler)
	parcelReq := &handler.Request{
		TrackingId:  "31857080137424",
		DestCountry: "Portugal",
		Zipcode:     "9760-180",
	}
	response := parcelService.GetParcel(parcelReq)
	fmt.Println(response.Shipments[len(response.Shipments)-1].LastState)
}
