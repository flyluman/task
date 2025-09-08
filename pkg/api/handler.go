package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task/internal/logger"
	"task/internal/services"
)

func GetUserRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()
	// set header -> application/json
	w.Header().Set("Content-Type", "application/json")

	// extract user_id from url
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"success": "false", "restaurants": "", "error": err.Error()})
		return
	}

	// get restaurants
	restaurants, err := services.GetUserRestaurants(id)

	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"success": "false", "restaurants": "", "error": err.Error()})
		return
	}

	// convirts restaurant slices into byte
	ret, err := json.Marshal(restaurants)

	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"success": "false", "restaurants": "", "error": err.Error()})
		return
	}

	// return result
	err = json.NewEncoder(w).Encode(map[string]string{"success": "true", "restaurants": string(ret), "error": ""})

	if err != nil {
		log.Error(err.Error())
	}
}

type Request struct {
	UserID     int `json:"user_id"`
	MenuItemID int `json:"menu_item_id"`
}

func PurchaseMenuItemHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()
	// set header -> application/json
	w.Header().Set("Content-Type", "application/json")

	// define request json
	var req Request

	// decode json from req body
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"success": "false", "error": err.Error()})
		return
	}

	// purchase(atomic)
	err = services.PurchaseMenuItem(req.UserID, req.MenuItemID)

	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"success": "false", "error": err.Error()})
		return
	}

	// return result
	err = json.NewEncoder(w).Encode(map[string]string{"success": "true", "error": ""})

	if err != nil {
		log.Error(err.Error())
	}
}
