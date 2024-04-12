package com.eagle.api.account;

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

    private final AccountService service;

    @Autowired
    public AccountController(final AccountService service) {
        this.service = service;
    }

    @GetMapping()
    public ResponseEntity<Collection<Account>> getAllAccounts() {
        return ResponseEntity.ok(service.getAllAccounts());
    }

    @GetMapping("{id}")
    public ResponseEntity<Account> getParcelById(@PathVariable("id") Long parcelId) {
        return service.getAccount(parcelId).map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }

    @PostMapping()
    public ResponseEntity<Account> saveOrUpdate(@RequestBody Account account) {
        final Optional<Account> savedAccount = service.saveOrUpdateAccount(account);
        return savedAccount.map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.BAD_REQUEST));
    }

    @DeleteMapping("{id}")
    public ResponseEntity<Account> deleteParcel(@PathVariable("id") Long accountId) {
        return service.deleteAccount(accountId).map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }
}
