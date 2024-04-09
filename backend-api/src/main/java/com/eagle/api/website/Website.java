package com.eagle.api.website;

import com.eagle.api.parcel.Parcel;
import jakarta.annotation.Nullable;
import jakarta.persistence.CascadeType;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.Id;
import jakarta.persistence.OneToMany;
import jakarta.persistence.Table;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import lombok.ToString;

import java.util.Collection;

@Getter
@Setter
@ToString
@AllArgsConstructor
@NoArgsConstructor
@Builder
@Entity
@Table(name = "website")
public class Website {

    @Id
    private String domain;

    @Column(name = "track_url")
    private String trackUrl;

    @Column(name = "classPath")
    private String clazz;

    @OneToMany(cascade = CascadeType.ALL, fetch = FetchType.LAZY)
    @Nullable
    private Collection<Parcel> parcels;
}
