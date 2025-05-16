package app

import (
	"crypto/tls"
	"fmt"
	"messenger/internal/config/loaders"
	"messenger/internal/config/models"
)

func loadAppCertificateConfig(certConfig models.Certificate) (tls.Config, error) {
	certificate, err := loaders.LoadCertificate(certConfig)
	if err != nil {
		return tls.Config{}, fmt.Errorf("ошибка загрузки сертификата: %w", err)
	}

	return tls.Config{
		Certificates: []tls.Certificate{certificate},
		MinVersion:   tls.VersionTLS12,
	}, nil
}
