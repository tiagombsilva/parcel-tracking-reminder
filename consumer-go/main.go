package main

import (
	"flag"
	"log"
	"parcelsApi/external"
	"parcelsApi/internal"
	"parcelsApi/internal/common/parcels"

	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr   = flag.String("addr", "localhost:9090", "The grpc server host:port")
	cronSchedule = flag.String("cron", "0 12 * * *", "Cron schedule")
)

func main() {
	SetupFlags()
	parcelHandler := external.NewParcelHandler()
	parcelService := external.NewParcelService(parcelHandler)
	grpcService := GetGrpcConnection(serverAddr)
	job := internal.NewJobImpl(parcelService, grpcService)

	parcelReq := &external.Request{
		TrackingId:  "LV997747362CN",
		DestCountry: "Portugal",
		//Zipcode:     "9760-180",
	}
	latestState := parcelService.GetLatestParcelState(parcelReq)
	log.Printf("package: %s", latestState)

	cron := cron.New()
	cron.AddJob(*cronSchedule, job)
	cron.Start()

	// Use a channel to keep the main goroutine alive
	// until an interrupt signal is received
	interrupt := make(chan struct{})
	<-interrupt
}

func GetGrpcConnection(serverAddr *string) internal.ParcelGrpcService {
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to setup connection to Grpc server")
	}
	grpcConn := parcels.NewParcelsClient(conn)
	return internal.NewParcelGrpcServiceImpl(grpcConn)
}

func SetupFlags() {
	flag.Parse()
	log.Printf("Server address: %s", *serverAddr)
	log.Printf("Cron Schedule: %s", *cronSchedule)
}
