package store

import (
	"encoding/json"
	"net/http"
)

type Site struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Location string `gorm:"type:varchar(100);not null"`
	Status   bool   `gorm:"not null;default:true"`
}

func SaveSiteInfo(writer *http.ResponseWriter, request *http.Request) {
	//store the site info to db
	db := MySQLConnectionHelper()
	//defer db.Close()

	//parse the request body
	var site Site
	err := json.NewDecoder(request.Body).Decode(&site)
	if err != nil {
		http.Error(*writer, err.Error(), http.StatusBadRequest)
		return
	}
	//save the site info to db
	result := db.Create(&site)
	if result.Error != nil {
		http.Error(*writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	//return the site info to client
	json.NewEncoder(*writer).Encode(site)

}

func GetAllSiteInfo(writer *http.ResponseWriter, request *http.Request) {

}

func GetSiteInfoByID(writer *http.ResponseWriter, request *http.Request) {

}

func UpdateSiteInfoByID(writer *http.ResponseWriter, request *http.Request) {

}

func DeleteSiteInfoByID(writer *http.ResponseWriter, request *http.Request) {

}
