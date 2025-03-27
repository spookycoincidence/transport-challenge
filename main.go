package main

import (
	"fmt"
	"transport-challenge/internal/infrastructure/persistence"
)

func main() {

	db := persistence.NewDatabase()

	routeRepo := persistence.NewRouteRepository

	id, err := routeRepo.CreateRoute("Route 1")
	if err != nil {
		fmt.Println("Error creating route:", err)
		return
	}
	fmt.Printf("Created route with ID: %d\n", id)

	routes, err := routeRepo.GetAllRoutes()
	if err != nil {
		fmt.Println("Error getting routes:", err)
		return
	}
	fmt.Println("All Routes:", routes)
}
