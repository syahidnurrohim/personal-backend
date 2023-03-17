package utils

import (
	"errors"
	"log"
	"personal-backend/config"
	"reflect"
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

func ToMap(val interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Map {
		return nil, errors.New("assigned value is not the type of map")
	}

	retMap := make(map[string]interface{})
	for _, key := range v.MapKeys() {
		strct := v.MapIndex(key)
		retMap[key.String()] = strct.Interface()
	}
	return retMap, nil
}

func ToSlice(val interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("assigned value is not the type of slice")
	}

	retSlice := make([]interface{}, v.Len())
	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			retSlice[i] = v.Index(i).Interface()
		}
	}

	return retSlice, nil
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
