package main

import (
	"flag"
	"parcelsApi/external"
	"parcelsApi/internal"
)

var (
	serverAddr = flag.String("addr", "localhost:9090", "The grpc server host:port")
)

func main() {
	parcelHandler := external.NewParcelHandler()
	parcelService := external.NewParcelService(parcelHandler)

	job := internal.NewJobImpl(*parcelService)
	grpcService := internal.GetGrpcConnection(serverAddr)
	job.Process(grpcService)
}
