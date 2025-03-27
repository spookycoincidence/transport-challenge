package persistence_test

import (
	"testing"
	"transport-challenge/internal/infrastructure/persistence"

	"github.com/stretchr/testify/assert"
)

func TestRouteRepository_CreateRoute(t *testing.T) {

	db := persistence.NewDatabase()

	routeRepo := persistence.NewRouteRepository(db)

	id, err := routeRepo.CreateRoute("Route 1")

	assert.Nil(t, err)

	assert.Greater(t, id, 0)

	routes, err := routeRepo.GetAllRoutes()
	assert.Nil(t, err)
	assert.Contains(t, routes, "Route 1")
}

func TestRouteRepository_CreateRoute_EmptyName(t *testing.T) {

	db := persistence.NewDatabase()

	routeRepo := persistence.NewRouteRepository(db)

	_, err := routeRepo.CreateRoute("")

	assert.NotNil(t, err)
	assert.Equal(t, "route name is empty", err.Error())
}
