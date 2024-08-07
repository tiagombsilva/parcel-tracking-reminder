package com.eagle.rest.parcel;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.Collection;
import java.util.Optional;

public interface ParcelRepository extends JpaRepository<Parcel, String> {

    @Query(value = "SELECT * FROM Parcel WHERE trackingCode = ?1", nativeQuery = true)
    Optional<Parcel> findByTrackingCode(String trackingCode);

    @Query(value = "SELECT * FROM Parcel WHERE done = FALSE", nativeQuery = true)
    Collection<Parcel> getInProgressParcels();
}
