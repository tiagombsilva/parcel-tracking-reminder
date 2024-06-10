package com.eagle.rest.parcel;

import com.eagle.rest.account.Account;
import com.fasterxml.jackson.annotation.JsonIgnore;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.Table;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import lombok.ToString;
import org.springframework.lang.Nullable;

import java.sql.Timestamp;

@Getter
@Setter
@ToString
@AllArgsConstructor
@NoArgsConstructor
@Builder
@EqualsAndHashCode(onlyExplicitlyIncluded = true)
@Entity
@Table(name = "parcel")
public class Parcel {
    @Id
    private String uuid;
    private String trackingCode;
    @Nullable
    private String name;
    private String origin;
    private String destination;
    @Column(name = "last_update")
    private Timestamp lastUpdate;
    private String status;
    private boolean done;
    private String zipCode;
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "discord_id", referencedColumnName = "discord_id")
    @JsonIgnore
    private Account account;
}
