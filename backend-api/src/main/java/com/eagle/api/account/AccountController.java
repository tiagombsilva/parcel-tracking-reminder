package com.eagle.api.account;

import com.eagle.api.parcel.Parcel;
import jakarta.websocket.server.PathParam;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;

import java.util.Collection;
import java.util.Optional;

public class AccountController {

    private final AccountService service;

    @Autowired
    public AccountController(final AccountService service) {
        this.service = service;
    }

    @GetMapping("/accounts")
    public ResponseEntity<Collection<Account>> getAllParcels() {
        return ResponseEntity.ok(service.getAllAccounts());
    }

    @GetMapping("/account/{id}")
    public ResponseEntity<Account> getParcelById(@PathParam("id") Long parcelId) {
        return service.getAccount(parcelId).map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }

    @PostMapping("/account")
    public ResponseEntity<Account> saveOrUpdate(@RequestBody Account account) {
        final Optional<Account> savedAccount = service.saveOrUpdateAccount(account);
        return savedAccount.map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.BAD_REQUEST));
    }

    @DeleteMapping("/account/{id}")
    public ResponseEntity<Account> deleteParcel(@PathParam("id") Long accountId) {
        return service.deleteAccount(accountId).map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }

    @GetMapping("/account/{id}/parcels")
    public ResponseEntity<Collection<Parcel>> getAllParcelsFromAccount(@PathParam("id") Long accountId) {
        return service.getAllParcelsFromAccount(accountId).map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }
}
