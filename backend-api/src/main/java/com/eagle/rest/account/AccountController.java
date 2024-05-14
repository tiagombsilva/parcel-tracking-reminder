package com.eagle.rest.account;

import com.eagle.rest.parcel.Parcel;
import com.eagle.rest.parcel.ParcelService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Collection;
import java.util.Optional;
import java.util.Set;

@RestController
@RequestMapping("/api/${api.version}/account")
public class AccountController {

    private final AccountService accountService;
    private final ParcelService parcelService;

    @Autowired
    public AccountController(final AccountService accountService, final ParcelService parcelService) {
        this.accountService = accountService;
        this.parcelService = parcelService;
    }

    @GetMapping()
    public ResponseEntity<Collection<Account>> getAllAccounts() {
        return ResponseEntity.ok(accountService.getAllAccounts());
    }

    @GetMapping("{id}/parcels")
    public ResponseEntity<Set<Parcel>> getAccountParcels(@PathVariable("id") Long accountId) {
        final Optional<Account> optionalAccount = accountService.getAccountById(accountId);
        return optionalAccount
                .map(Account::getParcels)
                .map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.BAD_REQUEST));
    }

    @PostMapping("{id}/parcel")
    public ResponseEntity<Parcel> saveOrUpdateParcel(@PathVariable("id") Long accountId,
                                                     @RequestBody Parcel parcel) {
        final Optional<Account> optionalAccount = accountService.getAccountById(accountId);
        parcelService.saveOrUpdateParcel(parcel);
        if (optionalAccount.isPresent()) {
            var acc = optionalAccount.get();
            acc.getParcels().add(parcel);
            accountService.saveOrUpdateAccount(acc);
            return ResponseEntity.ok(parcel);
        }
        return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
    }

    @PostMapping()
    public ResponseEntity<Account> saveOrUpdateAccount(@RequestBody Account account) {
        final Optional<Account> savedAccount = accountService.saveOrUpdateAccount(account);
        return savedAccount.map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.BAD_REQUEST));
    }

    @DeleteMapping("{id}")
    public ResponseEntity<Account> deleteAccount(@PathVariable("id") Long accountId) {
        return accountService.deleteAccount(accountId).map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }
}
