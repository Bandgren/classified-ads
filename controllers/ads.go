package controllers

import (
	"encoding/json"
	"github.com/bandgren/classified-ads/database"
	"github.com/bandgren/classified-ads/entities"
	"net/http"
)

func CreateAd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ad entities.Ad
	json.NewDecoder(r.Body).Decode(&ad)
	database.Instance.Create(&ad)
	json.NewEncoder(w).Encode(ad)
}
func GetAds(w http.ResponseWriter, r *http.Request) {
	var ads []entities.Ad
	database.Instance.Find(&ads)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ads)
}
