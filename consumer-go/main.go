package main

import (
	"flag"
	"log"
	"net/http"
	"parcelsApi/external"
	"parcelsApi/internal"
	"parcelsApi/internal/common/parcels"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr   = flag.String("addr", "localhost:9090", "The grpc server host:port")
	cronSchedule = flag.String("cron", "0 12 * * *", "Cron schedule")
	apiUrl       = "https://parcelsapp.com/api/v3/shipments/tracking"
)

func main() {
	SetupFlags()
	parcelService := GetParcelService()
	grpcService := GetGrpcConnection(serverAddr)
	job := internal.NewJobImpl(parcelService, grpcService)

	job.Run()
	//cron := cron.New()
	//cron.AddJob(*cronSchedule, job)
	//cron.Start()

	//interrupt := make(chan struct{})
	//<-interrupt
}

func GetParcelService() *external.ParcelService {
	handler := external.NewParcelHandler(http.DefaultClient, apiUrl)
	return external.NewParcelService(handler)
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
