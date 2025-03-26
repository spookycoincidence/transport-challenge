package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// ServicioPush maneja el envío de notificaciones push
type ServicioPush struct {
	config ConfigPush
}

// Payload representa la estructura de una notificación push
type Payload struct {
	IDCompra     int    `json:"id_compra"`
	Tipo         string `json:"tipo"`
	Descripcion  string `json:"descripcion"`
	Destinatario string `json:"destinatario"`
}

// NuevoServicioPush crea una nueva instancia del servicio de notificaciones push
func NuevoServicioPush(config ConfigPush) *ServicioPush {
	return &ServicioPush{
		config: config,
	}
}

// Enviar envía una notificación push
func (s *ServicioPush) Enviar(notificacion Notificacion) error {
	// Validar configuración
	if s.config.ServidorAPI == "" || s.config.ClaveAPI == "" {
		return fmt.Errorf("configuración de push incompleta")
	}

	// Crear payload
	payload := Payload{
		IDCompra:     notificacion.IDCompra,
		Tipo:         string(notificacion.Tipo),
		Descripcion:  notificacion.Descripcion,
		Destinatario: notificacion.Destinatario,
	}

	// Convertir payload a JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error al serializar payload: %w", err)
	}

	// Crear solicitud HTTP
	req, err := http.NewRequest("POST", s.config.ServidorAPI, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error al crear solicitud HTTP: %w", err)
	}

	// Configurar headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.config.ClaveAPI))

	// Crear cliente HTTP
	cliente := &http.Client{}

	// Enviar solicitud
	resp, err := cliente.Do(req)
	if err != nil {
		return fmt.Errorf("error al enviar notificación push: %w", err)
	}
	defer resp.Body.Close()

	// Verificar respuesta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error en respuesta push: código de estado %d", resp.StatusCode)
	}

	return nil
}
