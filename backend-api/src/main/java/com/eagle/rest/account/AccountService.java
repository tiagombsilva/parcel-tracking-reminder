package com.eagle.rest.account;

import com.eagle.rest.parcel.Parcel;
import com.eagle.rest.parcel.ParcelService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Collection;
import java.util.List;
import java.util.Optional;

@Service
public class AccountService {

    private final AccountRepository repository;
    private final ParcelService parcelService;

    @Autowired
    public AccountService(final AccountRepository repository, final ParcelService parcelService) {
        this.repository = repository;
        this.parcelService = parcelService;
    }

    public Collection<Account> getAllAccounts() {
        return repository.findAll();
    }

    public Optional<Account> getAccount(final String discordId) {
        return repository.findById(discordId);
    }

    public Optional<Account> saveOrUpdateAccount(final Account account) {
        return Optional.of(repository.save(account));
    }

    public Optional<Account> deleteAccount(final String discordId) {
        final Optional<Account> AccountOptional = repository.findById(discordId);
        AccountOptional.ifPresent(p -> repository.deleteById(discordId));
        return AccountOptional;
    }

    public Optional<Account> getAccountById(final String discordId) {
        return repository.findById(discordId);
    }

    public Collection<Parcel> getAccountParcels(final String discordId) {
        final Optional<Account> optionalAccount = getAccountById(discordId);
        if (optionalAccount.isPresent()) {
            return optionalAccount.get().getParcels();
        }
        return List.of();
    }

    public Optional<Parcel> saveOrUpdateParcel(final String discordId, final Parcel parcel) {
        final Optional<Account> optionalAccount = getAccountById(discordId);
        parcelService.saveOrUpdateParcel(parcel);
        if (optionalAccount.isPresent()) {
            var acc = optionalAccount.get();
            acc.getParcels().add(parcel);
            saveOrUpdateAccount(acc);
            return Optional.of(parcel);
        }
        return Optional.empty();
    }
}
