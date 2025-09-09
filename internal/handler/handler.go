package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task/internal/service"
	"task/model"
	"task/pkg/logger"
)

type RespVal map[string]interface{}

func WriteJSON(w http.ResponseWriter, status int, resp RespVal) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(resp)

	if err != nil {
		logger.Log.Error(err.Error())
	}
}

func WriteError(w http.ResponseWriter, status int, err error) {
	logger.Log.Error(err.Error())
	WriteJSON(w, status, RespVal{
		"success": false,
		"error":   err.Error(),
	})
}

func GetUserRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	// extract user_id from url
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	// get restaurants
	restaurants, err := service.GetUserRestaurants(id)

	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// convirts restaurant slices into byte
	ret, err := json.Marshal(restaurants)

	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// return result
	WriteJSON(w, http.StatusOK, RespVal{
		"success":     true,
		"restaurants": string(ret),
		"error":       "",
	})
}

func PurchaseMenuItemHandler(w http.ResponseWriter, r *http.Request) {
	// define request json
	var req model.Request

	// decode json from req body
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	// purchase(atomic)
	err = service.PurchaseMenuItem(req.UserID, req.MenuItemID)

	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// return result
	WriteJSON(w, http.StatusOK, RespVal{
		"success": true,
		"error":   "",
	})
}
