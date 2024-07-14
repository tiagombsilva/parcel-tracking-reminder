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

import java.time.ZonedDateTime;

import static com.eagle.utils.Helper.getOrDefault;

@GrpcService
public class GrpcParcelService extends ParcelsGrpc.ParcelsImplBase {

    private final ParcelService parcelService;

    @Autowired
    public GrpcParcelService(final ParcelService parcelService) {
        this.parcelService = parcelService;
    }

    static ParcelMessage getParcelMessage(final Parcel parcel) {
        return ParcelMessage.newBuilder()
                .setUuid(parcel.getUuid())
                .setName(parcel.getName())
                .setDestination(getOrDefault(parcel.getDestination()))
                .setLastUpdate(getOrDefault(parcel.getLastUpdate()))
                .setOrigin(getOrDefault(parcel.getOrigin()))
                .setTrackingCode(parcel.getTrackingCode())
                .setStatus(getOrDefault(parcel.getStatus()))
                .setZipCode(getOrDefault(parcel.getZipCode()))
                .setIsDone(parcel.isDone())
                .build();
    }

    @Override
    public void getParcels(Empty request, StreamObserver<ParcelMessage> responseObserver) {
        var parcels = parcelService.getAllParcels();
        for (var parcel : parcels) {
            ParcelMessage response = getParcelMessage(parcel);
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
            ParcelMessage response = getParcelMessage(parcel);
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } else {
            responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
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
        parcel.setTrackingCode(request.getTrackingCode());
        parcel.setLastUpdate(ZonedDateTime.parse(request.getLastUpdate()));
        parcel.setZipCode(request.getZipCode());
        parcel.setDone(request.getIsDone());
        var savedParcel = parcelService.saveOrUpdateParcel(parcel);
        var spm = SaveParcelMessage.newBuilder().setIsSaved(savedParcel != null).build();
        responseObserver.onNext(spm);
        responseObserver.onCompleted();
    }
}
