package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task/internal/service"
	"task/model"
)

type UserLogger interface {
	Debug(msg string, args ...any)
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}

type UserHandler struct {
	logger      UserLogger
	UserService service.UserService
}

var userHandler *UserHandler

func NewUserHandler(l UserLogger, s service.UserService) *UserHandler {
	userHandler = &UserHandler{logger: l, UserService: s}
	return userHandler
}

type RespVal map[string]interface{}

func WriteJSON(w http.ResponseWriter, status int, resp RespVal) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(resp)

	if err != nil {
		userHandler.logger.Error(err.Error())
	}
}

func WriteError(w http.ResponseWriter, status int, err error) {
	userHandler.logger.Error(strings.ReplaceAll(err.Error(), "\n", ", "))
	WriteJSON(w, status, RespVal{
		"success": false,
		"error":   strings.ReplaceAll(err.Error(), "\n", ", "),
	})
}

func (h *UserHandler) GetUserRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	// extract user id from url
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	// get restaurants
	restaurants, err := h.UserService.GetUserRestaurants(id)

	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// return result
	WriteJSON(w, http.StatusOK, RespVal{
		"success":     true,
		"restaurants": restaurants,
		"error":       "",
	})
}

func (h *UserHandler) PurchaseMenuItemHandler(w http.ResponseWriter, r *http.Request) {
	// define request json
	var req model.Request

	// decode json from req body
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	// purchase(atomic)
	err = h.UserService.PurchaseMenuItem(req.UserID, req.MenuItemID)

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
