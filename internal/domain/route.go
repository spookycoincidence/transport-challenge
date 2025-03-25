package domain

import (
	"time"
)

// RouteStatus representa el estado actual de una ruta
type RouteStatus string

// Definición de estados posibles para una ruta
const (
	RouteStatusPending    RouteStatus = "PENDING"
	RouteStatusInProgress RouteStatus = "IN_PROGRESS"
	RouteStatusCompleted  RouteStatus = "COMPLETED"
	RouteStatusCancelled  RouteStatus = "CANCELLED"
)

// Route representa una ruta de distribución en el sistema de logística
type Route struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Vehicle   string      `json:"vehicle"`
	Driver    string      `json:"driver"`
	Status    RouteStatus `json:"status"`
	Purchases []Purchase  `json:"purchases"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// Purchase representa una compra asociada a una ruta
type Purchase struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// Validate realiza validaciones de negocio para una ruta
func (r *Route) Validate() error {
	if r.Name == "" {
		return ErrInvalidRouteName
	}

	if r.Vehicle == "" {
		return ErrInvalidVehicle
	}

	if r.Driver == "" {
		return ErrInvalidDriver
	}

	return nil
}

func (r *Route) validate() error {
	if r.Name == "" {
		return NewDomainError(
			ErrorCodes.ValidationError,
			"Route name is required",
			ErrInvalidRouteName,
		)
	}

	if r.Vehicle == "" {
		return NewDomainError(
			ErrorCodes.ValidationError,
			"Vehicle information is required",
			ErrInvalidVehicle,
		)
	}

	if r.Driver == "" {
		return NewDomainError(
			ErrorCodes.ValidationError,
			"Driver information is required",
			ErrInvalidDriver,
		)
	}

	return nil
}
