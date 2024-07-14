package com.eagle.rest.exception;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(HttpStatus.NOT_MODIFIED)
public class ResourceHasBondsException extends Exception {

    public ResourceHasBondsException(final String message) {
        super(message);
    }
}
