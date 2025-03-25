package application

import (
	"fmt"
	"time"
	"transport-challenge/internal/domain"
)

type RouteService struct {
	routeRepo domain.RouteRepository
}

func NewRouteService(repo domain.RouteRepository) *RouteService {
	return &RouteService{
		routeRepo: repo,
	}
}

// CreateRoute crea una nueva ruta con validaciones de negocio
func (s *RouteService) CreateRoute(route *domain.Route) (int, error) {
	// Validaciones de negocio
	if err := route.Validate(); err != nil {
		return 0, err
	}

	// Establecer valores por defecto
	route.Status = domain.RouteStatusPending
	route.CreatedAt = time.Now()
	route.UpdatedAt = time.Now()

	// Crear ruta
	id, err := s.routeRepo.Create(*route)
	if err != nil {
		return 0, fmt.Errorf("failed to create route: %w", err)
	}

	return id, nil
}

// GetRouteByID recupera una ruta por su ID
func (s *RouteService) GetRouteByID(id int) (domain.Route, error) {
	route, err := s.routeRepo.GetByID(id)
	if err != nil {
		return domain.Route{}, fmt.Errorf("failed to retrieve route: %w", err)
	}

	return route, nil
}

// UpdateRoute actualiza una ruta existente
func (s *RouteService) UpdateRoute(id int, route *domain.Route) error {
	// Validar la ruta
	if err := route.Validate(); err != nil {
		return err
	}

	// Verificar si la ruta existe
	existingRoute, err := s.routeRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("route not found: %w", err)
	}

	// Actualizar campos modificables
	existingRoute.Name = route.Name
	existingRoute.Vehicle = route.Vehicle
	existingRoute.Driver = route.Driver
	existingRoute.Status = route.Status
	existingRoute.UpdatedAt = time.Now()

	// Persistir cambios
	if err := s.routeRepo.Update(id, existingRoute); err != nil {
		return fmt.Errorf("failed to update route: %w", err)
	}

	return nil
}

// Métodos existentes previamente...

// GetRoutesByStatus recupera rutas por su estado
func (s *RouteService) GetRoutesByStatus(status domain.RouteStatus) ([]domain.Route, error) {
	// Si no se proporciona un estado, devolver todas las rutas
	if status == "" {
		return s.routeRepo.List()
	}

	routes, err := s.routeRepo.FindByStatus(status)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve routes: %w", err)
	}

	return routes, nil
}

// AssignPurchaseToRoute asigna una compra a una ruta
func (s *RouteService) AssignPurchaseToRoute(routeID int, purchase domain.Purchase) error {
	// Recuperar la ruta existente
	route, err := s.routeRepo.GetByID(routeID)
	if err != nil {
		return fmt.Errorf("route not found: %w", err)
	}

	// Validar que la ruta esté en estado adecuado
	if route.Status != domain.RouteStatusPending && route.Status != domain.RouteStatusInProgress {
		return fmt.Errorf("cannot assign purchase to route with status %s", route.Status)
	}

	// Asignar compra a la ruta
	err = s.routeRepo.AssignPurchaseToRoute(routeID, purchase)
	if err != nil {
		return fmt.Errorf("failed to assign purchase to route: %w", err)
	}

	// Actualizar estado de la ruta si es necesario
	if route.Status == domain.RouteStatusPending {
		route.Status = domain.RouteStatusInProgress
		route.UpdatedAt = time.Now()

		if err := s.routeRepo.Update(routeID, route); err != nil {
			return fmt.Errorf("failed to update route status: %w", err)
		}
	}

	return nil
}

// CompleteRoute marca una ruta como completada
func (s *RouteService) CompleteRoute(routeID int) error {
	// Recuperar la ruta
	route, err := s.routeRepo.GetByID(routeID)
	if err != nil {
		return fmt.Errorf("route not found: %w", err)
	}

	// Validar que todos los paquetes estén entregados
	allDelivered := true
	for _, purchase := range route.Purchases {
		if purchase.Status != "DELIVERED" {
			allDelivered = false
			break
		}
	}

	if !allDelivered {
		return fmt.Errorf("cannot complete route: not all purchases are delivered")
	}

	// Actualizar estado de la ruta
	route.Status = domain.RouteStatusCompleted
	route.UpdatedAt = time.Now()

	if err := s.routeRepo.Update(routeID, route); err != nil {
		return fmt.Errorf("failed to complete route: %w", err)
	}

	return nil
}
