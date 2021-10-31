package handlers

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/devices/src/logger"
)

// When doesn't has an error
// Then should return false
func TestErrorHandlerNoError(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	errorHandler := NewErrorHandler(loggerInstance)
	want := false

	// act
	got := errorHandler.HandleError(nil, "Test step")

	// assert
	if got != want {
		t.Errorf("%v should be false %v", got, want)
	}
}

// When has an error
// Then should return true and log the error
func TestErrorHandlerWError(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := logger.NewLogger(buffer, logger.Info)
	errorHandler := NewErrorHandler(loggerInstance)
	expectedLog := "Level:Error Message:Test Error Step:Test step"
	want := true

	// act
	got := errorHandler.HandleError(errors.New("Test Error"), "Test step")

	// assert
	if got != want {
		t.Errorf("%v should be false %v", got, want)
	}

	messageLogged := buffer.String()

	if !strings.Contains(messageLogged, expectedLog) {
		t.Errorf("%v doesn't contains %v", messageLogged, expectedLog)
	}
}
