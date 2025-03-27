package persistence

import (
	"fmt"
	"sync"
	"transport-challenge/internal/domain"
)

type InMemoryRouteRepository struct {
	mu     sync.RWMutex
	routes map[int]domain.Route
	nextID int
}

func NewRouteRepository() *InMemoryRouteRepository {
	return &InMemoryRouteRepository{
		routes: make(map[int]domain.Route),
		nextID: 1,
	}
}

func (r *InMemoryRouteRepository) Create(route domain.Route) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Validar la ruta
	if err := route.Validate(); err != nil {
		return 0, err
	}

	// Asignar ID
	route.ID = r.nextID
	r.routes[r.nextID] = route

	// Incrementar el próximo ID
	r.nextID++

	return route.ID, nil
}

func (r *InMemoryRouteRepository) GetByID(id int) (domain.Route, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	route, exists := r.routes[id]
	if !exists {
		return domain.Route{}, domain.ErrNotFound
	}

	return route, nil
}

func (r *InMemoryRouteRepository) Update(id int, route domain.Route) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.routes[id]
	if !exists {
		return domain.ErrNotFound
	}

	if err := route.Validate(); err != nil {
		return err
	}

	r.routes[id] = route

	return nil
}

func (r *InMemoryRouteRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.routes[id]
	if !exists {
		return domain.ErrNotFound
	}

	delete(r.routes, id)
	return nil
}

func (r *InMemoryRouteRepository) List() ([]domain.Route, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	routes := make([]domain.Route, 0, len(r.routes))
	for _, route := range r.routes {
		routes = append(routes, route)
	}

	return routes, nil
}

func (r *InMemoryRouteRepository) FindByStatus(status domain.RouteStatus) ([]domain.Route, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var matchedRoutes []domain.Route
	for _, route := range r.routes {
		if route.Status == status {
			matchedRoutes = append(matchedRoutes, route)
		}
	}

	return matchedRoutes, nil
}

func (r *InMemoryRouteRepository) AssignPurchaseToRoute(routeID int, purchase domain.Purchase) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	route, exists := r.routes[routeID]
	if !exists {
		return fmt.Errorf("route with ID %d not found", routeID)
	}

	for _, existingPurchase := range route.Purchases {
		if existingPurchase.ID == purchase.ID {
			return fmt.Errorf("purchase with ID %d already exists in route", purchase.ID)
		}
	}

	route.Purchases = append(route.Purchases, purchase)
	r.routes[routeID] = route

	return nil
}

var _ domain.RouteRepository = &InMemoryRouteRepository{}
