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

// CreateSite godoc
// @Summary      Create a new site
// @Description  Adds a new site
// @Tags         sites
// @Accept       json
// @Produce      json
// @Param        site  body      Site  true  "Site to create"
// @Success      201    {object}  Site
// @Failure      400    {object}  map[string]string "Invalid input"
// @Router       /sites/v1.0 [post]
func SaveSiteInfo(writer http.ResponseWriter, request *http.Request) {
	//store the site info to db
	db := MySQLConnectionHelper()
	//defer db.Close()

	//parse the request body
	var site Site
	err := json.NewDecoder(request.Body).Decode(&site)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	//save the site info to db
	result := db.Create(&site)
	if result.Error != nil {
		http.Error(writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	//return the site info to client
	json.NewEncoder(writer).Encode(site)

}

// GetAllSites godoc
// @Summary      Get all sites
// @Description  Returns list of sites
// @Tags         sites
// @Accept       json
// @Produce      json
// @Success      200  {array}   Site
// @Router       /sites/v1.0 [get]
func GetAllSiteInfo(writer http.ResponseWriter, request *http.Request) {
	var sites []Site
	db := MySQLConnectionHelper()
	result := db.Find(&sites)
	if result.Error != nil {

		http.Error(writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	//return the site info to client
	json.NewEncoder(writer).Encode(sites)

}

// GetSiteById godoc
// @Summary Get details of requested site
// @Description Get details of requestesd site
// @Tags sites
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the Site"
// @Success 200 {object} Site
// @Failure 400 {object} map[string]string "Invalid ID supplied"
// @Failure 404 {object} map[string]string "Site not found"
// @Router /sites/v1.0/{id} [get]
func GetSiteInfoByID(writer http.ResponseWriter, request *http.Request) {
	var site Site
	db := MySQLConnectionHelper()
	result := db.First(&site, request.URL.Query().Get("id"))
	if result.Error != nil {

		http.Error(writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	//return the site info to client
	json.NewEncoder(writer).Encode(site)

}

// UpdateSite godoc
// @Summary Update existing site
// @Description Update existing site with the input payload
// @Tags sites
// @Accept  json
// @Produce  json
// @Param site body Site true "Update site"
// @Success 200 {object} Site
// @Router /sites/v1.0 [put]
func UpdateSiteInfoByID(writer http.ResponseWriter, request *http.Request) {
	var site Site
	db := MySQLConnectionHelper()
	err := json.NewDecoder(request.Body).Decode(&site)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := db.Model(&site).Updates(Site{Name: site.Name, Location: site.Location, Status: site.Status})
	//send result
	if result.Error != nil {
		http.Error(writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(site)

}

// DeleteSite godoc
// @Summary Delete site
// @Description Delete site by id
// @Tags sites
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the Site"
// @Success 200 {object} Site
// @Router /sites/v1.0/{id} [delete]
func DeleteSiteInfoByID(writer http.ResponseWriter, request *http.Request) {
	db := MySQLConnectionHelper()
	result := db.Delete(&Site{}, request.URL.Query().Get("id"))
	if result.Error != nil {
		http.Error(writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode("Site deleted successfully")

}
