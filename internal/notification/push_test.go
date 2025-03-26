package notification

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNuevoServicioPush(t *testing.T) {
	config := ConfigPush{
		ServidorAPI: "https://api-push.ejemplo.com",
		ClaveAPI:    "token_prueba",
	}

	servicio := NuevoServicioPush(config)

	if servicio == nil {
		t.Fatal("Se esperaba que el servicio de push no fuera nulo")
	}

	if servicio.config != config {
		t.Errorf("La configuración del servicio no coincide con la configuración proporcionada")
	}
}

func TestEnviarPush(t *testing.T) {
	testCases := []struct {
		nombre            string
		configuracionPush ConfigPush
		notificacion      Notificacion
		respuestaServidor int
		esperaError       bool
	}{
		{
			nombre: "Envío de Notificación Push Exitoso",
			configuracionPush: ConfigPush{
				ServidorAPI: "/test-push",
				ClaveAPI:    "token_valido",
			},
			notificacion: Notificacion{
				Tipo:         NotificacionCompraCreada,
				IDCompra:     123,
				Descripcion:  "Compra creada exitosamente",
				Destinatario: "usuario@ejemplo.com",
			},
			respuestaServidor: http.StatusOK,
			esperaError:       false,
		},
		{
			nombre: "Configuración Push Incompleta",
			configuracionPush: ConfigPush{
				ServidorAPI: "",
				ClaveAPI:    "",
			},
			notificacion: Notificacion{
				Tipo:         NotificacionCompraCreada,
				IDCompra:     123,
				Descripcion:  "Compra creada",
				Destinatario: "usuario@ejemplo.com",
			},
			respuestaServidor: http.StatusOK,
			esperaError:       true,
		},
		{
			nombre: "Error en Respuesta del Servidor",
			configuracionPush: ConfigPush{
				ServidorAPI: "/test-push",
				ClaveAPI:    "token_valido",
			},
			notificacion: Notificacion{
				Tipo:         NotificacionCompraCreada,
				IDCompra:     123,
				Descripcion:  "Compra creada",
				Destinatario: "usuario@ejemplo.com",
			},
			respuestaServidor: http.StatusInternalServerError,
			esperaError:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.nombre, func(t *testing.T) {
			// Crear un servidor de prueba
			servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.respuestaServidor)
			}))
			defer servidor.Close()

			// Modificar la configuración con el URL del servidor de prueba
			configPush := tc.configuracionPush
			if configPush.ServidorAPI == "/test-push" {
				configPush.ServidorAPI = servidor.URL
			}

			servicio := NuevoServicioPush(configPush)
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
