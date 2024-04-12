package com.eagle.api.website;

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
import java.util.Optional;

@RestController
@RequestMapping("/api/${api.version}/website")
public class WebsiteController {

    private final WebsiteService service;

    @Autowired
    public WebsiteController(final WebsiteService service) {
        this.service = service;
    }

    @GetMapping()
    public ResponseEntity<Collection<Website>> getAllWebsites() {
        return ResponseEntity.ok(service.getAllWebsites());
    }

    @GetMapping("{domain}")
    public ResponseEntity<Website> getWebsiteById(@PathVariable("domain") String domain) {
        return service.getWebsite(domain).map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }

    @PostMapping()
    public ResponseEntity<Website> saveOrUpdate(@RequestBody Website website) {
        final Optional<Website> savedWebsite = service.saveOrUpdateWebsite(website);
        return savedWebsite.map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.BAD_REQUEST));
    }

    @DeleteMapping("{domain}")
    public ResponseEntity<Website> deleteParcel(@PathVariable("domain") String domain) {
        return service.deleteWebsite(domain).map(ResponseEntity::ok)
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }
}
