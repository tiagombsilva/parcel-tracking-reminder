package internal

import (
	"log"
	"parcelsApi/external"
)

type JobImpl struct {
	parcelService *external.ParcelService
	grpcService   ParcelGrpcService
}

func NewJobImpl(parcelService *external.ParcelService, grpcService ParcelGrpcService) *JobImpl {
	return &JobImpl{
		parcelService: parcelService,
		grpcService:   grpcService,
	}
}

func (job *JobImpl) Run() {
	grpcResponse, err := job.grpcService.GetAllParcels()
	if err != nil {
		log.Printf("Failed to to get data from GRPC")
	}
	log.Printf("Received from GRPC: %s", grpcResponse)

	parcelReq := &external.Request{
		TrackingId:  "1Z67CPT40409456025",
		DestCountry: "Portugal",
		//Zipcode:     "9760-180",
	}
	response := job.parcelService.GetParcel(parcelReq)
	log.Printf("Received package: %s", response.Shipments[len(response.Shipments)-1].LastState)
}
