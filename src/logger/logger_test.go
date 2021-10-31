package logger

import (
	"bytes"
	"strings"
	"testing"
)

// When all values are provided
// Then should have message and step
func TestLogInfo(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := NewLogger(buffer, Info)
	want := "Level:Info Message:Test message Step:Test step}"

	// act
	loggerInstance.LogInfo("Test message", "Test step")

	// assert
	got := buffer.String()

	if !strings.Contains(got, want) {
		t.Errorf("%q doesn't contains %q", got, want)
	}
}

// When all values are empty
// Then should have empty message and step
func TestLogInfoEmpty(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := NewLogger(buffer, Info)
	want := "Level:Info Message: Step:}"

	// act
	loggerInstance.LogInfo("", "")

	// assert
	got := buffer.String()

	if !strings.Contains(got, want) {
		t.Errorf("%q doesn't contains %q", got, want)
	}
}

// When log level is higher that info
// Then should log nothing
func TestLogInfoHigherLogLevel(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := NewLogger(buffer, Warning)
	want := ""

	// act
	loggerInstance.LogInfo("", "")

	// assert
	got := buffer.String()

	if !strings.Contains(got, want) {
		t.Errorf("%q doesn't contains %q", got, want)
	}
}

// when all values are provided
// Then should have message and step
func TestLogWarning(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := NewLogger(buffer, Info)
	want := "Level:Warning Message:Test message Step:Test step}"

	// act
	loggerInstance.LogWarning("Test message", "Test step")

	// assert
	got := buffer.String()

	if !strings.Contains(got, want) {
		t.Errorf("%q doesn't contains %q", got, want)
	}
}

// When all values are empty
// Then should have empty message and step
func TestLogWarningEmpty(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := NewLogger(buffer, Warning)
	want := "Level:Warning Message: Step:}"

	// act
	loggerInstance.LogWarning("", "")

	// assert
	got := buffer.String()

	if !strings.Contains(got, want) {
		t.Errorf("%q doesn't contains %q", got, want)
	}
}

// When log level is higher that warning
// Then should log nothing
func TestLogWarningHigherLogLevel(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := NewLogger(buffer, Error)
	want := ""

	// act
	loggerInstance.LogWarning("Test message", "Test step")

	// assert
	got := buffer.String()

	if !strings.Contains(got, want) {
		t.Errorf("%q doesn't contains %q", got, want)
	}
}

// when all values are provided
// Then should have message and step
func TestLogError(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := NewLogger(buffer, Error)
	want := "Level:Error Message:Test message Step:Test step}"

	// act
	loggerInstance.LogError("Test message", "Test step")

	// assert
	got := buffer.String()

	if !strings.Contains(got, want) {
		t.Errorf("%q doesn't contains %q", got, want)
	}
}

// When all values are empty
// Then should have empty message and step
func TestLogErrorEmpty(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := NewLogger(buffer, Error)
	want := "Level:Error Message: Step:}"

	// act
	loggerInstance.LogError("", "")

	// assert
	got := buffer.String()

	if !strings.Contains(got, want) {
		t.Errorf("%q doesn't contains %q", got, want)
	}
}

// When log level is higher that error
// Then should log nothing
func TestLogErrorHigherLogLevel(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	loggerInstance := NewLogger(buffer, Disable)
	want := ""

	// act
	loggerInstance.LogError("", "")

	// assert
	got := buffer.String()

	if !strings.Contains(got, want) {
		t.Errorf("%q doesn't contains %q", got, want)
	}
}
