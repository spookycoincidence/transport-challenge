package notification

import (
	"os"
	"testing"
)

func TestCargarConfiguracionDesdeVariablesEntorno(t *testing.T) {
	// Configurar variables de entorno
	os.Setenv("EMAIL_HABILITADO", "true")
	os.Setenv("PUSH_HABILITADO", "true")
	os.Setenv("SMTP_HOST", "smtp.ejemplo.com")
	os.Setenv("SMTP_USUARIO", "usuario_prueba")
	os.Setenv("SMTP_CLAVE", "clave_prueba")
	os.Setenv("SMTP_REMITENTE", "remitente@ejemplo.com")
	os.Setenv("PUSH_SERVIDOR_API", "https://api-push.ejemplo.com")
	os.Setenv("PUSH_CLAVE_API", "token_push")
	defer func() {
		// Limpiar variables de entorno
		os.Unsetenv("EMAIL_HABILITADO")
		os.Unsetenv("PUSH_HABILITADO")
		os.Unsetenv("SMTP_HOST")
		os.Unsetenv("SMTP_USUARIO")
		os.Unsetenv("SMTP_CLAVE")
		os.Unsetenv("SMTP_REMITENTE")
		os.Unsetenv("PUSH_SERVIDOR_API")
		os.Unsetenv("PUSH_CLAVE_API")
	}()

	// Ejecutar función de carga de configuración
	configuracion := CargarConfiguracionDesdeVariablesEntorno()

	// Verificar que los valores se hayan cargado correctamente
	if !configuracion.EmailHabilitado {
		t.Errorf("Se esperaba que EmailHabilitado fuera true")
	}

	if !configuracion.PushHabilitado {
		t.Errorf("Se esperaba que PushHabilitado fuera true")
	}

	if configuracion.ConfiguracionSMTP.Host != "smtp.ejemplo.com" {
		t.Errorf("Host SMTP incorrecto. Esperado: smtp.ejemplo.com, Obtenido: %s", configuracion.ConfiguracionSMTP.Host)
	}

	if configuracion.ConfiguracionSMTP.Usuario != "usuario_prueba" {
		t.Errorf("Usuario SMTP incorrecto. Esperado: usuario_prueba, Obtenido: %s", configuracion.ConfiguracionSMTP.Usuario)
	}

	if configuracion.ConfiguracionPush.ServidorAPI != "https://api-push.ejemplo.com" {
		t.Errorf("Servidor Push incorrecto. Esperado: https://api-push.ejemplo.com, Obtenido: %s", configuracion.ConfiguracionPush.ServidorAPI)
	}
}

func TestValidarConfiguracion(t *testing.T) {
	testCases := []struct {
		nombre        string
		configuracion ConfiguracionNotificaciones
		esperaError   bool
		mensajeError  string
	}{
		{
			nombre: "Configuración Completa",
			configuracion: ConfiguracionNotificaciones{
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
			},
			esperaError: false,
		},
		{
			nombre: "Email Habilitado Sin Host",
			configuracion: ConfiguracionNotificaciones{
				EmailHabilitado: true,
				ConfiguracionSMTP: ConfigSMTP{
					Host: "",
				},
			},
			esperaError:  true,
			mensajeError: "host SMTP es requerido para notificaciones por email",
		},
		{
			nombre: "Email Habilitado Sin Usuario",
			configuracion: ConfiguracionNotificaciones{
				EmailHabilitado: true,
				ConfiguracionSMTP: ConfigSMTP{
					Host:    "smtp.ejemplo.com",
					Usuario: "",
				},
			},
			esperaError:  true,
			mensajeError: "usuario SMTP es requerido para notificaciones por email",
		},
		{
			nombre: "Push Habilitado Sin Servidor",
			configuracion: ConfiguracionNotificaciones{
				PushHabilitado: true,
				ConfiguracionPush: ConfigPush{
					ServidorAPI: "",
				},
			},
			esperaError:  true,
			mensajeError: "servidor API es requerido para notificaciones push",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.nombre, func(t *testing.T) {
			err := tc.configuracion.Validar()

			if tc.esperaError && err == nil {
				t.Errorf("Se esperaba un error para el caso: %s", tc.nombre)
			}

			if tc.esperaError && err != nil {
				if err.Error() != tc.mensajeError {
					t.Errorf("Mensaje de error incorrecto. Esperado: %s, Obtenido: %s",
						tc.mensajeError, err.Error())
				}
			}

			if !tc.esperaError && err != nil {
				t.Errorf("No se esperaba un error para el caso: %s. Error: %v", tc.nombre, err)
			}
		})
	}
}
