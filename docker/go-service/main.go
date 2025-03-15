package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"go-service/external"
	"go-service/internal"
	"go-service/internal/common/parcels"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	cert, err := os.ReadFile("server-cert.pem")
	if err != nil {
		log.Fatalf("failed to read server certificate: %v", err)
	}

	// Create a certificate pool and add the server's certificate
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Fatalf("failed to add server certificate to pool")
	}

	// Create TLS configuration
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	// Create gRPC credentials from the TLS configuration
	creds := credentials.NewTLS(tlsConfig)

	conn, err := grpc.NewClient(*serverAddr, grpc.WithTransportCredentials(creds))
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
