package com.eagle.grpc;

import com.eagle.grpc.accounts.lib.AccountMessage;
import com.eagle.grpc.accounts.lib.AccountParcelMessage;
import com.eagle.grpc.accounts.lib.AccountReq;
import com.eagle.grpc.accounts.lib.AccountSaveParcel;
import com.eagle.grpc.accounts.lib.AccountsGrpc;
import com.eagle.grpc.accounts.lib.SaveAccountResponse;
import com.eagle.rest.account.Account;
import com.eagle.rest.account.AccountService;
import com.eagle.rest.parcel.Parcel;
import com.google.protobuf.Empty;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;

import java.sql.Timestamp;

@GrpcService
public class GrpcAccountService extends AccountsGrpc.AccountsImplBase {

    private final AccountService accountService;

    @Autowired
    public GrpcAccountService(final AccountService accountService) {
        this.accountService = accountService;
    }

    @Override
    public void getAccounts(Empty request, StreamObserver<AccountMessage> responseObserver) {
        var accounts = accountService.getAllAccounts();
        for (var account : accounts) {
            AccountMessage accountResponse = AccountMessage.newBuilder()
                    .setDiscordId(account.getDiscordId())
                    .setName(account.getName())
                    .build();
            responseObserver.onNext(accountResponse);
        }
        if (accounts.isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
        } else {
            responseObserver.onCompleted();
        }
    }

    @Override
    public void getAccountByDiscordId(AccountReq request, StreamObserver<AccountMessage> responseObserver) {
        var accountOptional = accountService.getAccountById(request.getDiscordId());
        if (accountOptional.isPresent()) {
            var account = accountOptional.get();
            AccountMessage accountResponse = AccountMessage.newBuilder()
                    .setDiscordId(account.getDiscordId())
                    .setName(account.getName())
                    .build();
            responseObserver.onNext(accountResponse);
        }
        responseObserver.onCompleted();
    }

    @Override
    public void saveAccount(AccountMessage request, StreamObserver<SaveAccountResponse> responseObserver) {
        final var account = new Account();
        account.setDiscordId(request.getDiscordId());
        account.setName(request.getName());
        var accountSave = accountService.saveOrUpdateAccount(account);
        responseObserver.onNext(SaveAccountResponse.newBuilder()
                .setIsSaved(accountSave.isPresent())
                .build());
        responseObserver.onCompleted();
    }

    @Override
    public void getAccountParcels(AccountReq request, StreamObserver<AccountParcelMessage> responseObserver) {
        final var account = accountService.getAccountById(request.getDiscordId());
        if (account.isPresent()) {
            for (var parcel : account.get().getParcels()) {
                responseObserver.onNext(AccountParcelMessage.newBuilder()
                        .setUuid(parcel.getUuid())
                        .setName(parcel.getName())
                        .setDestination(parcel.getDestination())
                        .setLastUpdate(parcel.getLastUpdate().toString())
                        .setOrigin(parcel.getOrigin())
                        .setTrackingCode(parcel.getTrackingCode())
                        .setStatus(parcel.getStatus())
                        .setZipCode(parcel.getZipCode())
                        .setIsDone(parcel.isDone())
                        .build());
            }
            responseObserver.onCompleted();
        } else {
            responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
        }
    }

    @Override
    public void saveParcel(AccountSaveParcel request, StreamObserver<SaveAccountResponse> responseObserver) {
        final var parcel = new Parcel();
        parcel.setUuid(request.getUuid());
        parcel.setName(request.getName());
        parcel.setOrigin(request.getOrigin());
        parcel.setDestination(request.getDestination());
        parcel.setLastUpdate(Timestamp.valueOf(request.getLastUpdate()));
        parcel.setTrackingCode(request.getTrackingCode());
        parcel.setStatus(request.getStatus());
        parcel.setZipCode(request.getZipCode());
        parcel.setDone(request.getIsDone());
        final var account = accountService.saveOrUpdateParcel(request.getDiscordId(), parcel);
        responseObserver.onNext(SaveAccountResponse.newBuilder()
                .setIsSaved(account.isPresent())
                .build());
        responseObserver.onCompleted();
    }
}
