package notification

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

// ServicioEmail maneja el envío de notificaciones por correo electrónico
type ServicioEmail struct {
	config ConfigSMTP
}

func NuevoServicioEmail(config ConfigSMTP) *ServicioEmail {
	return &ServicioEmail{
		config: config,
	}
}

func (s *ServicioEmail) Enviar(notificacion Notificacion) error {

	if s.config.Host == "" || s.config.Usuario == "" || s.config.Clave == "" {
		return fmt.Errorf("configuración de email incompleta")
	}

	// Se construye el mensaje
	asunto := fmt.Sprintf("Notificación de Compra #%d", notificacion.IDCompra)
	cuerpo := fmt.Sprintf(`
		Tipo de Notificación: %s
		ID de Compra: %d
		Descripción: %s
	`, notificacion.Tipo, notificacion.IDCompra, notificacion.Descripcion)

	mensaje := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s", s.config.Remitente, notificacion.Destinatario, asunto, cuerpo)

	auth := smtp.PlainAuth("", s.config.Usuario, s.config.Clave, s.config.Host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         s.config.Host,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", s.config.Host, s.config.Puerto), tlsConfig)
	if err != nil {
		return fmt.Errorf("error al establecer conexión TLS: %w", err)
	}

	cliente, err := smtp.NewClient(conn, s.config.Host)
	if err != nil {
		return fmt.Errorf("error al crear cliente SMTP: %w", err)
	}
	defer cliente.Close()

	if err = cliente.Auth(auth); err != nil {
		return fmt.Errorf("error de autenticación: %w", err)
	}

	if err = cliente.Mail(s.config.Remitente); err != nil {
		return fmt.Errorf("error al configurar remitente: %w", err)
	}

	if err = cliente.Rcpt(notificacion.Destinatario); err != nil {
		return fmt.Errorf("error al configurar destinatario: %w", err)
	}

	w, err := cliente.Data()
	if err != nil {
		return fmt.Errorf("error al preparar datos: %w", err)
	}

	_, err = w.Write([]byte(mensaje))
	if err != nil {
		return fmt.Errorf("error al escribir mensaje: %w", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("error al cerrar escritura: %w", err)
	}

	err = cliente.Quit()
	if err != nil {
		return fmt.Errorf("error al cerrar conexión SMTP: %w", err)
	}

	return nil
}
