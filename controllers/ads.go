package controllers

import (
	"encoding/json"
	"github.com/bandgren/classified-ads/database"
	"github.com/bandgren/classified-ads/entities"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateAd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payload entities.AdPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if payload.Email == "" || payload.Subject == "" || payload.Body == "" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ad := entities.Ad{
		Email:   payload.Email,
		Subject: payload.Subject,
		Price:   payload.Price,
		Body:    payload.Body,
	}
	database.Instance.Create(&ad)
	err = json.NewEncoder(w).Encode(ad)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func GetAds(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	var ads []entities.Ad
	sortBy := v.Get("sortBy")
	if sortBy == "" {
		sortBy = "created_at"
	}
	orderBy := v.Get("orderBy")
	if orderBy == "" {
		orderBy = "desc"
	}
	database.Instance.Order(sortBy + " " + orderBy).Find(&ads)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(ads)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DeleteAd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	adId := mux.Vars(r)["id"]
	var ad entities.Ad
	test := database.Instance.Delete(&ad, adId)
	if test.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode("Ad Not Found!")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	err := json.NewEncoder(w).Encode("Ad Deleted Successfully!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
