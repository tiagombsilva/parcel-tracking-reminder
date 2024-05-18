package com.eagle.rest.account;

import com.eagle.rest.parcel.Parcel;
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

@RestController
@RequestMapping("/api/${api.version}/account")
public class AccountController {

    private final AccountService accountService;

    @Autowired
    public AccountController(final AccountService accountService) {
        this.accountService = accountService;
    }

    @GetMapping()
    public ResponseEntity<Collection<Account>> getAllAccounts() {
        return ResponseEntity.ok(accountService.getAllAccounts());
    }

    @GetMapping("{discordId}/parcels")
    public ResponseEntity<Collection<Parcel>> getAccountParcels(@PathVariable("discordId") String discordId) {
        return ResponseEntity.ok(accountService.getAccountParcels(discordId));
    }

    @PostMapping("{discordId}/parcel")
    public ResponseEntity<Parcel> saveOrUpdateParcel(@PathVariable("discordId") String discordId,
                                                     @RequestBody Parcel parcel) {
        return accountService.saveOrUpdateParcel(discordId, parcel).map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }

    @PostMapping()
    public ResponseEntity<Account> saveOrUpdateAccount(@RequestBody Account account) {
        final Optional<Account> savedAccount = accountService.saveOrUpdateAccount(account);
        return savedAccount.map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.BAD_REQUEST));
    }

    @DeleteMapping("{discordId}")
    public ResponseEntity<Account> deleteAccount(@PathVariable("discordId") String discordId) {
        return accountService.deleteAccount(discordId).map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }
}
