package com.eagle.rest.account;

import com.eagle.rest.exception.ResourceNotFoundException;
import com.eagle.rest.parcel.Parcel;
import com.eagle.rest.parcel.ParcelService;
import io.micrometer.common.util.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Collection;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

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

    public void deleteAccount(final String discordId) {
        final Account account = repository.findById(discordId).orElseThrow(
                () -> new ResourceNotFoundException("Parcel not found")
        );
        repository.delete(account);
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

    public Parcel saveOrUpdateParcel(final String discordId, final Parcel parcel) throws
            ResourceNotFoundException {
        final Account account = getAccountById(discordId)
                .orElseThrow(() -> new ResourceNotFoundException("Account not found"));
        parcel.setAccount(account);
        parcelService.saveOrUpdateParcel(parcel);
        account.getParcels().add(parcel);
        repository.save(account);
        return parcel;
    }
}
