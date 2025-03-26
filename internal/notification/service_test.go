package notification

import (
	"fmt"
	"testing"
)

// Mock de ServicioEmail para pruebas
type mockServicioEmail struct {
	enviado bool
	error   error
}

func (m *mockServicioEmail) Enviar(notificacion Notificacion) error {
	m.enviado = true
	return m.error
}

// Mock de ServicioPush para pruebas
type mockServicioPush struct {
	enviado bool
	error   error
}

func (m *mockServicioPush) Enviar(notificacion Notificacion) error {
	m.enviado = true
	return m.error
}

func TestNuevoServicioNotificaciones(t *testing.T) {
	config := ConfiguracionNotificaciones{
		EmailHabilitado: true,
		PushHabilitado:  true,
		ConfiguracionSMTP: ConfigSMTP{
			Host:      "smtp.ejemplo.com",
			Usuario:   "usuario",
			Clave:     "clave",
			Remitente: "remitente@ejemplo.com",
		},
		ConfiguracionPush: ConfigPush{
			ServidorAPI: "https://api-push.ejemplo.com",
			ClaveAPI:    "token",
		},
	}

	servicio := NuevoServicioNotificaciones(config)

	if servicio == nil {
		t.Fatal("Se esperaba que el servicio de notificaciones no fuera nulo")
	}
}

func TestNotificar(t *testing.T) {
	testCases := []struct {
		nombre           string
		configuracion    ConfiguracionNotificaciones
		notificacion     Notificacion
		esperaErrorEmail bool
		esperaErrorPush  bool
	}{
		{
			nombre: "Notificación Exitosa",
			configuracion: ConfiguracionNotificaciones{
				EmailHabilitado: true,
				PushHabilitado:  true,
			},
			notificacion: Notificacion{
				Tipo:         NotificacionCompraCreada,
				IDCompra:     123,
				Descripcion:  "Compra creada exitosamente",
				Destinatario: "cliente@ejemplo.com",
			},
			esperaErrorEmail: false,
			esperaErrorPush:  false,
		},
		{
			nombre: "Error en Email",
			configuracion: ConfiguracionNotificaciones{
				EmailHabilitado: true,
				PushHabilitado:  false,
			},
			notificacion: Notificacion{
				Tipo:         NotificacionCompraCreada,
				IDCompra:     123,
				Descripcion:  "Compra creada con error en email",
				Destinatario: "cliente@ejemplo.com",
			},
			esperaErrorEmail: true,
			esperaErrorPush:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.nombre, func(t *testing.T) {
			mockEmail := &mockServicioEmail{error: nil}
			mockPush := &mockServicioPush{error: nil}

			if tc.esperaErrorEmail {
				mockEmail.error = fmt.Errorf("error simulado en email")
			}
			if tc.esperaErrorPush {
				mockPush.error = fmt.Errorf("error simulado en push")
			}

			servicio := &ServicioNotificaciones{
				config:       tc.configuracion,
				emailService: mockEmail,
				pushService:  mockPush,
			}

			err := servicio.Notificar(tc.notificacion)

			if tc.configuracion.EmailHabilitado {
				if !mockEmail.enviado {
					t.Errorf("Se esperaba que el servicio de email enviara una notificación")
				}
				if tc.esperaErrorEmail && err == nil {
					t.Errorf("Se esperaba un error en email")
				}
			}

			if tc.configuracion.PushHabilitado {
				if !mockPush.enviado {
					t.Errorf("Se esperaba que el servicio de push enviara una notificación")
				}
				if tc.esperaErrorPush && err == nil {
					t.Errorf("Se esperaba un error en push")
				}
			}
		})
	}
}

func TestNotificarCambioEstadoCompra(t *testing.T) {
	config := ConfiguracionNotificaciones{
		EmailHabilitado: true,
		PushHabilitado:  true,
	}

	servicio := NuevoServicioNotificaciones(config)

	err := servicio.NotificarCambioEstadoCompra(
		456,
		NotificacionCompraEntregada,
		"Compra entregada con éxito",
		"cliente@ejemplo.com",
	)

	if err != nil {
		t.Errorf("No se esperaba un error al notificar cambio de estado: %v", err)
	}
}
