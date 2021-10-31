package services

import (
	"bytes"
	"errors"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/devices/src/handlers"
	"github.com/devices/src/logger"
	"github.com/devices/src/models"
)

// When all goes well
// Then should return expect object
func TestGetByIdWithoutError(t *testing.T) {
	// arrange
	db, mock, _ := sqlmock.New()
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	errorHandler := handlers.NewErrorHandler(loggerInstance)
	deviceService := NewDeviceService(db, loggerInstance, errorHandler)

	expectedDevice := models.Device{Id: 0, Name: "testName", Brand: "testBrand", Created: "testCreated"}

	deviceMockRow := mock.NewRows([]string{"id", "name", "brand", "created"}).AddRow("0", "testName", "testBrand", "testCreated")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, brand, created FROM devices where id = $1`)).WithArgs(0).WillReturnRows(deviceMockRow)

	// act
	got := deviceService.GetById(0)

	// assert
	if got != expectedDevice {
		t.Errorf("%v obtained doesn't match expected %v", got, expectedDevice)
	}
}

// When occurs some error
// Then should return empty object and log error
func TestGetByIdWithError(t *testing.T) {
	// arrange
	db, mock, _ := sqlmock.New()
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	errorHandler := handlers.NewErrorHandler(loggerInstance)
	deviceService := NewDeviceService(db, loggerInstance, errorHandler)

	expectedDevice := models.Device{}
	expectedError := "Test Error"

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, brand, created FROM devices where id = $1`)).WithArgs(0).WillReturnError(errors.New(expectedError))

	// act
	got := deviceService.GetById(0)

	// assert
	if got != expectedDevice {
		t.Errorf("%v obtained doesn't match expected %v", got, expectedDevice)
	}

	messageLogged := buffer.String()

	if !strings.Contains(messageLogged, expectedError) {
		t.Errorf("%v doesn't contains %v", messageLogged, expectedError)
	}
}

// When all goes well
// Then should return id of added object
func TestAddWithoutError(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	db, mock, _ := sqlmock.New()
	errorHandler := handlers.NewErrorHandler(loggerInstance)
	deviceService := NewDeviceService(db, loggerInstance, errorHandler)

	deviceToAdd := models.Device{Name: "testName", Brand: "testBrand", Created: "testCreated"}
	expectedId := 0

	idMockRow := mock.NewRows([]string{"id"}).AddRow("0")
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO devices(name, brand, created) VALUES ($1, $2, $3) returning id;`)).
		WithArgs(deviceToAdd.Name, deviceToAdd.Brand, deviceToAdd.Created).
		WillReturnRows(idMockRow)

	// act
	got := deviceService.Add(deviceToAdd)

	// assert
	if got != expectedId {
		t.Errorf("%v obtained doesn't match expected %v", got, expectedId)
	}
}

// When occurs some error
// Then should return -1 and log the error
func TestAddWithError(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	db, mock, _ := sqlmock.New()
	errorHandler := handlers.NewErrorHandler(loggerInstance)
	deviceService := NewDeviceService(db, loggerInstance, errorHandler)

	deviceToAdd := models.Device{Name: "testName", Brand: "testBrand", Created: "testCreated"}
	expectedId := -1
	expectedError := "Test Error"

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO devices(name, brand, created) VALUES ($1, $2, $3) returning id;`)).
		WithArgs(deviceToAdd.Name, deviceToAdd.Brand, deviceToAdd.Created).
		WillReturnError(errors.New(expectedError))

	// act
	got := deviceService.Add(deviceToAdd)

	// assert
	if got != expectedId {
		t.Errorf("%v obtained doesn't match expected %v", got, expectedId)
	}

	messageLogged := buffer.String()

	if !strings.Contains(messageLogged, expectedError) {
		t.Errorf("%v doesn't contains %v", messageLogged, expectedError)
	}
}

// When all goes well
// Then should return array of devices
func TestGetAllWithoutError(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	db, mock, _ := sqlmock.New()
	errorHandler := handlers.NewErrorHandler(loggerInstance)
	deviceService := NewDeviceService(db, loggerInstance, errorHandler)

	expectedFirstDevice := models.Device{Id: 1, Name: "testName1", Brand: "testBrand1", Created: "testCreated1"}

	mockRows := mock.NewRows([]string{"id", "name", "brand", "created"}).
		AddRow("1", "testName1", "testBrand1", "testCreated1").
		AddRow("2", "testName2", "testBrand2", "testCreated2")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM device`)).
		WillReturnRows(mockRows)

	// act
	got := deviceService.GetAll()

	// assert
	if got[0] != expectedFirstDevice {
		t.Errorf("%v obtained doesn't match expected %v", got, expectedFirstDevice)
	}

	if len(got) != 2 {
		t.Errorf("expected len of 2 actual %v", len(got))
	}
}

// When occurs some error
// Then should return empty devices array and log the error
func TestGetAllWithError(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	db, mock, _ := sqlmock.New()
	errorHandler := handlers.NewErrorHandler(loggerInstance)
	deviceService := NewDeviceService(db, loggerInstance, errorHandler)

	expectedError := "Test Error"

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM device`)).WillReturnError(errors.New(expectedError))

	// act
	got := deviceService.GetAll()

	// assert
	if len(got) != 0 {
		t.Errorf("expected len of 0 actual %v", len(got))
	}

	messageLogged := buffer.String()

	if !strings.Contains(messageLogged, expectedError) {
		t.Errorf("%v doesn't contains %v", messageLogged, expectedError)
	}
}

