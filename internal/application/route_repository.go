package application

import (
	"log"
	"transport-challenge/internal/domain"
)

type RouteRepository interface {
	Save(route domain.Route) error
	FindByID(id int) (domain.Route, error)
}

type MySQLRouteRepository struct {
	// Aca irían las configuraciones o conexiones reales de MySQL
}

func NewMySQLRouteRepository() *MySQLRouteRepository {
	return &MySQLRouteRepository{}
}

func (r *MySQLRouteRepository) Save(route domain.Route) error {
	log.Println("Saving route to database", route)
	return nil
}

func (r *MySQLRouteRepository) FindByID(id int) (domain.Route, error) {
	log.Println("Finding route by ID", id)
	return domain.Route{}, nil
}
