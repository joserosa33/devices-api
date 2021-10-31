package routers

import (
	"github.com/devices/src/controllers"

	"github.com/gorilla/mux"
)

func GetDeviceRouter(deviceController controllers.DeviceController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/device", deviceController.All).Methods("GET")
	router.HandleFunc("/device", deviceController.Add).Methods("POST")
	router.HandleFunc("/device/{id}", deviceController.Get).Methods("GET")
	router.HandleFunc("/device", deviceController.Update).Methods("PUT")
	router.HandleFunc("/device/{id}", deviceController.Delete).Methods("DELETE")
	router.HandleFunc("/device/brand/{brand}", deviceController.GetByBrand).Methods("GET")

	return router
}
