package services

import (
	"database/sql"
	"fmt"

	"github.com/devices/src/handlers"
	"github.com/devices/src/logger"
	"github.com/devices/src/models"
)

type deviceService struct {
	database     *sql.DB
	logger       logger.Logger
	errorHandler handlers.ErrorHandler
}

type DeviceService interface {
	Add(device models.Device) int
	GetAll() []models.Device
	GetById(deviceId int) models.Device
	Update(device models.Device) int
	DeleteById(deviceId int) int
	GetByBrand(brand string) []models.Device
}

func NewDeviceService(sql *sql.DB, logger logger.Logger, errorHandler handlers.ErrorHandler) DeviceService {
	return &deviceService{
		database:     sql,
		logger:       logger,
		errorHandler: errorHandler,
	}
}

func (d *deviceService) GetById(deviceId int) models.Device {
	var row = d.database.QueryRow("SELECT id, name, brand, created FROM devices where id = $1", deviceId)

	var id int
	var name string
	var brand string
	var created string

	err := row.Scan(&id, &name, &brand, &created)

	if d.errorHandler.HandleError(err, "DeviceService/GetById") {
		return models.Device{}
	}

	return models.Device{Id: id, Name: name, Brand: brand, Created: created}
}

func (d *deviceService) Add(device models.Device) int {
	var id int

	row := d.database.QueryRow(
		"INSERT INTO devices(name, brand, created) VALUES ($1, $2, $3) returning id;",
		device.Name, device.Brand, device.Created)

	err := row.Scan(&id)

	if d.errorHandler.HandleError(err, "DeviceService/Add") {
		return -1
	}

	d.logger.LogInfo("Added to database new item wiht id: "+fmt.Sprint(id), "DeviceService/Add")
	return id
}

func (d *deviceService) GetAll() []models.Device {
	var devices []models.Device

	rows, err := d.database.Query("SELECT * FROM devices")

	if d.errorHandler.HandleError(err, "DeviceService/GetAll") {
		return devices
	}

	for rows.Next() {
		var id int
		var name string
		var brand string
		var created string

		err := rows.Scan(&id, &name, &brand, &created)

		d.errorHandler.HandleError(err, "DeviceService/GetAll/Next")

		devices = append(devices, models.Device{Id: id, Name: name, Brand: brand, Created: created})

	}

	d.logger.LogInfo("Get all items", "DeviceService/GetAll")

	return devices
}

func (d *deviceService) Update(device models.Device) int {
	var id int

	row := d.database.QueryRow(
		"UPDATE devices SET name=$1, brand=$2, created=$3 WHERE id=$4 returning id",
		device.Name, device.Brand, device.Created, device.Id)

	err := row.Scan(&id)

	if d.errorHandler.HandleError(err, "DeviceService/Update") {
		return -1
	}

	d.logger.LogInfo("Updated item wiht id: "+fmt.Sprint(id), "DeviceService/Update")
	return id
}

func (d *deviceService) DeleteById(deviceId int) int {
	var id int
	row := d.database.QueryRow("DELETE FROM devices where id = $1 returning id", deviceId)

	err := row.Scan(&id)

	if d.errorHandler.HandleError(err, "DeviceService/DeleteById") {
		return -1
	}

	d.logger.LogInfo("Deleted item wiht id: "+fmt.Sprint(id), "DeviceService/DeleteById")
	return id
}

func (d *deviceService) GetByBrand(brand string) []models.Device {
	var devices []models.Device

	rows, err := d.database.Query("SELECT * FROM devices where brand = $1", brand)

	if d.errorHandler.HandleError(err, "DeviceService/SearchByBrand") {
		return devices
	}

	for rows.Next() {
		var id int
		var name string
		var brand string
		var created string

		err := rows.Scan(&id, &name, &brand, &created)

		d.errorHandler.HandleError(err, "DeviceService/SearchByBrand/Next")

		devices = append(devices, models.Device{Id: id, Name: name, Brand: brand, Created: created})
	}

	d.logger.LogInfo("Get items wiht brand: "+brand, "DeviceService/SearchByBrand")

	return devices
}
