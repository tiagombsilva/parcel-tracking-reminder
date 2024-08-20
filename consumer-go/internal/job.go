package internal

import (
	"log"
	"parcelsApi/external"
	"parcelsApi/internal/common/parcels"
	"sync"
)

type JobImpl struct {
	parcelService external.ParcelService
	grpcService   ParcelGrpcService
	wg            sync.WaitGroup
}

func NewJobImpl(parcelService external.ParcelService, grpcService ParcelGrpcService) *JobImpl {
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
		if resp.GetIsDone() {
			continue
		}
		job.wg.Add(1)
		go job.updatePackageToLatestState(resp)
	}
	job.wg.Wait()
	log.Print("All Parcels updated!")
}

func (job *JobImpl) updatePackageToLatestState(res *parcels.ParcelMessage) {
	log.Printf("Updating Parcel '%s'", res)
	latestState := job.getLatestState(res)
	job.updateToLatestState(res, latestState)
	defer job.wg.Done()
}

func (job *JobImpl) getLatestState(res *parcels.ParcelMessage) *external.State {
	parcelReq := &external.Request{
		TrackingId:  res.GetTrackingCode(),
		DestCountry: res.GetDestination(),
		Zipcode:     res.GetZipCode(),
	}
	latestState, err := job.parcelService.GetLatestParcelState(parcelReq)
	if err != nil {
		log.Print(err.Error())
		return nil
	} else {
		log.Printf("Latest state found: %s", latestState)
		return latestState
	}
}

func (job *JobImpl) updateToLatestState(res *parcels.ParcelMessage, latestState *external.State) {
	parcelMessage := &parcels.ParcelMessage{
		Uuid:         res.Uuid,
		TrackingCode: res.TrackingCode,
		Name:         res.Name,
		Origin:       res.Origin,
		Destination:  res.Destination,
		LastUpdate:   &latestState.Date,
		Status:       &latestState.Status,
		ZipCode:      res.ZipCode,
		IsDone:       res.IsDone,
	}
	log.Printf("Saving new status...")
	if res.LastUpdate != &latestState.Date || res.Uuid == nil || *res.Uuid == "" {
		_, err := job.grpcService.SaveOrUpdateParcel(parcelMessage)
		if err != nil {
			log.Printf("Failed to save new Status %s", err)
		} else {
			log.Printf("Successfully saved new Status for parcel: %s", res)
		}
	}
}
