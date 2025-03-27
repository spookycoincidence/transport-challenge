package domain

import (
	"errors"
	"fmt"
)

// Errores generales
var (
	ErrNotFound = errors.New("record not found")
)

// Errores específicos de Ruta
var (
	ErrInvalidRouteName      = errors.New("route name is required and cannot be empty")
	ErrInvalidVehicle        = errors.New("vehicle information is required")
	ErrInvalidDriver         = errors.New("driver information is required")
	ErrRouteAlreadyCompleted = errors.New("route has already been completed")
	ErrInvalidRouteStatus    = errors.New("invalid route status")
)

// Errores específicos de Compra
var (
	ErrInvalidPurchaseID     = errors.New("purchase ID is invalid")
	ErrPurchaseAlreadyExists = errors.New("purchase already exists in route")
)

// DomainError error personalizado para errores de dominio
type DomainError struct {
	Code    string
	Message string
	Err     error
}

// Interfaz de error
func (e *DomainError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap permite que el error sea desenrollado
func (e *DomainError) Unwrap() error {
	return e.Err
}

// NewDomainError crea un nuevo error de dominio
func NewDomainError(code, message string, err error) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// ErrorCodes define códigos de error
var ErrorCodes = struct {
	NotFound            string
	ValidationError     string
	AlreadyExists       string
	InvalidState        string
	InternalServerError string
}{
	NotFound:            "NOT_FOUND",
	ValidationError:     "VALIDATION_ERROR",
	AlreadyExists:       "ALREADY_EXISTS",
	InvalidState:        "INVALID_STATE",
	InternalServerError: "INTERNAL_SERVER_ERROR",
}

// IsNotFoundError verifica si el error es de tipo "no encontrado"
func IsNotFoundError(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// IsValidationError verifica si el error es de validación
func IsValidationError(err error) bool {
	var domainErr *DomainError
	return errors.As(err, &domainErr) && domainErr.Code == ErrorCodes.ValidationError
}
