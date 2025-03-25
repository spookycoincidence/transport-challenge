package persistence

import (
	"errors"
	"sync"
)

type Database struct {
	mu  sync.Mutex
	ids map[string]int // Mapa ficticio para simular una tabla de secuencia por nombre de entidad
}

func NewDatabase() *Database {
	return &Database{
		ids: make(map[string]int),
	}
}

// GetNextID simula la obtención del siguiente ID para una tabla de secuencia.
func (db *Database) GetNextID(tableName string) (int, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Simular la creación de un nuevo ID para la "tabla" indicada por tableName
	if tableName == "" {
		return 0, errors.New("table name cannot be empty")
	}

	// Obtener el siguiente ID para esa "tabla"
	db.ids[tableName]++
	return db.ids[tableName], nil
}
