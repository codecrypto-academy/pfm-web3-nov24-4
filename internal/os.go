package internal

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
		log.Infof("Directorio creado exitosamente: %s", path)
	} else {
		log.Infof("El directorio ya existe: %s", path)
	}
	return nil
}
