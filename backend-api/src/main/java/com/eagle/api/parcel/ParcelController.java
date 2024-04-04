package com.eagle.api.parcel;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.servlet.support.ServletUriComponentsBuilder;

import java.net.URI;
import java.util.ArrayList;
import java.util.Collection;

@RestController
@RequestMapping("/api")
public class ParcelController {

	private final ParcelRepository parcelRepository;

	@Autowired
	public ParcelController(ParcelRepository parcelRepository) {
		this.parcelRepository = parcelRepository;
	}

	@GetMapping("/parcels")
	public ResponseEntity<Collection<Parcel>> getAllParcels(){
		final Collection<Parcel> parcels = new ArrayList<>();
		parcelRepository.findAll().addAll(parcels);
		return ResponseEntity.ok(parcels);
	}

	@PostMapping("/parcel")
	public ResponseEntity<Parcel> newParcel(@RequestBody Parcel parcel){
		final Parcel savedParcel = parcelRepository.save(parcel);
		URI location = ServletUriComponentsBuilder
				.fromCurrentRequest()
				.path("/{id}")
				.buildAndExpand(savedParcel.getId())
				.toUri();
		return ResponseEntity.created(location).body(savedParcel);
	}
}
