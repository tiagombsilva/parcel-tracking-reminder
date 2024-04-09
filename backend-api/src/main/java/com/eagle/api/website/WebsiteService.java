package com.eagle.api.website;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Collection;
import java.util.Optional;

@Service
public class WebsiteService {

    private final WebsiteRepository repository;

    @Autowired
    public WebsiteService(final WebsiteRepository repository) {
        this.repository = repository;
    }

    public Collection<Website> getAllWebsites() {
        return repository.findAll();
    }

    public Optional<Website> getWebsite(final String domain) {
        return repository.findById(domain);
    }

    public Optional<Website> saveOrUpdateWebsite(final Website website) {
        return Optional.of(repository.save(website));
    }

    public Optional<Website> deleteWebsite(final String domain) {
        final Optional<Website> websiteOptional = repository.findById(domain);
        websiteOptional.ifPresent(p -> repository.deleteById(domain));
        return websiteOptional;
    }
}
