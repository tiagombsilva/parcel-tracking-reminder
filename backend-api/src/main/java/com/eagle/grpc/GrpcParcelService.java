package com.eagle.grpc;

import com.eagle.grpc.parcels.lib.ParcelsGrpc;
import com.eagle.grpc.parcels.lib.ParcelsResponse;
import com.google.protobuf.Empty;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;

@GrpcService
public class GrpcParcelService extends ParcelsGrpc.ParcelsImplBase {
    @Override
    public void getParcels(Empty request, StreamObserver<ParcelsResponse> responseObserver) {
        ParcelsResponse response = ParcelsResponse.newBuilder().setName("Hello").build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}
