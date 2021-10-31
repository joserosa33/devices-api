package handlers

import (
	"github.com/devices/src/logger"
)

type errorHandler struct {
	logger logger.Logger
}

type ErrorHandler interface {
	HandleError(err error, step string) bool
}

func NewErrorHandler(logger logger.Logger) ErrorHandler {
	return &errorHandler{
		logger: logger,
	}
}

func (handler *errorHandler) HandleError(err error, step string) bool {
	if err != nil {
		handler.logger.LogError(err.Error(), step)
		return true
	}

	return false
}
