package com.eagle.api.parcel;

import jakarta.websocket.server.PathParam;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Collection;
import java.util.Optional;

@RestController
@RequestMapping("/api")
public class ParcelController {

	private final ParcelService service;

	@Autowired
	public ParcelController(final ParcelService service) {
		this.service = service;
	}

	@GetMapping("/parcels")
	public ResponseEntity<Collection<Parcel>> getAllParcels() {
		return ResponseEntity.ok(service.getAllParcels());
	}

	@GetMapping("/parcel/{id}")
	public ResponseEntity<Parcel> getParcelById(@PathParam("id") Long parcelId) {
		return service.getParcel(parcelId).map(ResponseEntity::ok)
				.orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
	}

	@PostMapping("/parcel")
	public ResponseEntity<Parcel> saveOrUpdate(@RequestBody Parcel parcel) {
		final Optional<Parcel> savedParcel = service.saveOrUpdateParcel(parcel);
		return savedParcel.map(ResponseEntity::ok)
				.orElse(new ResponseEntity<>(HttpStatus.BAD_REQUEST));
	}

	@DeleteMapping("/parcel/{id}")
	public ResponseEntity<Parcel> deleteParcel(@PathParam("id") Long parcelId) {
		return service.deleteParcel(parcelId).map(ResponseEntity::ok)
				.orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
	}

}
