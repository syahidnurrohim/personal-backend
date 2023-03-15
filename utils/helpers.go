package utils

import (
	"errors"
	"log"
	"personal-backend/config"
	"time"
)

func GetIndependentMapKeysValues(val map[string]interface{}) ([]string, []interface{}) {
	keys := make([]string, 0, len(val))
	values := make([]interface{}, 0, len(val))
	for k, v := range val {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}

func DB() *pgConnection {
	if config.PGC == nil {
		err := errors.New("database not initialized")
		Logger().Error(err)
		log.Fatal(err)
		return nil
	}
	return PGConnection(config.PGC)
}

func SetInterval(fn func(), interval time.Duration) {
	ticker := time.NewTicker(interval)
	fn()
	go func() {
		for {
			select {
			case <-ticker.C:
				fn()
			}
		}
	}()
}
