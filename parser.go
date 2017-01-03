package configparser

import (
	"errors"
	"io"
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
	// ErrDecoding is returned when the content of the config reader could not be decoded
	// with the provided format
	ErrDecoding = errors.New("Error decoding the reader")
)

// NewJSONConfigReader reads from reader which should be valid json, and returns a ConfigParser
func NewJSONConfigReader(reader io.Reader) (ConfigParser, error) {
	return nil, nil
}
