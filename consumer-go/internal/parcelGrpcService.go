package internal

import (
	"context"
	"parcelsApi/internal/common/parcels"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ParcelGrpcService interface {
	GetAllParcels() (parcels.Parcels_GetParcelsClient, error)
}

type ParcelGrpcServiceImpl struct {
	parcelClient parcels.ParcelsClient
}

func NewParcelGrpcServiceImpl(parcelClient parcels.ParcelsClient) *ParcelGrpcServiceImpl {
	return &ParcelGrpcServiceImpl{
		parcelClient: parcelClient,
	}
}

func (service *ParcelGrpcServiceImpl) GetAllParcels() (parcels.Parcels_GetParcelsClient, error) {
	return service.parcelClient.GetParcels(context.Background(), &emptypb.Empty{})
}
