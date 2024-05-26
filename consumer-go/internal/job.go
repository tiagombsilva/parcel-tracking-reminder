package internal

import (
	"log"
	"parcelsApi/external"
	"parcelsApi/internal/common/parcels"
	"sync"
)

type JobImpl struct {
	parcelService *external.ParcelService
	grpcService   ParcelGrpcService
	wg            sync.WaitGroup
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
		return
	}
	log.Printf("Updating all Parcels from GRPC...")
	for {
		resp, err := grpcResponse.Recv()
		if err != nil {
			break
		}
		job.wg.Add(1)
		go job.updatePackageToLatestState(resp)
	}
	job.wg.Wait()
	log.Print("All Parcels updated!")
}

func (job *JobImpl) updatePackageToLatestState(res *parcels.ParcelMessage) {
	defer job.wg.Done()
	log.Printf("Updating Parcel '%s'", res)

	// parcelReq := &external.Request{
	// 	TrackingId:  "LV997747362CN",
	// 	DestCountry: "Portugal",
	// 	Zipcode:     "9760-180",
	// }
	// latestState, err := parcelService.GetLatestParcelState(parcelReq)
	// if err != nil {
	// 	log.Print(err.Error())
	// } else {
	// 	log.Printf("package: %s", latestState)
	// }
}
