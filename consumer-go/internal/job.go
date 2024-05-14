package internal

import (
	"log"
	"parcelsApi/external"
	"parcelsApi/internal/common/parcels"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Job interface {
	Process()
}

type JobImpl struct {
	parcelService external.ParcelService
}

func NewJobImpl(parcelService external.ParcelService) *JobImpl {
	return &JobImpl{
		parcelService: parcelService,
	}
}

func (job *JobImpl) Process(grpcService *ParcelGrpcServiceImpl) {
	grpcResponse, err := grpcService.GetParcels()
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

func GetGrpcConnection(serverAddr *string) *ParcelGrpcServiceImpl {
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to setup connection to Grpc server")
	}
	grpcConn := parcels.NewParcelsClient(conn)
	return NewParcelGrpcServiceImpl(grpcConn)
}
