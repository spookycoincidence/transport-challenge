package notification

import (
	"fmt"
	"os"
)

// ConfiguracionNotificaciones representa la configuración para diferentes canales de notificación
type ConfiguracionNotificaciones struct {
	EmailHabilitado   bool
	PushHabilitado    bool
	ConfiguracionSMTP ConfigSMTP
	ConfiguracionPush ConfigPush
}

// ConfigSMTP contiene la configuración para envío de emails
type ConfigSMTP struct {
	Host      string
	Puerto    int
	Usuario   string
	Clave     string
	Remitente string
}

// ConfigPush contiene la configuración para notificaciones push
type ConfigPush struct {
	ServidorAPI string
	ClaveAPI    string
}

// CargarConfiguracionDesdeVariablesEntorno carga la configuración de notificaciones desde variables de entorno
func CargarConfiguracionDesdeVariablesEntorno() ConfiguracionNotificaciones {
	return ConfiguracionNotificaciones{
		EmailHabilitado: os.Getenv("EMAIL_HABILITADO") == "true",
		PushHabilitado:  os.Getenv("PUSH_HABILITADO") == "true",
		ConfiguracionSMTP: ConfigSMTP{
			Host:      os.Getenv("SMTP_HOST"),
			Puerto:    8025, // Puerto por defecto, se puede configurar
			Usuario:   os.Getenv("SMTP_USUARIO"),
			Clave:     os.Getenv("SMTP_CLAVE"),
			Remitente: os.Getenv("SMTP_REMITENTE"),
		},
		ConfiguracionPush: ConfigPush{
			ServidorAPI: os.Getenv("PUSH_SERVIDOR_API"),
			ClaveAPI:    os.Getenv("PUSH_CLAVE_API"),
		},
	}
}

// Validar verifica que la configuración tenga los parámetros necesarios
func (c *ConfiguracionNotificaciones) Validar() error {
	if c.EmailHabilitado {
		if c.ConfiguracionSMTP.Host == "" {
			return fmt.Errorf("host SMTP es requerido para notificaciones por email")
		}
		if c.ConfiguracionSMTP.Usuario == "" {
			return fmt.Errorf("usuario SMTP es requerido para notificaciones por email")
		}
	}

	if c.PushHabilitado {
		if c.ConfiguracionPush.ServidorAPI == "" {
			return fmt.Errorf("servidor API es requerido para notificaciones push")
		}
	}

	return nil
}
