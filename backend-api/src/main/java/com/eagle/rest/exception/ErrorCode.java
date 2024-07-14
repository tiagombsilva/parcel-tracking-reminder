package com.eagle.rest.exception;

import lombok.Getter;

@Getter
public enum ErrorCode {
    NOT_FOUND(404),
    UNEXPECTED(500);

    final private int errorCode;

    ErrorCode(final int errorCode) {
        this.errorCode = errorCode;
    }
}
