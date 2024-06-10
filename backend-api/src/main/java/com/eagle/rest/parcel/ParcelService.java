package com.eagle.rest.parcel;

import com.eagle.rest.exception.ResourceHasBondsException;
import com.eagle.rest.exception.ResourceNotFoundException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Collection;
import java.util.Optional;

@Service
public class ParcelService {

    private final ParcelRepository repository;

    @Autowired
    public ParcelService(final ParcelRepository repository) {
        this.repository = repository;
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
        return repository.save(parcel);
    }

    public void deleteParcel(final String parcelId) throws ResourceHasBondsException {
        final Parcel parcel = repository.findById(parcelId).orElseThrow(
                () -> new ResourceNotFoundException("Parcel not found"));
        try {
            repository.deleteById(parcelId);
        } catch (Exception e) {
            throw new ResourceHasBondsException("Parcel not found");
        }
    }

    public Optional<Parcel> getParcelByTrackingCode(final String trackingCode) {
        return repository.findByTrackingCode(trackingCode);
    }
}
