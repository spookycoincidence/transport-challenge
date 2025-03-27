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

type Payload struct {
	IDCompra     int    `json:"id_compra"`
	Tipo         string `json:"tipo"`
	Descripcion  string `json:"descripcion"`
	Destinatario string `json:"destinatario"`
}

func NuevoServicioPush(config ConfigPush) *ServicioPush {
	return &ServicioPush{
		config: config,
	}
}

func (s *ServicioPush) Enviar(notificacion Notificacion) error {

	if s.config.ServidorAPI == "" || s.config.ClaveAPI == "" {
		return fmt.Errorf("configuración de push incompleta")
	}

	payload := Payload{
		IDCompra:     notificacion.IDCompra,
		Tipo:         string(notificacion.Tipo),
		Descripcion:  notificacion.Descripcion,
		Destinatario: notificacion.Destinatario,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error al serializar payload: %w", err)
	}

	req, err := http.NewRequest("POST", s.config.ServidorAPI, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error al crear solicitud HTTP: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.config.ClaveAPI))

	cliente := &http.Client{}

	resp, err := cliente.Do(req)
	if err != nil {
		return fmt.Errorf("error al enviar notificación push: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error en respuesta push: código de estado %d", resp.StatusCode)
	}

	return nil
}
