package com.eagle.rest.parcel;

import com.eagle.rest.account.Account;
import com.eagle.rest.account.AccountRepository;
import com.eagle.rest.exception.ResourceNotFoundException;
import io.micrometer.common.util.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Collection;
import java.util.Optional;
import java.util.UUID;

@Service
public class ParcelService {

    private final ParcelRepository repository;
    private final AccountRepository accountRepository;

    @Autowired
    public ParcelService(final ParcelRepository repository, final AccountRepository accountRepository) {
        this.repository = repository;
        this.accountRepository = accountRepository;
    }

    public Collection<Parcel> getAllParcels() {
        return repository.findAll();
    }

    public Collection<Parcel> getAllParcelsInProgress() {
        return repository.getInProgressParcels();
    }

    public Optional<Parcel> getParcel(final String uuid) {
        return repository.findById(uuid);
    }

    public Parcel saveOrUpdateParcel(final Parcel parcel) {
        var savedParcel = repository.findById(parcel.getUuid());
        if (savedParcel.isPresent() && parcel.getAccount() == null) {
            parcel.setAccount(savedParcel.get().getAccount());
        }
        if (StringUtils.isEmpty(parcel.getUuid())) {
            parcel.setUuid(UUID.randomUUID().toString());
        }
        return repository.save(parcel);
    }

    @Transactional
    public void deleteParcel(final String parcelId) {
        final Parcel parcel = repository.findById(parcelId).orElseThrow(
                () -> new ResourceNotFoundException("Parcel not found")
        );
        Account account = parcel.getAccount();

        if (account != null) {
            account.getParcels().remove(parcel);
            parcel.setAccount(null);
            accountRepository.save(account);
        }

        repository.deleteById(parcel.getUuid());
    }

    public Optional<Parcel> getParcelByTrackingCode(final String trackingCode) {
        return repository.findByTrackingCode(trackingCode);
    }
}
