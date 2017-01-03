package configparser

import (
	"encoding/json"
	"errors"
	"io"
	"sync"

	"github.com/kingzbauer/json_cli/jsongear"
)

// ConfigParser provides the needed machinery to retrieve configuration values
// of any basic type
type ConfigParser interface {
	// Get searches and retrieves any kind of data type
	Get(string) interface{}
	// GetFloat checks for the given key value of type `float`
	// If such a key exists but of a different data type, returns `ErrIncompatibleType` error
	// If it misses the key, returns ErrorDoesNotExist
	GetFloat(string) (float64, error)
	// GetBool checks for the given key value of type `bool`
	// If such a key exists but of a different data type, returns `ErrIncompatibleType` error
	// If it misses the key, returns ErrorDoesNotExist
	GetBool(string) (bool, error)
	// GetString checks for the given key value of type `string`
	// If such a key exists but of a different data type, returns `ErrIncompatibleType` error
	// If it misses the key, returns ErrorDoesNotExist
	GetString(string) (string, error)
}

var (
	// ErrDoesNotExist is the error returned when the value being retrieved is not defined
	// in the specified configuration
	ErrDoesNotExist = errors.New("Config value does not exist")
	// ErrIncompatibleType is the error returned when the value being retrieved is not
	// compatible with the specified
	ErrIncompatibleType = errors.New("Incompatible type for stored value")
)

// NewJSONConfigReader reads from reader which should be valid json, and returns a ConfigParser
func NewJSONConfigReader(reader io.Reader) (ConfigParser, error) {
	config := &configParser{
		floatMap:  map[string]float64{},
		boolMap:   map[string]bool{},
		stringMap: map[string]string{},
	}
	if err := json.NewDecoder(reader).Decode(&config.data); err != nil {
		return nil, err
	}

	return config, nil
}

type configParser struct {
	data       interface{}
	floatMap   map[string]float64
	boolMap    map[string]bool
	stringMap  map[string]string
	floatLock  sync.RWMutex
	boolLock   sync.RWMutex
	stringLock sync.RWMutex
}

func (config *configParser) Get(key string) interface{} {
	return jsongear.Get(key, config.data)
}

func (config *configParser) GetFloat(key string) (float64, error) {
	config.floatLock.RLock()
	if v, ok := config.floatMap[key]; ok {
		config.floatLock.RUnlock()
		return v, nil
	}
	config.floatLock.RUnlock()

	v := jsongear.Get(key, config.data)
	if v == nil {
		return 0, ErrDoesNotExist
	}

	if floatV, ok := v.(float64); ok {
		config.floatLock.Lock()
		config.floatMap[key] = floatV
		config.floatLock.Unlock()
		return floatV, nil
	}
	return 0, ErrIncompatibleType
}

func (config *configParser) GetString(key string) (string, error) {
	config.stringLock.RLock()
	if v, ok := config.stringMap[key]; ok {
		config.stringLock.RUnlock()
		return v, nil
	}
	config.stringLock.RUnlock()

	v := jsongear.Get(key, config.data)
	if v == nil {
		return "", ErrDoesNotExist
	}

	if strV, ok := v.(string); ok {
		config.stringLock.Lock()
		config.stringMap[key] = strV
		config.stringLock.Unlock()
		return strV, nil
	}
	return "", ErrIncompatibleType
}

func (config *configParser) GetBool(key string) (bool, error) {
	config.boolLock.RLock()
	if v, ok := config.boolMap[key]; ok {
		config.boolLock.RUnlock()
		return v, nil
	}
	config.boolLock.RUnlock()

	v := jsongear.Get(key, config.data)
	if v == nil {
		return false, ErrDoesNotExist
	}

	if boolV, ok := v.(bool); ok {
		config.boolLock.Lock()
		config.boolMap[key] = boolV
		config.boolLock.Unlock()
		return boolV, nil
	}
	return false, ErrIncompatibleType
}
