package com.eagle.api.account;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class AccountService {

	private final AccountRepository repository;

	@Autowired
	public AccountService(final AccountRepository repository) {
		this.repository = repository;
	}
}
