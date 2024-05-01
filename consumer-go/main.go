package main

import (
	"flag"
	"log"
	"parcelsApi/external"
	"parcelsApi/internal"
	"parcelsApi/internal/common/parcels"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:9090", "The grpc server host:port")
)

func main() {
	parcelHandler := external.NewParcelHandler()
	parcelService := external.NewParcelService(parcelHandler)
	parcelReq := &external.Request{
		TrackingId:  "1Z67CPT40409456025",
		DestCountry: "Portugal",
		//Zipcode:     "9760-180",
	}
	response := parcelService.GetParcel(parcelReq)
	log.Printf("Received package: %s", response.Shipments[len(response.Shipments)-1].LastState)
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to setup connection to Grpc server")
	}
	grpcConn := parcels.NewParcelsClient(conn)
	service := internal.NewParcelGrpcServiceImpl(grpcConn)

	grpcResponse, err := service.GetParcels()
	if err != nil {
		log.Fatal("Failed to to get data from GRPC")
	}
	log.Printf("Received from GRPC: %s", grpcResponse)
}