// When all goes well
// Then should return id of updated object
func TestUpdateWithoutError(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	db, mock, _ := sqlmock.New()
	errorHandler := handlers.NewErrorHandler(loggerInstance)
	deviceService := NewDeviceService(db, loggerInstance, errorHandler)

	deviceToUpdate := models.Device{Id: 0, Name: "testName", Brand: "testBrand", Created: "testCreated"}
	expectedId := 0

	idMockRow := mock.NewRows([]string{"id"}).AddRow("0")
	mock.ExpectQuery(regexp.QuoteMeta(`UPDATE devices SET name=$1, brand=$2, created=$3 WHERE id=$4 returnin id`)).
		WithArgs(deviceToUpdate.Name, deviceToUpdate.Brand, deviceToUpdate.Created, deviceToUpdate.Id).
		WillReturnRows(idMockRow)

	// act
	got := deviceService.Update(deviceToUpdate)

	// assert
	if got != expectedId {
		t.Errorf("%v obtained doesn't match expected %v", got, expectedId)
	}
}

// When occurs some error
// Then should return -1 and log the error
func TestUpdateWithError(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	db, mock, _ := sqlmock.New()
	errorHandler := handlers.NewErrorHandler(loggerInstance)
	deviceService := NewDeviceService(db, loggerInstance, errorHandler)

	deviceToUpdate := models.Device{Id: 0, Name: "testName", Brand: "testBrand", Created: "testCreated"}
	expectedId := -1
	expectedError := "Test Error"

	mock.ExpectQuery(regexp.QuoteMeta(`UPDATE devices SET name=$1, brand=$2, created=$3 WHERE id=$4 returnin id`)).
		WithArgs(deviceToUpdate.Name, deviceToUpdate.Brand, deviceToUpdate.Created, deviceToUpdate.Id).
		WillReturnError(errors.New(expectedError))

	// act
	got := deviceService.Update(deviceToUpdate)

	// assert
	if got != expectedId {
		t.Errorf("%v obtained doesn't match expected %v", got, expectedId)
	}

	messageLogged := buffer.String()

	if !strings.Contains(messageLogged, expectedError) {
		t.Errorf("%v doesn't contains %v", messageLogged, expectedError)
	}
}

// When all goes well
// Then should return id of deleted object
func TestDeleteWithoutError(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	db, mock, _ := sqlmock.New()
	errorHandler := handlers.NewErrorHandler(loggerInstance)
	deviceService := NewDeviceService(db, loggerInstance, errorHandler)

	idToDelete := 0

	mockRow := mock.NewRows([]string{"id"}).AddRow("0")
	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM devices where id = $1 returning id`)).
		WithArgs(idToDelete).
		WillReturnRows(mockRow)

	// act
	got := deviceService.DeleteById(idToDelete)

	// assert
	if got != idToDelete {
		t.Errorf("%v obtained doesn't match expected %v", got, idToDelete)
	}
}

// When occurs some error
// Then should return -1 and log the error
func TestDeleteWithError(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	db, mock, _ := sqlmock.New()
	errorHandler := handlers.NewErrorHandler(loggerInstance)
	deviceService := NewDeviceService(db, loggerInstance, errorHandler)

	idToDelete := 0
	expectedId := -1
	expectedError := "Test Error"

	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM devices where id = $1 returning id`)).
		WithArgs(idToDelete).
		WillReturnError(errors.New(expectedError))

	// act
	got := deviceService.DeleteById(idToDelete)

	// assert
	if got != expectedId {
		t.Errorf("%v obtained doesn't match expected %v", got, expectedId)
	}

	messageLogged := buffer.String()

	if !strings.Contains(messageLogged, expectedError) {
		t.Errorf("%v doesn't contains %v", messageLogged, expectedError)
	}
}

// When all goes well
// Then should return array of devices
func TestGetByBrandWithoutError(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	db, mock, _ := sqlmock.New()
	errorHandler := handlers.NewErrorHandler(loggerInstance)
	deviceService := NewDeviceService(db, loggerInstance, errorHandler)

	expectedFirstDevice := models.Device{Id: 1, Name: "testName1", Brand: "testBrand", Created: "testCreated1"}

	mockRows := mock.NewRows([]string{"id", "name", "brand", "created"}).
		AddRow("1", "testName1", "testBrand", "testCreated1").
		AddRow("2", "testName2", "testBrand", "testCreated2")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM devices where brand = $1`)).
		WithArgs(expectedFirstDevice.Brand).
		WillReturnRows(mockRows)

	// act
	got := deviceService.GetByBrand(expectedFirstDevice.Brand)

	// assert
	if got[0] != expectedFirstDevice {
		t.Errorf("%v obtained doesn't match expected %v", got, expectedFirstDevice)
	}

	if len(got) != 2 {
		t.Errorf("expected len of 2 actual %v", len(got))
	}
}

// When occurs some error
// Then should return empty devices array and log the error
func TestGetByBrandWithError(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	db, mock, _ := sqlmock.New()
	errorHandler := handlers.NewErrorHandler(loggerInstance)
	deviceService := NewDeviceService(db, loggerInstance, errorHandler)

	expectedError := "Test Error"

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM devices where brand = $1`)).WillReturnError(errors.New(expectedError))

	// act
	got := deviceService.GetByBrand("Test Brand")

	// assert
	if len(got) != 0 {
		t.Errorf("expected len of 0 actual %v", len(got))
	}

	messageLogged := buffer.String()

	if !strings.Contains(messageLogged, expectedError) {
		t.Errorf("%v doesn't contains %v", messageLogged, expectedError)
	}
}
