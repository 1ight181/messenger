package loaders

import (
	"crypto/tls"
	"fmt"
	"messenger/internal/config/models"
)

// LoadCertificate загружает X.509 сертификат и соответствующий закрытый ключ
// из указанных файловых путей, предоставленных в конфигурации Certificate.
//
// Параметры:
//   - certificateConfig (models.Certificate): Структура, содержащая пути и
//     имена файлов для сертификата и ключа.
//
// Возвращает:
//   - (tls.Certificate): Загруженный TLS сертификат.
//   - (error): Ошибка, если сертификат или ключ не удалось загрузить.
//
// Функция формирует полные пути к файлам сертификата и ключа, используя
// предоставленные пути к директориям и имена файлов, затем пытается загрузить
// их с помощью tls.LoadX509KeyPair. В случае неудачи возвращается ошибка с
// деталями об именах файлов сертификата и ключа.
func LoadCertificate(certificateConfig models.Certificate) (tls.Certificate, error) {
	certificatePath := certificateConfig.CertificatePath
	certificateKeyPath := certificateConfig.KeyPath
	certificateFileName := certificateConfig.CertificateFileName
	certificateKeyFileName := certificateConfig.KeyFileName

	certificateFullPath := fmt.Sprintf("%s%s", certificatePath, certificateFileName)
	certificateKeyFullPath := fmt.Sprintf("%s%s", certificateKeyPath, certificateKeyFileName)

	cert, err := tls.LoadX509KeyPair(certificateFullPath, certificateKeyFullPath)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("не удалось загрузить сертификат \n"+
			"сертификат: %s,\n"+
			"ключ: %s \n"+
			"ошибка: %s",
			certificateFullPath,
			certificateKeyFullPath,
			err)
	}

	return cert, nil
}
