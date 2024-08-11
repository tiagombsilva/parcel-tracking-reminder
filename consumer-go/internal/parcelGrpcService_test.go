package internal

import (
	"context"
	"errors"
	"parcelsApi/internal/common/parcels"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type dummyExternalGrpcService struct {
}

type dummyGetParcelsClient struct {
	count uint32
}

func (dummyGetParcelsClient *dummyGetParcelsClient) Recv() (*parcels.ParcelMessage, error) {
	if dummyGetParcelsClient.count > 0 {
		dummyGetParcelsClient.count--
		return &parcels.ParcelMessage{
			TrackingCode: "123",
		}, nil
	}
	return nil, errors.New("Parcels empty")
}

func (dummyGetParcelsClient *dummyGetParcelsClient) CloseSend() error {
	return nil
}

func (dummyGetParcelsClient *dummyGetParcelsClient) Context() context.Context {
	return nil
}

func (dummyGetParcelsClient *dummyGetParcelsClient) Header() (metadata.MD, error) {
	return nil, nil
}

func (dummyGetParcelsClient *dummyGetParcelsClient) Trailer() metadata.MD {
	return nil
}

func (dummyGetParcelsClient *dummyGetParcelsClient) RecvMsg(T any) error {
	return nil
}

func (dummyGetParcelsClient *dummyGetParcelsClient) SendMsg(T any) error {
	return nil
}

func (dummyGrpcService *dummyExternalGrpcService) GetParcels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (parcels.Parcels_GetParcelsClient, error) {
	return &dummyGetParcelsClient{}, nil
}

func (dummyGrpcService *dummyExternalGrpcService) GetParcelByTrackingCode(ctx context.Context, in *parcels.ParcelReq, opts ...grpc.CallOption) (*parcels.ParcelMessage, error) {
	return &parcels.ParcelMessage{}, nil
}

func (dummyGrpcService *dummyExternalGrpcService) SaveParcel(ctx context.Context, in *parcels.ParcelMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func TestGetAllParcels(t *testing.T) {
	parcelServiceImpl := NewParcelGrpcServiceImpl(&dummyExternalGrpcService{})
	client, err := parcelServiceImpl.GetAllParcels()
	if err != nil {
		t.Error("Expecting client created")
	}
	message, err := client.Recv()
	if err != nil {
		t.Error("Expecting message")
	}

	if message.TrackingCode != "123" {
		t.Error("Expecting same Uuid")
	}
}

func TestSaveParcel(t *testing.T) {
	parcelServiceImpl := NewParcelGrpcServiceImpl(&dummyExternalGrpcService{})
	_, err := parcelServiceImpl.SaveParcel(&parcels.ParcelMessage{})
	if err != nil {
		t.Error("Expected success Message")
	}
}
