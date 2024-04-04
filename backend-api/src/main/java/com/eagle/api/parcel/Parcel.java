package com.eagle.api.parcel;

import com.eagle.api.account.Account;
import jakarta.annotation.Nullable;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.Table;
import lombok.Getter;
import lombok.Setter;

import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name = "parcel")
public class Parcel {
	@Id
	@GeneratedValue(strategy = GenerationType.AUTO)
	long id;
	@Column(name = "tracking")
	String tracking;
	@Column(name = "name")
	@Nullable
	String name;
	@Column(name = "last_update")
	Timestamp lastUpdate;
	@Column(name = "domain")
	@Nullable
	String domain;
	String state;
	@ManyToOne(fetch = FetchType.LAZY)
	@JoinColumn(name = "account_id")
	Account account;

	@Override
	public String toString() {
		return String.format("Parcel{id=%d, tracking=%s, name=%s, lastUpdate=%s, state=%s, account=%s}",
				id, tracking, name, lastUpdate, state, account);
	}
}
