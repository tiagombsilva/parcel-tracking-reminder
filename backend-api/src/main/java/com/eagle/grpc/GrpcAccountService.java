package com.eagle.grpc;

import com.eagle.grpc.accounts.lib.AccountReq;
import com.eagle.grpc.accounts.lib.AccountResponse;
import com.eagle.grpc.accounts.lib.AccountsGrpc;
import com.eagle.rest.account.AccountService;
import com.eagle.rest.parcel.ParcelService;
import com.google.protobuf.Empty;
import io.grpc.stub.StreamObserver;
import org.springframework.beans.factory.annotation.Autowired;

public class GrpcAccountService extends AccountsGrpc.AccountsImplBase {

    private final AccountService accountService;

    @Autowired
    public GrpcAccountService(final AccountService accountService, final ParcelService parcelService) {
        this.accountService = accountService;
    }

    @Override
    public void getAccounts(Empty request, StreamObserver<AccountResponse> responseObserver) {
        var accounts = accountService.getAllAccounts();
        for (var account : accounts) {
            AccountResponse accountResponse = AccountResponse.newBuilder()
                    .setId(account.getId())
                    .setDiscordId(account.getDiscordId())
                    .setName(account.getName())
                    .build();
            responseObserver.onNext(accountResponse);
        }
        responseObserver.onCompleted();
    }

    @Override
    public void getAccountById(AccountReq request, StreamObserver<AccountResponse> responseObserver) {
        var accountOptional = accountService.getAccountById(request.getId());
        if (accountOptional.isPresent()) {
            var account = accountOptional.get();
            AccountResponse accountResponse = AccountResponse.newBuilder()
                    .setId(account.getId())
                    .setDiscordId(account.getDiscordId())
                    .setName(account.getName())
                    .build();
            responseObserver.onNext(accountResponse);
        }
        responseObserver.onCompleted();
    }
}
