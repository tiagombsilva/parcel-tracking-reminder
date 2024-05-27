package com.eagle.rest.account;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;

import java.util.List;
import java.util.Optional;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.BDDMockito.given;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(AccountController.class)
public class AccountControllerTests {

    @Autowired
    private MockMvc mockMvc;
    @MockBean
    private AccountService service;
    @Autowired
    private ObjectMapper mapper;

    @Test
    public void getAllParcels() throws Exception {
        final var account = getDummyAccount("discordId");
        given(service.getAllAccounts()).willReturn(List.of(account));
        final var parcelJson = getJson(account);
        mockMvc.perform(MockMvcRequestBuilders.get("/api/v1/account")).andDo(print())
                .andExpect(status().isOk())
                .andExpect(content().json("[%s]".formatted(parcelJson)));
    }

    @Test
    public void getAcountById() throws Exception {
        final var account = getDummyAccount("discordId");
        given(service.getAccountById("discordId")).willReturn(Optional.of(account));
        mockMvc.perform(MockMvcRequestBuilders.get("/api/v1/account/{id}", "discordId"))
                .andExpect(status().isOk())
                .andExpect(content().json("%s".formatted(getJson(account))));
    }

    @Test
    public void saveOrUpdate() throws Exception {
        final var account = getDummyAccount("discordId");
        final var parcelJson = getJson(account);
        given(service.saveOrUpdateAccount(any())).willReturn(Optional.of(account));
        mockMvc.perform(post("/api/v1/account")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(parcelJson))
                .andExpect(status().isOk())
                .andExpect(content().json("%s".formatted(getJson(account))));
    }

    @Test
    public void deleteAccount() throws Exception {
        final var account = getDummyAccount("discordId");
        given(service.deleteAccount("discordId")).willReturn(Optional.of(account));
        mockMvc.perform(MockMvcRequestBuilders.delete("/api/v1/account/{id}", "discordId"))
                .andExpect(status().isOk())
                .andExpect(content().json("%s".formatted(getJson(account))));
    }

    private String getJson(final Account parcel) throws JsonProcessingException {
        return mapper.writeValueAsString(parcel);
    }

    private Account getDummyAccount(String discordId) {
        return Account.builder().discordId(discordId).name("dummy account").build();
    }
}
