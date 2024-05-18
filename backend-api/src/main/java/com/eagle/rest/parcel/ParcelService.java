package com.eagle.rest.parcel;

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

    public Optional<Parcel> getParcel(final Long uuid) {
        return repository.findById(uuid);
    }

    public Optional<Parcel> saveOrUpdateParcel(final Parcel parcel) {
        return Optional.of(repository.save(parcel));
    }

    public Optional<Parcel> deleteParcel(final Long parcelId) {
        final Optional<Parcel> parcelOptional = repository.findById(parcelId);
        parcelOptional.ifPresent(p -> repository.deleteById(parcelId));
        return parcelOptional;
    }

    public Optional<Parcel> getParcelByTrackingCode(final String trackingCode) {
        return repository.findByTrackingCode(trackingCode);
    }
}
