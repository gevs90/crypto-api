package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gevs90/crypto-api/api/models"
	"github.com/gevs90/crypto-api/api/routes"
	"github.com/gevs90/crypto-api/config"
	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	config.LoadAppConfig()
	routes := routes.NewRouter()
	var port string = config.AppConfig.Port
	http.ListenAndServe(port, routes)

	code := m.Run()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.NewRouter().StrictSlash(true).ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetNonExistentRecord(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/encrypt/11", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
	var m models.ResponseEncryptedText
	json.Unmarshal(response.Body.Bytes(), &m)

	if m.ID != 0 {
		t.Errorf("Expected the m.ID key of the response to be set to '0'.")
	}
}
