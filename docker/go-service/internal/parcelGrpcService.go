package internal

import (
	"context"
	"go-service/internal/common/parcels"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ParcelGrpcService interface {
	GetParcelsInProgress() (parcels.Parcels_GetParcelsInProgressClient, error)
	SaveOrUpdateParcel(parcelMessage *parcels.ParcelMessage) (*emptypb.Empty, error)
}

type ParcelGrpcServiceImpl struct {
	parcelClient parcels.ParcelsClient
}

func NewParcelGrpcServiceImpl(parcelClient parcels.ParcelsClient) *ParcelGrpcServiceImpl {
	return &ParcelGrpcServiceImpl{
		parcelClient: parcelClient,
	}
}

func (service *ParcelGrpcServiceImpl) GetParcelsInProgress() (parcels.Parcels_GetParcelsInProgressClient, error) {
	return service.parcelClient.GetParcelsInProgress(context.Background(), &emptypb.Empty{})
}

func (service *ParcelGrpcServiceImpl) SaveOrUpdateParcel(parcelMessage *parcels.ParcelMessage) (*emptypb.Empty, error) {
	return service.parcelClient.SaveOrUpdateParcel(context.Background(), parcelMessage)
}
