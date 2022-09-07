package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/bandgren/classified-ads/database"
	"github.com/bandgren/classified-ads/entities"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

func TestCreateAd(t *testing.T) {
	os.Setenv("SQL_CONNECTION_STRING", "root:root@tcp(127.0.0.1:3306)/ads?parseTime=true")
	database.Start()
	payload := entities.AdPayload{
		Email:   "JohannesTestar.test.se",
		Subject: "Riktigt fin bil",
		Price:   999999,
		Body:    "cString",
	}
	data, _ := json.Marshal(payload)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/ads", bytes.NewBuffer(data))

	CreateAd(rr, req)
	if rr.Result().StatusCode != 200 {
		t.Errorf("Status code returned, %d, did not match expected code %d", rr.Result().StatusCode, 200)
	}
}

func TestGetAds(t *testing.T) {
	os.Setenv("SQL_CONNECTION_STRING", "root:root@tcp(127.0.0.1:3306)/ads?parseTime=true")
	database.Start()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/ads", nil)

	GetAds(rr, req)
	if rr.Result().StatusCode != 200 {
		t.Errorf("Status code returned, %d, did not match expected code %d", rr.Result().StatusCode, 200)
	}
	var ads []entities.Ad
	err := json.NewDecoder(rr.Body).Decode(&ads)
	if err != nil {
		t.Errorf("failed to decode response body")
	}
	if len(ads) == 0 {
		t.Errorf("No ads was returned from GetAds")
	}
}

func TestDeleteAd(t *testing.T) {
	os.Setenv("SQL_CONNECTION_STRING", "root:root@tcp(127.0.0.1:3306)/ads?parseTime=true")
	database.Start()
	payload := entities.AdPayload{
		Email:   "JohannesTestar.test.se",
		Subject: "Riktigt fin bil",
		Price:   999999,
		Body:    "cString",
	}
	data, _ := json.Marshal(payload)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/ads", bytes.NewBuffer(data))
	CreateAd(rr, req)
	var ad entities.Ad
	err := json.NewDecoder(rr.Body).Decode(&ad)
	if err != nil {
		t.Errorf("failed to decode response body")
	}

	rr = httptest.NewRecorder()
	req = httptest.NewRequest("DELETE", "/api/ads", nil)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(uint64(ad.ID), 10)})

	DeleteAd(rr, req)
	if rr.Result().StatusCode != 200 {
		t.Errorf("Status code returned, %d, did not match expected code %d", rr.Result().StatusCode, 200)
	}
}
