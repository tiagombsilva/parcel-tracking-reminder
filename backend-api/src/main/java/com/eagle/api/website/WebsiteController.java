package com.eagle.api.website;

import jakarta.websocket.server.PathParam;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;

import java.util.Collection;
import java.util.Optional;

public class WebsiteController {

    private final WebsiteService service;

    @Autowired
    public WebsiteController(final WebsiteService service) {
        this.service = service;
    }

    @GetMapping("/websites")
    public ResponseEntity<Collection<Website>> getAllWebsites() {
        return ResponseEntity.ok(service.getAllWebsites());
    }

    @GetMapping("/parcel/{domain}")
    public ResponseEntity<Website> getWebsiteById(@PathParam("domain") String domain) {
        return service.getWebsite(domain).map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }

    @PostMapping("/website")
    public ResponseEntity<Website> saveOrUpdate(@RequestBody Website website) {
        final Optional<Website> savedWebsite = service.saveOrUpdateWebsite(website);
        return savedWebsite.map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.BAD_REQUEST));
    }

    @DeleteMapping("/website/{domain}")
    public ResponseEntity<Website> deleteParcel(@PathParam("domain") String domain) {
        return service.deleteWebsite(domain).map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }
}
