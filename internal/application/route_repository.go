package persistence

import (
	"log"
	"transport-challenge/internal/domain"
)

type RouteRepository interface {
	Save(route domain.Route) error
	FindByID(id int) (domain.Route, error)
}

type MySQLRouteRepository struct {
	// Aquí irían las configuraciones o conexiones de MySQL
}

func NewMySQLRouteRepository() *MySQLRouteRepository {
	return &MySQLRouteRepository{}
}

func (r *MySQLRouteRepository) Save(route domain.Route) error {
	// Simulación de la lógica de persistencia
	log.Println("Saving route to database", route)
	return nil
}

func (r *MySQLRouteRepository) FindByID(id int) (domain.Route, error) {
	// Simulación de la búsqueda de una ruta
	log.Println("Finding route by ID", id)
	return domain.Route{}, nil
}
