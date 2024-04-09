package com.eagle.api.parcel;

import com.eagle.api.account.Account;
import com.eagle.api.website.Website;
import jakarta.annotation.Nullable;
import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import java.sql.Timestamp;

@Getter
@Setter
@ToString
@Entity
@Table(name = "parcel")
public class Parcel {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private long id;
    @Column(name = "tracking")
    private String tracking;
    @Column(name = "name")
    @Nullable
    private String name;
    @Column(name = "last_update")
    private Timestamp lastUpdate;
    @Column(name = "domain")
    @ManyToOne
    @Nullable
    private Website website;
    private String state;
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "account_id")
    private Account account;
}
