package notification

import (
	"testing"
)

func TestNuevoServicioEmail(t *testing.T) {
	config := ConfigSMTP{
		Host:      "smtp.ejemplo.com",
		Puerto:    587,
		Usuario:   "usuario_prueba",
		Clave:     "clave_prueba",
		Remitente: "remitente@ejemplo.com",
	}

	servicio := NuevoServicioEmail(config)

	if servicio == nil {
		t.Fatal("Se esperaba que el servicio de email no fuera nulo")
	}

	if servicio.config != config {
		t.Errorf("La configuración del servicio no coincide con la configuración proporcionada")
	}
}

func TestEnviarEmail(t *testing.T) {
	testCases := []struct {
		nombre            string
		configuracionSMTP ConfigSMTP
		notificacion      Notificacion
		esperaError       bool
	}{
		{
			nombre: "Configuración SMTP Válida",
			configuracionSMTP: ConfigSMTP{
				Host:      "smtp.ejemplo.com",
				Puerto:    587,
				Usuario:   "usuario_prueba",
				Clave:     "clave_prueba",
				Remitente: "remitente@ejemplo.com",
			},
			notificacion: Notificacion{
				Tipo:         NotificacionCompraCreada,
				IDCompra:     123,
				Descripcion:  "Compra creada exitosamente",
				Destinatario: "cliente@ejemplo.com",
			},
			esperaError: false,
		},
		{
			nombre: "Configuración SMTP Incompleta - Sin Host",
			configuracionSMTP: ConfigSMTP{
				Host:      "",
				Usuario:   "usuario_prueba",
				Clave:     "clave_prueba",
				Remitente: "remitente@ejemplo.com",
			},
			notificacion: Notificacion{
				Tipo:         NotificacionCompraCreada,
				IDCompra:     123,
				Descripcion:  "Compra creada",
				Destinatario: "cliente@ejemplo.com",
			},
			esperaError: true,
		},
		{
			nombre: "Configuración SMTP Incompleta - Sin Usuario",
			configuracionSMTP: ConfigSMTP{
				Host:      "smtp.ejemplo.com",
				Usuario:   "",
				Clave:     "clave_prueba",
				Remitente: "remitente@ejemplo.com",
			},
			notificacion: Notificacion{
				Tipo:         NotificacionCompraCreada,
				IDCompra:     123,
				Descripcion:  "Compra creada",
				Destinatario: "cliente@ejemplo.com",
			},
			esperaError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.nombre, func(t *testing.T) {
			servicio := NuevoServicioEmail(tc.configuracionSMTP)
			err := servicio.Enviar(tc.notificacion)

			if tc.esperaError && err == nil {
				t.Errorf("Se esperaba un error para el caso: %s", tc.nombre)
			}

			if !tc.esperaError && err != nil {
				t.Errorf("No se esperaba un error para el caso: %s. Error: %v", tc.nombre, err)
			}
		})
	}
}
