package store

import (
	"encoding/json"
	"net/http"
)

type Site struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(100);uniqueIndex;not null"`
	Location string `json:"location" gorm:"type:varchar(100);not null"`
	Status   bool   `json:"status" gorm:"not null;default:true"`
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
	var sites []Site
	db := MySQLConnectionHelper()
	result := db.Find(&sites)
	if result.Error != nil {

		http.Error(*writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	//return the site info to client
	json.NewEncoder(*writer).Encode(sites)

}

func GetSiteInfoByID(writer *http.ResponseWriter, request *http.Request) {
	var site Site
	db := MySQLConnectionHelper()
	result := db.First(&site, request.URL.Query().Get("id"))
	if result.Error != nil {

		http.Error(*writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	//return the site info to client
	json.NewEncoder(*writer).Encode(site)

}

func UpdateSiteInfoByID(writer *http.ResponseWriter, request *http.Request) {
	var site Site
	db := MySQLConnectionHelper()
	err := json.NewDecoder(request.Body).Decode(&site)
	if err != nil {
		http.Error(*writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := db.Model(&site).Updates(Site{Name: site.Name, Location: site.Location, Status: site.Status})
	//send result
	if result.Error != nil {
		http.Error(*writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(*writer).Encode(site)

}

func DeleteSiteInfoByID(writer *http.ResponseWriter, request *http.Request) {
	db := MySQLConnectionHelper()
	result := db.Delete(&Site{}, request.URL.Query().Get("id"))
	if result.Error != nil {
		http.Error(*writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(*writer).Encode("Site deleted successfully")

}
