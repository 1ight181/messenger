package models

import (
	"errors"
)

type Certificate struct {
	CertificateFileName string `mapstructure:"cert_file_name"`
	KeyFileName         string `mapstructure:"key_file_name"`
	CertificatePath     string `mapstructure:"cert_file_path"`
	KeyPath             string `mapstructure:"key_file_path"`
}

// Validate проверяет структуру Certificate на наличие обязательных полей
// и гарантирует, что ни одно из них не пустое. Возвращает ошибку, если
// отсутствует одно из следующих полей:
// - CertificateFileName: имя файла сертификата.
// - KeyFileName: имя файла ключа.
// - CertificatePath: путь к файлу сертификата.
// - KeyPath: путь к файлу ключа.
func (c *Certificate) Validate() error {
	if c.CertificateFileName == "" {
		return errors.New("требуется имя файла сертификата")
	}
	if c.KeyFileName == "" {
		return errors.New("требуется имя файла ключа")
	}
	if c.CertificatePath == "" {
		return errors.New("требуется путь к сертификату")
	}
	if c.KeyPath == "" {
		return errors.New("требуется путь к ключу")
	}
	return nil
}
