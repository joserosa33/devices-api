package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/devices/src/handlers"
	"github.com/devices/src/models"
	"github.com/devices/src/services"

	"github.com/gorilla/mux"
)

type deviceController struct {
	deviceService services.DeviceService
	errorHandler  handlers.ErrorHandler
}

type DeviceController interface {
	Add(w http.ResponseWriter, r *http.Request)
	All(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetByBrand(w http.ResponseWriter, r *http.Request)
}

func NewDeviceController(s services.DeviceService, errorHandler handlers.ErrorHandler) DeviceController {
	return &deviceController{
		deviceService: s,
		errorHandler:  errorHandler,
	}
}

func (d *deviceController) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stringId := params["id"]

	id, err := strconv.Atoi(stringId)

	if d.errorHandler.HandleError(err, "DeviceController/Get") {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var device = d.deviceService.GetById(id)

	if device == (models.Device{}) {
		http.Error(w, "Error getting device", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&device)
}

func (d *deviceController) Add(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	_ = json.NewDecoder(r.Body).Decode(&device)

	id := d.deviceService.Add(device)

	if id < 0 {
		http.Error(w, "Error adding to database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Item added wiht id: " + fmt.Sprint(id))
}

func (d *deviceController) All(w http.ResponseWriter, r *http.Request) {
	var all = d.deviceService.GetAll()
	json.NewEncoder(w).Encode(all)
}

func (d *deviceController) Update(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	_ = json.NewDecoder(r.Body).Decode(&device)

	id := d.deviceService.Update(device)

	if id < 0 {
		http.Error(w, "Error updating item", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Item updated wiht id: " + fmt.Sprint(id))
}

func (d *deviceController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stringId := params["id"]
	idToDelete, err := strconv.Atoi(stringId)

	if d.errorHandler.HandleError(err, "DeviceController/Get") {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id := d.deviceService.DeleteById(idToDelete)

	if id < 0 {
		http.Error(w, "Error deleting item", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Item deleted wiht id: " + fmt.Sprint(id))
}

func (d *deviceController) GetByBrand(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	brand := params["brand"]

	var all = d.deviceService.GetByBrand(brand)

	json.NewEncoder(w).Encode(all)
}
