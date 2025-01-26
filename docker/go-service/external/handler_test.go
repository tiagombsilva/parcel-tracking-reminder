package external

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHandlerPost(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got '%s'", r.Method)
		}

		if r.Header.Get("Content-type") != "application/json" {
			t.Errorf("Expected content-type to be application/json, got '%s'", r.Header.Get("Content-type"))
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}

		var req ParcelRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			t.Errorf("Expected a valid JSON body, got '%s'", string(body))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success"}`))
	}))

	defer testServer.Close()

	os.Setenv("parcelsApiToken", "test_api_key")

	handler := NewParcelHandler(testServer.Client(), testServer.URL)

	req := &Request{
		TrackingId:  "123",
		DestCountry: "testCountry",
		Zipcode:     "9760-123",
	}

	response, err := handler.PostParcel(req)
	if err != nil {
		t.Fatalf("Unexpected error, got '%s'", err)
	}

	expectedResponse := `{"status":"success"}`
	if strings.TrimSpace(string(response)) != expectedResponse {
		t.Errorf("Expected response '%s', but got '%s'", expectedResponse, string(response))
	}
}
