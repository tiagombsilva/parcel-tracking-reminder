package com.eagle.rest.parcel;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.Collection;

public interface ParcelRepository extends JpaRepository<Parcel, String> {

    @Query(value = "SELECT * FROM Parcel WHERE done = FALSE", nativeQuery = true)
    Collection<Parcel> getInProgressParcels();
}
