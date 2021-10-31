package routers

import (
	"github.com/devices/src/controllers"

	"github.com/gorilla/mux"
)

func GetDeviceRouter(deviceController controllers.DeviceController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/devices", deviceController.All).Methods("GET")
	router.HandleFunc("/devices", deviceController.Add).Methods("POST")
	router.HandleFunc("/devices/{id}", deviceController.Get).Methods("GET")
	router.HandleFunc("/devices", deviceController.Update).Methods("PUT")
	router.HandleFunc("/devices/{id}", deviceController.Delete).Methods("DELETE")
	router.HandleFunc("/devices/brands/{brand}", deviceController.GetByBrand).Methods("GET")

	return router
}
