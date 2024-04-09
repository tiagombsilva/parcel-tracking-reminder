package com.eagle.api.website;

import com.eagle.api.parcel.Parcel;
import jakarta.annotation.Nullable;
import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import java.util.Collection;

@Getter
@Setter
@ToString
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
