package main

import (
	"fmt"
	"transport-challenge/internal/infrastructure/persistence"
)

func main() {
	// Crear una instancia de la base de datos simulada
	db := persistence.NewDatabase()

	// Crear una nueva instancia del repositorio de rutas pasando la base de datos
	routeRepo := persistence.NewRouteRepository(db)

	// Crear una nueva ruta
	id, err := routeRepo.CreateRoute("Route 1")
	if err != nil {
		fmt.Println("Error creating route:", err)
		return
	}
	fmt.Printf("Created route with ID: %d\n", id)

	// Obtener todas las rutas
	routes, err := routeRepo.GetAllRoutes()
	if err != nil {
		fmt.Println("Error getting routes:", err)
		return
	}
	fmt.Println("All Routes:", routes)
}
