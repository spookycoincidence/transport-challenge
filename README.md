# Challenge de Transporte 🚚

Este repositorio contiene la solución para el desafío técnico de gestión de rutas y distribución de productos en un sistema de microservicios. La aplicación REST desarrollada permite gestionar rutas de distribución, asignar compras a rutas y monitorear el estado de las compras.

## Descripción del Desafío  💡 

La empresa tiene un sistema de microservicios que gestiona la compra de productos y su posterior entrega mediante la generación de rutas y distribución en vehículos. Durante la distribución, el sistema monitorea el estado de las rutas y las compras asociadas, enviando notificaciones cuando las compras son despachadas o entregadas.


## Requisitos del Desafío  📝 

- **Crear una nueva ruta**: Con información del vehículo y conductor asignados.
- **Agregar compras a una ruta**: Permitir asignar compras a una ruta de distribución sin restricciones específicas de carga.
- **Consultar rutas**: Obtener detalles de una ruta, incluyendo las compras asociadas y su estado.

## Tecnologías Utilizadas  🛠️ 

- **Go**: Lenguaje de programación principal.
- **REST API**: Exposición de servicios para gestionar rutas y compras.
- **Testify**: Librería para realizar tests unitarios y de integración.
- **Gorilla Mux**: Librería para la definición y manejo de rutas en la API.
- **Postman**: Herramienta utilizada para realizar las pruebas manuales de la API.
- **MySQL**: Implementación de un repositorio ficticio para simular la interacción con una base de datos MySQL.

## Instalación ⚙️ 

1. Clonar el repositorio:
   ```bash
   git clone https://github.com/spookycoincidence/transport-challenge.git
   cd transport-challenge
   ```

2. Ejecutar la aplicación:
   ```bash
   go run main.go
   ```
   La API estará disponible en `http://localhost:8080`

## Endpoints de la API 🔧

### Crear Nueva Ruta
- **Endpoint**: `POST /routes`
- **Cuerpo**: Información de vehículo y conductor 🚗
- **Respuesta**: Detalles de la ruta creada

### Asignar Compra a Ruta
- **Endpoint**: `POST /routes/{route_id}/purchases`
- **Cuerpo**: Información de compra para asignar a la ruta 📦
- **Respuesta**: Confirmación de asignación de compra

### Obtener Todas las Rutas
- **Endpoint**: `GET /routes`
- **Respuesta**: Lista de todas las rutas con detalles

### Obtener Ruta Específica
- **Endpoint**: `GET /routes/{id}`
- **Parámetros**: ID de Ruta  🔑
- **Respuesta**: Detalles de ruta, incluyendo compras asociadas y su estado

### Actualizar Ruta
- **Endpoint**: `PUT /routes/{id}`
- **Parámetros**: ID de Ruta
- **Cuerpo**: Información actualizada de ruta  🔄
- **Respuesta**: Detalles de ruta actualizados

## Ejemplos en Postman 🖥️

### Crear una Ruta
- **URL**: `http://localhost:8080/routes`
- **Método**: POST
- **Cuerpo**:
  ```json
  {
    "vehicle": "ABC-123",
    "driver": "Julián"
  }
  ```

### Obtener Ruta Específica
- **URL**: `http://localhost:8080/routes/1`
- **Método**: GET

## Arquitectura del Sistema 🏗️

La aplicación sigue una arquitectura de microservicios con los siguientes componentes principales:

- **Microservicio de Rutas**: Gestiona creación de rutas y asignación de compras.  🛣️
- **Persistencia en memoria**: Los datos se mantienen en memoria, simulando la interacción con una base de datos MySQL. 🧠
- **API REST**: Expuesta utilizando Gorilla Mux para gestionar las rutas y las solicitudes de la API.

### Módulos Principales

- **Aplication**: Lógica de negocio para creación de rutas y asignación de compras.
- **Domain**: Modelos de datos y reglas de negocio.
- **Infraestructure**: Capa de persistencia en memoria e interfaz de servicio HTTP.  🔌






   
