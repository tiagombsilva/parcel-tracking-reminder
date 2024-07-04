package com.eagle.utils;

public class Helper {

    public static String getOrDefault(final Object value) {
        return value == null ? "" : value.toString();
    }
}
