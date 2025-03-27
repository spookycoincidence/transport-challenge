package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"transport-challenge/internal/application"
	"transport-challenge/internal/domain"

	"github.com/gorilla/mux"
)

type Server struct {
	Router       *mux.Router
	RouteService *application.RouteService
}

func NewServer(routeService *application.RouteService) *Server {
	router := mux.NewRouter()

	server := &Server{
		Router:       router,
		RouteService: routeService,
	}

	// Define rutas
	server.routes()

	return server
}

func (s *Server) routes() {
	s.Router.HandleFunc("/routes", s.CreateRoute).Methods("POST")
	s.Router.HandleFunc("/routes", s.GetRoutes).Methods("GET")
	s.Router.HandleFunc("/routes/{id}", s.GetRouteByID).Methods("GET")
	s.Router.HandleFunc("/routes/{id}", s.UpdateRoute).Methods("PUT")
}

func (s *Server) CreateRoute(w http.ResponseWriter, r *http.Request) {
	var route domain.Route

	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := route.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := s.RouteService.CreateRoute(&route)
	if err != nil {
		http.Error(w, "Error creating route: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"id": id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s *Server) GetRoutes(w http.ResponseWriter, r *http.Request) {

	routes, err := s.RouteService.GetRoutesByStatus("")
	if err != nil {
		http.Error(w, "Error retrieving routes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(routes)
}

func (s *Server) GetRouteByID(w http.ResponseWriter, r *http.Request) {
	// Obtiene el ID de la ruta desde los parámetros de la URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid route ID", http.StatusBadRequest)
		return
	}

	// Busca la ruta por ID
	route, err := s.RouteService.GetRouteByID(id)
	if err != nil {
		if err == domain.ErrNotFound {
			http.Error(w, "Route not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error retrieving route: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(route)
}

func (s *Server) UpdateRoute(w http.ResponseWriter, r *http.Request) {
	// Obtiene el ID de la ruta desde los parámetros de la URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid route ID", http.StatusBadRequest)
		return
	}

	var route domain.Route
	err = json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := route.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.RouteService.UpdateRoute(id, &route)
	if err != nil {
		if err == domain.ErrNotFound {
			http.Error(w, "Route not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error updating route: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) Start() {
	log.Println("Iniciando servidor...")
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
