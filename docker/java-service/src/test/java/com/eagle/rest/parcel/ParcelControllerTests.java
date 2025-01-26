package com.eagle.rest.parcel;

import com.eagle.rest.account.Account;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;

import java.time.ZonedDateTime;
import java.util.List;
import java.util.Optional;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.BDDMockito.given;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(ParcelController.class)
public class ParcelControllerTests {

    @Autowired
    private MockMvc mockMvc;
    @MockBean
    private ParcelService service;
    @Autowired
    private ObjectMapper mapper;

    @Test
    public void getAllParcels() throws Exception {
        final var account = getDummyAccount("discordId");
        final var parcel = getDummyParcel(account);
        given(service.getAllParcels()).willReturn(List.of(parcel));
        final var parcelJson = getJson(parcel);
        mockMvc.perform(MockMvcRequestBuilders.get("/api/v1/parcel")).andDo(print())
                .andExpect(status().isOk())
                .andExpect(content().json("[%s]".formatted(parcelJson)));
    }

    @Test
    public void getParcelById() throws Exception {
        final var account = getDummyAccount("discordId");
        final var parcel = getDummyParcel(account);
        given(service.getParcel("1")).willReturn(Optional.of(parcel));
        mockMvc.perform(MockMvcRequestBuilders.get("/api/v1/parcel/{id}", 1L))
                .andExpect(status().isOk())
                .andExpect(content().json("%s".formatted(getJson(parcel))));
    }

    @Test
    public void saveOrUpdate() throws Exception {
        final var account = getDummyAccount("discordId");
        final var parcel = getDummyParcel(account);
        final var parcelJson = getJson(parcel);
        given(service.saveOrUpdateParcel(any())).willReturn(parcel);
        mockMvc.perform(post("/api/v1/parcel")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(parcelJson))
                .andExpect(status().isOk())
                .andExpect(content().json("%s".formatted(getJson(parcel))));
    }

    @Test
    public void deleteParcel() throws Exception {
        mockMvc.perform(MockMvcRequestBuilders.delete("/api/v1/parcel/{id}", 1L))
                .andExpect(status().isNoContent());
    }

    private String getJson(final Parcel parcel) throws JsonProcessingException {
        return mapper.writeValueAsString(parcel);
    }

    private Account getDummyAccount(String discordId) {
        return Account.builder().discordId(discordId).name("dummy account").build();
    }

    private Parcel getDummyParcel(final Account account) {
        return Parcel.builder()
                .uuid("123l123l123")
                .name("Dummy parcel")
                .trackingCode("123123")
                .origin("PT")
                .destination("PT")
                .lastUpdate(ZonedDateTime.now())
                .account(account)
                .build();
    }
}
