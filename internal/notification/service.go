package notification

import (
	"fmt"
	"log"
)

// TipoNotificacion representa los diferentes tipos de notificaciones
type TipoNotificacion string

const (
	NotificacionCompraCreada    TipoNotificacion = "COMPRA_CREADA"
	NotificacionCompraEnRuta    TipoNotificacion = "COMPRA_EN_RUTA"
	NotificacionCompraEntregada TipoNotificacion = "COMPRA_ENTREGADA"
	NotificacionCompraEnError   TipoNotificacion = "COMPRA_EN_ERROR"
)

type Notificacion struct {
	Tipo         TipoNotificacion
	IDCompra     int
	Descripcion  string
	Destinatario string
}

// ServicioNotificaciones gestiona el envío de notificaciones
type ServicioNotificaciones struct {
	config       ConfiguracionNotificaciones
	emailService *ServicioEmail
	pushService  *ServicioPush
}

func NuevoServicioNotificaciones(config ConfiguracionNotificaciones) *ServicioNotificaciones {
	return &ServicioNotificaciones{
		config:       config,
		emailService: NuevoServicioEmail(config.ConfiguracionSMTP),
		pushService:  NuevoServicioPush(config.ConfiguracionPush),
	}
}

func (s *ServicioNotificaciones) Notificar(notificacion Notificacion) error {
	var errores []error

	if s.config.EmailHabilitado {
		if err := s.emailService.Enviar(notificacion); err != nil {
			errores = append(errores, fmt.Errorf("error en notificación por email: %w", err))
		}
	}

	if s.config.PushHabilitado {
		if err := s.pushService.Enviar(notificacion); err != nil {
			errores = append(errores, fmt.Errorf("error en notificación push: %w", err))
		}
	}

	if len(errores) > 0 {
		for _, err := range errores {
			log.Printf("Error de notificación: %v", err)
		}
		return fmt.Errorf("ocurrieron %d errores al enviar notificaciones", len(errores))
	}

	return nil
}

func (s *ServicioNotificaciones) NotificarCambioEstadoCompra(idCompra int, tipo TipoNotificacion, descripcion string, destinatario string) error {
	notificacion := Notificacion{
		Tipo:         tipo,
		IDCompra:     idCompra,
		Descripcion:  descripcion,
		Destinatario: destinatario,
	}

	return s.Notificar(notificacion)
}
