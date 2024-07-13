package internal

import (
	"parcelsApi/external"
	"parcelsApi/internal/common/parcels"
	"testing"
)

type dummyGrpcService struct{}
type dummyParcelService struct{}

func (dps *dummyParcelService) GetLatestParcelState(parcelReq *external.Request) (*external.State, error) {
	return &external.State{
		Status: "ok",
	}, nil
}
func (dps *dummyParcelService) GetParcel(parcelReq *external.Request) (*external.Tracking, error) {
	return &external.Tracking{
		UUID: "123",
	}, nil
}

func (dummyGrpcService *dummyGrpcService) GetAllParcels() (parcels.Parcels_GetParcelsClient, error) {
	return &dummyGetParcelsClient{count: 2}, nil
}

func (dummyGrpcService *dummyGrpcService) SaveParcel(parcelMessage *parcels.ParcelMessage) (*parcels.SaveParcelMessage, error) {
	return &parcels.SaveParcelMessage{
		IsSaved: true,
	}, nil
}

func TestRun(t *testing.T) {
	job := NewJobImpl(&dummyParcelService{}, &dummyGrpcService{})
	job.Run()
}
