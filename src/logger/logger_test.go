package logger

import (
	"bytes"
	"strings"
	"testing"
)

// When_AllValuesProvided_And_TestLogInfo_Then_ShouldHaveTestMessageAndStep
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

// When_AllValuesAreEmpty_And_TestLogInfo_Then_ShouldHaveEmptyTestMessageAndStep
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

// When_LogelLevelIsHigherThanInfo_And_TestLogInfo_Then_ShouldLogNothing
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

// When_AllValuesProvided_And_TestLogWarning_Then_ShouldHaveTestMessageAndStep
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

// When_AllValuesAreEmpty_And_TestLogWarning_Then_ShouldHaveEmptyTestMessageAndStep
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

// When_LogLevelHigherThanWarning_And_TestLogWarning_Then_ShouldLogNothing
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

// When_AllValuesAreEmpty_And_TestLogError_Then_ShouldHaveTestMessageAndStep
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

// When_AllValuesAreEmpty_And_TestLogError_Then_ShouldHaveEmptyTestMessageAndStep
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

// When_LogelLevelIsHigherThanError_And_TestLogError_Then_ShouldLogNothing
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
