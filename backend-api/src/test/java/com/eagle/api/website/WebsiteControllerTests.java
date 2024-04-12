package com.eagle.api.website;

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

@WebMvcTest(WebsiteController.class)
public class WebsiteControllerTests {

    @Autowired
    private MockMvc mockMvc;
    @MockBean
    private WebsiteService service;
    @Autowired
    private ObjectMapper mapper;

    @Test
    public void getAllWebsites() throws Exception {
        final var website = getDummyWebsite("dummy");
        given(service.getAllWebsites()).willReturn(List.of(website));
        final var parcelJson = getJson(website);
        mockMvc.perform(MockMvcRequestBuilders.get("/api/v1/website")).andDo(print())
                .andExpect(status().isOk())
                .andExpect(content().json("[%s]".formatted(parcelJson)));
    }

    @Test
    public void getWebsiteByDomain() throws Exception {
        final var website = getDummyWebsite("dummy");
        given(service.getWebsite("dummy")).willReturn(Optional.of(website));
        mockMvc.perform(MockMvcRequestBuilders.get("/api/v1/website/{domain}", "dummy"))
                .andExpect(status().isOk())
                .andExpect(content().json("%s".formatted(getJson(website))));
    }

    @Test
    public void saveOrUpdate() throws Exception {
        final var website = getDummyWebsite("dummy");
        final var parcelJson = getJson(website);
        given(service.saveOrUpdateWebsite(any())).willReturn(Optional.of(website));
        mockMvc.perform(post("/api/v1/website")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(parcelJson))
                .andExpect(status().isOk())
                .andExpect(content().json("%s".formatted(getJson(website))));
    }

    @Test
    public void deleteWebsite() throws Exception {
        final var website = getDummyWebsite("dummy");
        given(service.deleteWebsite("dummy")).willReturn(Optional.of(website));
        mockMvc.perform(MockMvcRequestBuilders.delete("/api/v1/website/{domain}", "dummy"))
                .andExpect(status().isOk())
                .andExpect(content().json("%s".formatted(getJson(website))));
    }

    private String getJson(final Website parcel) throws JsonProcessingException {
        return mapper.writeValueAsString(parcel);
    }

    private Website getDummyWebsite(final String domain) {
        return Website.builder().domain(domain).build();
    }

}
