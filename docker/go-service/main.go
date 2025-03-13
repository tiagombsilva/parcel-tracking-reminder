package main

import (
	"flag"
	"go-service/external"
	"go-service/internal"
	"go-service/internal/common/parcels"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr   = flag.String("addr", "java-service:9090", "The grpc server host:port")
	cronSchedule = flag.String("cron", "0 12 * * *", "Cron schedule")
	apiUrl       = "https://parcelsapp.com/api/v3/shipments/tracking"
)

func main() {
	SetupFlags()
	parcelService := GetParcelService()
	grpcService := GetGrpcConnection(serverAddr)
	job := internal.NewJobImpl(parcelService, grpcService)
	log.Printf(*serverAddr)
	job.Run()
	//cron := cron.New()
	//cron.AddJob(*cronSchedule, job)
	//cron.Start()

	//interrupt := make(chan struct{})
	//<-interrupt
}

func GetParcelService() external.ParcelService {
	config, err := internal.ReadConfig("resources/config.json")
	if err != nil {
		log.Fatal("Failed to read config")
	}
	handler := external.NewParcelHandler(http.DefaultClient, apiUrl, config.ParcelsApiToken)
	return external.NewParcelService(handler)
}

func GetGrpcConnection(serverAddr *string) internal.ParcelGrpcService {
	conn, err := grpc.NewClient(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
