package internal

import (
	"go-service/external"
	"go-service/internal/common/parcels"
	"log"
	"strings"
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
	grpcResponse, err := job.grpcService.GetParcelsInProgress()
	if err != nil {
		log.Printf("Failed to to get data from GRPC")
		return
	}
	log.Printf("Starting to update parcels")
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
	log.Printf("Updating Parcel '%s'", res)
	latestUpdate, err := job.getLatestUpdate(res)
	if err != nil {
		log.Printf("Failed to get latest update %s", err)
	} else {
		log.Printf("Latest state found: %s", latestUpdate)
		job.updateToLatestState(res, latestUpdate)
	}
	defer job.wg.Done()
}

func (job *JobImpl) getLatestUpdate(res *parcels.ParcelMessage) (*external.UpdateResponse, error) {
	parcelReq := &external.Request{
		TrackingId:  res.GetTrackingCode(),
		DestCountry: res.GetDestination(),
		Zipcode:     res.GetZipCode(),
	}
	return job.parcelService.GetParcelUpdates(parcelReq)
}

func (job *JobImpl) updateToLatestState(res *parcels.ParcelMessage, lastUpdate *external.UpdateResponse) {
	parcelMessage := &parcels.ParcelMessage{
		Uuid:         res.Uuid,
		TrackingCode: res.TrackingCode,
		Name:         res.Name,
		Origin:       res.Origin,
		Destination:  res.Destination,
		LastUpdate:   &lastUpdate.State.Date,
		Status:       &lastUpdate.State.Status,
		ZipCode:      res.ZipCode,
		IsDone:       lastUpdate.Done,
	}
	if job.shouldUpdate(res.LastUpdate, &lastUpdate.State.Date) {
		log.Printf("Saving new status...")
		_, err := job.grpcService.SaveOrUpdateParcel(parcelMessage)
		if err != nil {
			log.Printf("Failed to save new Status %s", err)
		} else {
			log.Printf("Successfully saved new Status for parcel: %s", res)
		}
	}
}

func (job *JobImpl) shouldUpdate(date *string, date2 *string) bool {
	if date == nil || len(*date) == 0 {
		return true
	}
	dateParsed := (*date)[:13]
	dateParsed2 := (*date2)[:13]
	return strings.Compare(dateParsed, dateParsed2) != 1
}
