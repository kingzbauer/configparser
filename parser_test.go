package configparser

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

var MapJSON = `{
        "Id": "d836a5e40aa8974d7076e791ba3c14726bf2dd2cd079652477d6827973969130",
        "Created": "2016-08-31T16:49:33.119587574Z",
        "Path": "/bin/bash",
        "Args": [{"bool": true}],
        "Level": 34.7,
        "Running": true,
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false
        }
     }`

var (
	config ConfigParser
)

func TestGet(t *testing.T) {
	expectedV := "d836a5e40aa8974d7076e791ba3c14726bf2dd2cd079652477d6827973969130"
	if !reflect.DeepEqual(expectedV, config.Get("Id")) {
		t.Errorf("Expected %#v, got %#v", expectedV, config.Get("Id"))
	}

	// Test missing value
	returnedV := config.Get("NoThere")
	if returnedV != nil {
		t.Errorf("Expected %#v, got %#v", nil, returnedV)
	}
}

func TestGetFloat(t *testing.T) {
	expected := float64(34.7)
	received, _ := config.GetFloat("Level")
	if expected != received {
		t.Errorf("Expected %#v, got %#v", expected, received)
	}

	// Test out a value that does not exist
	_, err := config.GetFloat("FloatNone")
	if err != ErrDoesNotExist {
		t.Errorf("Expected %#v, got %#v", ErrDoesNotExist, received)
	}

	// test out incompatible types
	_, err = config.GetFloat("Id")
	if err != ErrIncompatibleType {
		t.Errorf("Expected %#v, got %#v", ErrIncompatibleType, err)
	}
}

func TestGetBool(t *testing.T) {
	received, _ := config.GetBool("Running")
	if !received {
		t.Errorf("Expected %#v, got %#v", true, received)
	}

	// Test out a value that does not exist
	_, err := config.GetBool("Nowhere")
	if err != ErrDoesNotExist {
		t.Errorf("Expected %#v, got %#v", ErrDoesNotExist, received)
	}

	// test out incompatible types
	_, err = config.GetBool("Id")
	if err != ErrIncompatibleType {
		t.Errorf("Expected %#v, got %#v", ErrIncompatibleType, err)
	}
}

func TestGetString(t *testing.T) {
	received, _ := config.GetString("Id")
	expectedV := "d836a5e40aa8974d7076e791ba3c14726bf2dd2cd079652477d6827973969130"
	if received != expectedV {
		t.Errorf("Expected %#v, got %#v", expectedV, received)
	}

	_, err := config.GetString("Nowhere")
	if err != ErrDoesNotExist {
		t.Errorf("Expected %#v, got %#v", expectedV, received)
	}

	_, err = config.GetString("Running")
	if err != ErrIncompatibleType {
		t.Errorf("Expected %#v, got %#v", expectedV, received)
	}
}

func TestMain(m *testing.M) {
	config, _ = NewJSONConfigReader(bytes.NewReader([]byte(MapJSON)))
	os.Exit(m.Run())
}
