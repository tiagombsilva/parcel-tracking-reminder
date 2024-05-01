package com.eagle.grpc;

import com.eagle.grpc.parcels.lib.ParcelReq;
import com.eagle.grpc.parcels.lib.ParcelsGrpc;
import com.eagle.grpc.parcels.lib.ParcelsResponse;
import com.eagle.rest.parcel.ParcelService;
import com.google.protobuf.Empty;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;

@GrpcService
public class GrpcParcelService extends ParcelsGrpc.ParcelsImplBase {

    private final ParcelService parcelService;

    @Autowired
    public GrpcParcelService(final ParcelService parcelService) {
        this.parcelService = parcelService;
    }

    @Override
    public void getParcels(Empty request, StreamObserver<ParcelsResponse> responseObserver) {
        var parcels = parcelService.getAllParcels();
        for (var parcel : parcels) {
            ParcelsResponse response = ParcelsResponse.newBuilder()
                    .setUuid(parcel.getUuid())
                    .setName(parcel.getName())
                    .setDestination(parcel.getDestination())
                    .setLastUpdate(parcel.getLastUpdate().toString())
                    .setOrigin(parcel.getOrigin())
                    .setTrackingCode(parcel.getTrackingCode())
                    .setStatus(parcel.getStatus())
                    .build();
            responseObserver.onNext(response);
        }
        responseObserver.onCompleted();
    }

    @Override
    public void getParcelByTrackingCode(ParcelReq request, StreamObserver<ParcelsResponse> responseObserver) {
        var parcelOptional = parcelService.getParcelByTrackingCode(request.getTrackingCode());
        if (parcelOptional.isPresent()) {
            var parcel = parcelOptional.get();
            ParcelsResponse response = ParcelsResponse.newBuilder()
                    .setUuid(parcel.getUuid())
                    .setName(parcel.getName())
                    .setDestination(parcel.getDestination())
                    .setLastUpdate(parcel.getLastUpdate().toString())
                    .setOrigin(parcel.getOrigin())
                    .setTrackingCode(parcel.getTrackingCode())
                    .setStatus(parcel.getStatus())
                    .build();
            responseObserver.onNext(response);
        }
        responseObserver.onCompleted();
    }
}
