package com.eagle.rest.parcel;

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

@RestController
@RequestMapping("/api/${api.version}/parcel")
public class ParcelController {

    private final ParcelService service;

    @Autowired
    public ParcelController(final ParcelService service) {
        this.service = service;
    }

    @GetMapping()
    public ResponseEntity<Collection<Parcel>> getAllParcels() {
        return ResponseEntity.ok(service.getAllParcels());
    }

    @GetMapping("/inProgress")
    public ResponseEntity<Collection<Parcel>> getAllParcelsInProgress() {
        return ResponseEntity.ok(service.getAllParcelsInProgress());
    }

    @GetMapping("{id}")
    public ResponseEntity<Parcel> getParcelById(@PathVariable("id") String parcelUuid) {
        return service.getParcel(parcelUuid).map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }

    @PostMapping()
    public Parcel saveOrUpdate(@RequestBody Parcel parcel) {
        return service.saveOrUpdateParcel(parcel);
    }

    @DeleteMapping("{id}")
    public ResponseEntity<Void> deleteParcel(@PathVariable("id") String parcelUuid) {
        service.deleteParcel(parcelUuid);
        return ResponseEntity.noContent().build();
    }
}
