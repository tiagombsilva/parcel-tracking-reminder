package internal

import (
	"parcelsApi/external"
	"parcelsApi/internal/common/parcels"
	"testing"

	"google.golang.org/protobuf/types/known/emptypb"
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

func (dummyGrpcService *dummyGrpcService) SaveOrUpdateParcel(parcelMessage *parcels.ParcelMessage) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func TestRun(t *testing.T) {
	job := NewJobImpl(&dummyParcelService{}, &dummyGrpcService{})
	job.Run()
}
