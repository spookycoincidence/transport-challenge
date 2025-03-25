package domain

type Repository[T any] interface {
	// Create agrega un nuevo elemento
	Create(item T) (int, error)

	// GetByID recupera un elemento por su identificador
	GetByID(id int) (T, error)

	// Update modifica un elemento existente
	Update(id int, item T) error

	// Delete elimina un elemento por su identificador
	Delete(id int) error

	// List recupera todos los elementos
	List() ([]T, error)
}

type RouteRepository interface {
	Repository[Route]

	FindByStatus(status RouteStatus) ([]Route, error)
	AssignPurchaseToRoute(routeID int, purchase Purchase) error
}
