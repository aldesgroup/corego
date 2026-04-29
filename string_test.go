//------------------------------------------------------------------------------
//
// Tests for the string util code
//
//------------------------------------------------------------------------------
// Copyright © Oscar Ayoun, Julien Wan
//------------------------------------------------------------------------------

package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPascalToSnake(t *testing.T) {
	assert.Equal(t, "url_pattern_test", PascalToSnake("URLPatternTest"))
	assert.Equal(t, "hello_world", PascalToSnake("HelloWorld"))
	assert.Equal(t, "hello_world", PascalToSnake("hello_world"))
	assert.Equal(t, "id", PascalToSnake("ID"))
	assert.Equal(t, "raw_json_name", PascalToSnake("RawJSONName"))
}

func TestPascalToCamel(t *testing.T) {
	assert.Equal(t, "urlPatternTest", PascalToCamel("URLPatternTest"))
	assert.Equal(t, "helloWorld", PascalToCamel("HelloWorld"))
	assert.Equal(t, "id", PascalToCamel("ID"))
	assert.Equal(t, "deviceId", PascalToCamel("DeviceID"))
	assert.Equal(t, "deviceIds", PascalToCamel("DeviceIDS"))
	assert.Equal(t, "deviceIdString", PascalToCamel("DeviceIDString"))
	assert.Equal(t, "rawJsonName", PascalToCamel("RawJSONName"))
	assert.Equal(t, "defaultHttpError", PascalToCamel("DefaultHTTPError"))
}
