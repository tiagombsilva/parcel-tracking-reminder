package com.eagle.api;

import com.eagle.api.account.Account;
import com.eagle.api.parcel.Parcel;
import com.eagle.api.parcel.ParcelController;
import com.eagle.api.parcel.ParcelService;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;

import java.sql.Timestamp;
import java.time.Instant;
import java.util.List;

import static org.mockito.BDDMockito.given;
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
        final var dummyParcel = getDummyParcel();
        given(service.getAllParcels()).willReturn(List.of(dummyParcel));
        final String dummyParcelJson = mapper.writeValueAsString(dummyParcel);
        mockMvc.perform(MockMvcRequestBuilders.get("/api/v1/parcel"))
                .andDo(print())
                .andExpect(status().isOk())
                .andExpect(content().json("[%s]".formatted(dummyParcelJson)));
    }

    private Parcel getDummyParcel() {
        final Account account = Account.builder()
                .id(1L)
                .name("dummy account")
                .build();
        return Parcel.builder()
                .id(1L)
                .name("Dummy parcel")
                .lastUpdate(Timestamp.from(Instant.now()))
                .account(account)
                .build();
    }
}
