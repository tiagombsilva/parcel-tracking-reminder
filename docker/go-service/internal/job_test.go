package internal

import (
	"go-service/external"
	"go-service/internal/common/parcels"
	"testing"

	"google.golang.org/protobuf/types/known/emptypb"
)

type dummyGrpcService struct{}
type dummyParcelService struct{}

func (dps *dummyParcelService) GetParcelUpdates(parcelReq *external.Request) (*external.UpdateResponse, error) {
	done := true
	return &external.UpdateResponse{
		State: &external.State{
			Status: "ok",
		},
		Done: &done,
	}, nil
}
func (dps *dummyParcelService) GetParcel(parcelReq *external.Request) (*external.Tracking, error) {
	return &external.Tracking{
		UUID: "123",
	}, nil
}

func (dummyGrpcService *dummyGrpcService) GetParcelsInProgress() (parcels.Parcels_GetParcelsInProgressClient, error) {
	return &dummyGetParcelsClient{count: 2}, nil
}

func (dummyGrpcService *dummyGrpcService) SaveOrUpdateParcel(parcelMessage *parcels.ParcelMessage) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func TestRun(t *testing.T) {
	job := NewJobImpl(&dummyParcelService{}, &dummyGrpcService{})
	job.Run()
}
