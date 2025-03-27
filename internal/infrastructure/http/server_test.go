package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"transport-challenge/internal/domain"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// MockRouteRepository simula un repositorio de rutas para pruebas
type MockRouteRepository struct {
	routes map[int]domain.Route
}

func NewMockRouteRepository() *MockRouteRepository {
	return &MockRouteRepository{
		routes: make(map[int]domain.Route),
	}
}

func (m *MockRouteRepository) Create(route domain.Route) (int, error) {
	route.ID = len(m.routes) + 1
	m.routes[route.ID] = route
	return route.ID, nil
}

func (m *MockRouteRepository) GetByID(id int) (domain.Route, error) {
	route, exists := m.routes[id]
	if !exists {
		return domain.Route{}, domain.ErrNotFound
	}
	return route, nil
}

func (m *MockRouteRepository) Update(id int, route domain.Route) error {
	if _, exists := m.routes[id]; !exists {
		return domain.ErrNotFound
	}
	route.ID = id
	m.routes[id] = route
	return nil
}

func (m *MockRouteRepository) Delete(id int) error {
	delete(m.routes, id)
	return nil
}

func (m *MockRouteRepository) List() ([]domain.Route, error) {
	routes := make([]domain.Route, 0, len(m.routes))
	for _, route := range m.routes {
		routes = append(routes, route)
	}
	return routes, nil
}

func (m *MockRouteRepository) FindByStatus(status domain.RouteStatus) ([]domain.Route, error) {
	var matchedRoutes []domain.Route
	for _, route := range m.routes {
		if route.Status == status {
			matchedRoutes = append(matchedRoutes, route)
		}
	}
	return matchedRoutes, nil
}

func (m *MockRouteRepository) AssignPurchaseToRoute(routeID int, purchase domain.Purchase) error {
	return nil
}

func TestCreateRoute(t *testing.T) {
	// Crea repositorio mock y servidor
	mockRepo := NewMockRouteRepository()
	server := NewServer(mockRepo)

	routeJSON := `{
		"name": "Test Route",
		"vehicle": "Truck",
		"driver": "Julian",
		"status": "PENDING"
	}`

	req, err := http.NewRequest("POST", "/routes", bytes.NewBufferString(routeJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	server.Router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "id")
	assert.Greater(t, int(response["id"].(float64)), 0)
}

func TestGetRouteByID(t *testing.T) {
	mockRepo := NewMockRouteRepository()

	testRoute := domain.Route{
		Name:    "Test Route",
		Vehicle: "Van",
		Driver:  "Ramona",
		Status:  domain.RouteStatusPending,
	}
	routeID, _ := mockRepo.Create(testRoute)

	server := NewServer(mockRepo)

	req, err := http.NewRequest("GET", "/routes/"+mux.Var{Key: "id", Value: "1"}, nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	server.Router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var route domain.Route
	err = json.Unmarshal(recorder.Body.Bytes(), &route)
	assert.NoError(t, err)
	assert.Equal(t, routeID, route.ID)
	assert.Equal(t, "Test Route", route.Name)
}

func TestUpdateRoute(t *testing.T) {
	mockRepo := NewMockRouteRepository()

	testRoute := domain.Route{
		Name:    "Original Route",
		Vehicle: "Truck",
		Driver:  "Julian",
		Status:  domain.RouteStatusPending,
	}
	routeID, _ := mockRepo.Create(testRoute)

	server := NewServer(mockRepo)

	updatedRouteJSON := `{
		"name": "Updated Route",
		"vehicle": "Bus",
		"driver": "Ramona",
		"status": "IN_PROGRESS"
	}`

	req, err := http.NewRequest("PUT", "/routes/"+mux.Var{Key: "id", Value: "1"}, bytes.NewBufferString(updatedRouteJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	server.Router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	updatedRoute, err := mockRepo.GetByID(routeID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Route", updatedRoute.Name)
	assert.Equal(t, "Bus", updatedRoute.Vehicle)
	assert.Equal(t, domain.RouteStatusInProgress, updatedRoute.Status)
}

func TestCreateRouteValidationError(t *testing.T) {
	mockRepo := NewMockRouteRepository()
	server := NewServer(mockRepo)

	routeJSON := `{
		"vehicle": "Truck",
		"driver": "Julian",
		"status": "PENDING"
	}`

	req, err := http.NewRequest("POST", "/routes", bytes.NewBufferString(routeJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	server.Router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
