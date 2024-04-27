package com.eagle.rest.account;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Collection;
import java.util.Optional;

@Service
public class AccountService {

    private final AccountRepository repository;

    @Autowired
    public AccountService(final AccountRepository repository) {
        this.repository = repository;
    }

    public Collection<Account> getAllAccounts() {
        return repository.findAll();
    }

    public Optional<Account> getAccount(final Long accountId) {
        return repository.findById(accountId);
    }

    public Optional<Account> saveOrUpdateAccount(final Account account) {
        return Optional.of(repository.save(account));
    }

    public Optional<Account> deleteAccount(final Long accountId) {
        final Optional<Account> AccountOptional = repository.findById(accountId);
        AccountOptional.ifPresent(p -> repository.deleteById(accountId));
        return AccountOptional;
    }
}
