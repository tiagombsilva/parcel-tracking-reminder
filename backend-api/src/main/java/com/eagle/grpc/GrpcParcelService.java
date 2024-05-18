package com.eagle.grpc;

import com.eagle.grpc.parcels.lib.ParcelMessage;
import com.eagle.grpc.parcels.lib.ParcelReq;
import com.eagle.grpc.parcels.lib.ParcelsGrpc;
import com.eagle.grpc.parcels.lib.SaveParcelMessage;
import com.eagle.rest.parcel.Parcel;
import com.eagle.rest.parcel.ParcelService;
import com.google.protobuf.Empty;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;

import java.sql.Timestamp;
import java.time.Instant;

@GrpcService
public class GrpcParcelService extends ParcelsGrpc.ParcelsImplBase {

    private final ParcelService parcelService;

    @Autowired
    public GrpcParcelService(final ParcelService parcelService) {
        this.parcelService = parcelService;
    }

    @Override
    public void getParcels(Empty request, StreamObserver<ParcelMessage> responseObserver) {
        var parcels = parcelService.getAllParcels();
        for (var parcel : parcels) {
            ParcelMessage response = ParcelMessage.newBuilder()
                    .setUuid(parcel.getUuid())
                    .setName(parcel.getName())
                    .setDestination(parcel.getDestination())
                    .setLastUpdate(parcel.getLastUpdate().toString())
                    .setOrigin(parcel.getOrigin())
                    .setTrackingCode(parcel.getTrackingCode())
                    .setStatus(parcel.getStatus())
                    .setIsDone(parcel.isDone())
                    .build();
            responseObserver.onNext(response);
        }
        if (parcels.isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
        } else {
            responseObserver.onCompleted();
        }
    }

    @Override
    public void getParcelByTrackingCode(ParcelReq request, StreamObserver<ParcelMessage> responseObserver) {
        var parcelOptional = parcelService.getParcelByTrackingCode(request.getTrackingCode());
        if (parcelOptional.isPresent()) {
            var parcel = parcelOptional.get();
            ParcelMessage response = ParcelMessage.newBuilder()
                    .setUuid(parcel.getUuid())
                    .setName(parcel.getName())
                    .setDestination(parcel.getDestination())
                    .setLastUpdate(parcel.getLastUpdate().toString())
                    .setOrigin(parcel.getOrigin())
                    .setTrackingCode(parcel.getTrackingCode())
                    .setStatus(parcel.getStatus())
                    .setIsDone(parcel.isDone())
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } else {
            responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
        }
    }

    @Override
    public void getParcelsInProgress(Empty request, StreamObserver<ParcelMessage> responseObserver) {
        var parcels = parcelService.getAllParcelsInProgress();
        for (var parcel : parcels) {
            ParcelMessage response = ParcelMessage.newBuilder()
                    .setUuid(parcel.getUuid())
                    .setName(parcel.getName())
                    .setDestination(parcel.getDestination())
                    .setLastUpdate(parcel.getLastUpdate().toString())
                    .setOrigin(parcel.getOrigin())
                    .setTrackingCode(parcel.getTrackingCode())
                    .setStatus(parcel.getStatus())
                    .setIsDone(parcel.isDone())
                    .build();
            responseObserver.onNext(response);
        }
        if (parcels.isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
        } else {
            responseObserver.onCompleted();
        }
    }

    @Override
    public void saveParcel(ParcelMessage request, StreamObserver<SaveParcelMessage> responseObserver) {
        var parcel = new Parcel();
        parcel.setUuid(request.getUuid());
        parcel.setName(request.getName());
        parcel.setOrigin(request.getOrigin());
        parcel.setDestination(request.getDestination());
        parcel.setStatus(request.getStatus());
        parcel.setLastUpdate(Timestamp.from(Instant.now()));
        parcel.setTrackingCode(request.getTrackingCode());
        parcel.setLastUpdate(Timestamp.valueOf(request.getLastUpdate()));
        parcel.setDone(request.getIsDone());
        var savedParcel = parcelService.saveOrUpdateParcel(parcel);
        var spm = SaveParcelMessage.newBuilder().setIsSaved(savedParcel.isPresent()).build();
        responseObserver.onNext(spm);
        responseObserver.onCompleted();
    }
}
