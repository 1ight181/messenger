package models

import "fmt"

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
		return fmt.Errorf("требуется имя файла сертификата")
	}
	if c.KeyFileName == "" {
		return fmt.Errorf("требуется имя файла ключа")
	}
	if c.CertificatePath == "" {
		return fmt.Errorf("требуется путь к сертификату")
	}
	if c.KeyPath == "" {
		return fmt.Errorf("требуется путь к ключу")
	}
	return nil
}
