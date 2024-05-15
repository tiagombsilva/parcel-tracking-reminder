package internal

import (
	"context"
	"parcelsApi/internal/common/parcels"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ParcelGrpcService interface {
	GetAllParcels() (*parcels.ParcelsResponse, error)
}

type ParcelGrpcServiceImpl struct {
	parcelService parcels.ParcelsClient
}

func NewParcelGrpcServiceImpl(parcelService parcels.ParcelsClient) *ParcelGrpcServiceImpl {
	return &ParcelGrpcServiceImpl{
		parcelService: parcelService,
	}
}

func (service *ParcelGrpcServiceImpl) GetAllParcels() (*parcels.ParcelsResponse, error) {
	return service.parcelService.GetParcels(context.Background(), &emptypb.Empty{})
}
