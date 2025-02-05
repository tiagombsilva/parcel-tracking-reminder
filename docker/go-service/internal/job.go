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
	log.Printf("Starting to update parcels. Current GRPC response: %v", grpcResponse)
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
	latestState := job.getLatestState(res)
	if latestState != nil {
		job.updateToLatestState(res, latestState)
	}
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
	}
	isDelivered := strings.Contains(latestState.Status, "Delivered")
	parcelMessage.IsDone = &isDelivered
	if job.checkDate(res.LastUpdate, &latestState.Date) {
		log.Printf("Saving new status...")
		_, err := job.grpcService.SaveOrUpdateParcel(parcelMessage)
		if err != nil {
			log.Printf("Failed to save new Status %s", err)
		} else {
			log.Printf("Successfully saved new Status for parcel: %s", res)
		}
	}
}

func (job *JobImpl) checkDate(date *string, date2 *string) bool {
	dateParsed := (*date)[:13]
	dateParsed2 := (*date2)[:13]
	return strings.Compare(dateParsed, dateParsed2) == 1
}
