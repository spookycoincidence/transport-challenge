package persistence_test

import (
	"testing"
	"transport-challenge/internal/infrastructure/persistence"

	"github.com/stretchr/testify/assert"
)

func TestRouteRepository_CreateRoute(t *testing.T) {
	// Se crea una base de datos ficticia
	db := persistence.NewDatabase()

	// Se crea un repositorio de rutas
	routeRepo := persistence.NewRouteRepository(db)

	// Se crea una nueva ruta
	id, err := routeRepo.CreateRoute("Route 1")

	// Nos aseguramos de que no haya errores
	assert.Nil(t, err)

	// Nos aseguramos de que el ID de la ruta sea mayor que 0 (porque es generado de manera secuencial)
	assert.Greater(t, id, 0)

	// Obtenemos todas las rutas y nos aseguramos de que la ruta "Route 1" esté presente
	routes, err := routeRepo.GetAllRoutes()
	assert.Nil(t, err)
	assert.Contains(t, routes, "Route 1")
}

func TestRouteRepository_CreateRoute_EmptyName(t *testing.T) {
	// Crear una base de datos ficticia
	db := persistence.NewDatabase()

	// Crear un repositorio de rutas con la base de datos simulada
	routeRepo := persistence.NewRouteRepository(db)

	// Intentar crear una ruta con un nombre vacío
	_, err := routeRepo.CreateRoute("")

	// Asegurarse de que se produzca un error debido a nombre vacío
	assert.NotNil(t, err)
	assert.Equal(t, "route name is empty", err.Error())
}
