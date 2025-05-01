package utility

import (
	"crypto/tls"
	"fmt"
	"messager/internal/config/models"
)

// LoadCertificateConfig создает полные пути к файлам сертификата и ключа
// на основе предоставленной конфигурации.
//
// Параметры:
//   - certificateConfig: Указатель на структуру models.Certificate, содержащую конфигурацию сертификата.
//
// Возвращает:
//   - string: Полный путь к файлу сертификата.
//   - string: Полный путь к файлу ключа сертификата.
//   - error: Ошибка, если какие-либо обязательные поля конфигурации пусты.
//
// Ошибки:
//   - Возвращается ошибка, если не удалось загрузить сертификат
func LoadCertificate(certificateConfig models.Certificate) (tls.Certificate, error) {
	certificatePath := certificateConfig.CertificatePath
	certificateKeyPath := certificateConfig.KeyPath
	certificateFileName := certificateConfig.CertificateFileName
	certificateKeyFileName := certificateConfig.KeyFileName

	certificateFullPath := fmt.Sprintf("%s/%s", certificatePath, certificateFileName)
	certificateKeyFullPath := fmt.Sprintf("%s/%s", certificateKeyPath, certificateKeyFileName)

	cert, err := tls.LoadX509KeyPair(certificateFullPath, certificateKeyFullPath)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("не удалось загрузить сертификат \n"+
			"сертификат: %s,\n"+
			"ключ: %s",
			certificateFileName,
			certificateKeyFileName)
	}

	return cert, nil
}
