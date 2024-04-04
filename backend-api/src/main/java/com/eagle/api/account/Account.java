package com.eagle.api.account;

import com.eagle.api.parcel.Parcel;
import jakarta.persistence.CascadeType;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.OneToMany;
import jakarta.persistence.Table;
import lombok.Getter;
import lombok.Setter;

import java.util.Set;

@Getter
@Setter
@Entity
@Table(name = "account")
public class Account{
		@Id
		@GeneratedValue(strategy = GenerationType.AUTO)
		long id;
		@Column(name = "discord_id")
		String discordId;
		String name;
		@OneToMany(cascade = CascadeType.ALL, fetch = FetchType.LAZY)
		Set<Parcel> parcels;

	@Override
	public String toString() {
		return String.format("User{id=%d, name=%s, parcels=%s}", id, name, parcels);
	}
}
