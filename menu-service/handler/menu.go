package handler

import (
	"Pengenalan-Microservice/database"
	"Pengenalan-Microservice/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"
)

type MenuHandler struct {
	Db *gorm.DB
}

func (handler *MenuHandler) AddMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	var menu database.Menu
	err = json.Unmarshal(body, &menu)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	err = menu.Insert(handler.Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
	}

	utils.WrapAPISuccess(w, r, "success", http.StatusOK)
}

func (handler *MenuHandler) GetMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	menu := database.Menu{}
	menus, err := menu.GetAll(handler.Db)
	if err != nil {
		utils.WrapAPIError(w, r, "failed get menu:"+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WrapAPIData(w, r, menus, 200, "success")
}
