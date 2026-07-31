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
	assert.Equal(t, "raw_json_name_1", PascalToSnake("RawJSONName1"))
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
	assert.Equal(t, "branch1Type", PascalToCamel("Branch1Type"))
}

func TestPascalToShort(t *testing.T) {
	assert.Equal(t, "UrlPatTes", PascalToShort("URLPatternTest"))
	assert.Equal(t, "HelWor", PascalToShort("HelloWorld"))
	assert.Equal(t, "Id", PascalToShort("ID"))
	assert.Equal(t, "DevId", PascalToShort("DeviceID"))
	assert.Equal(t, "DevIds", PascalToShort("DeviceIDS"))
	assert.Equal(t, "DevIdStr", PascalToShort("DeviceIDString"))
	assert.Equal(t, "RawJsoNam", PascalToShort("RawJSONName"))
	assert.Equal(t, "DefHttErr", PascalToShort("DefaultHTTPError"))
	assert.Equal(t, "Bra1Typ", PascalToShort("Branch1Type"))
	assert.Equal(t, "RanAccMem", PascalToShort("RandomAccessMemory"))
}

func TestElude(t *testing.T) {
	assert.Equal(t, "Hello World", Elude("Hello World", 20))
	assert.Equal(t, "Hell (...)", Elude("Hello World", 10))
	assert.Equal(t, "Hello (...)", Elude("Hello World", 11))
	assert.Equal(t, "Hello  (...)", Elude("Hello World", 12))
	assert.Equal(t, "Hello W (...)", Elude("Hello World", 13))
	assert.Equal(t, "Hello Wo (...)", Elude("Hello World", 14))
	assert.Equal(t, "Hello Wor (...)", Elude("Hello World", 15))
}

func TestKebabToPascal(t *testing.T) {
	assert.Equal(t, "UrlPatTes", KebabToPascal("url-pat-tes"))
	assert.Equal(t, "HelWor", KebabToPascal("hel-wor"))
	assert.Equal(t, "Id", KebabToPascal("id"))
	assert.Equal(t, "DevId", KebabToPascal("dev-id"))
	assert.Equal(t, "DevIds", KebabToPascal("dev-ids"))
	assert.Equal(t, "DevIdStr", KebabToPascal("dev-id-str"))
	assert.Equal(t, "RawJsoNam", KebabToPascal("raw-jso-nam"))
	assert.Equal(t, "DefHttErr", KebabToPascal("def-htt-err"))
	assert.Equal(t, "Bra1Typ", KebabToPascal("bra-1-typ"))
	assert.Equal(t, "RanAccMem", KebabToPascal("ran-acc-mem"))
}

func TestSnakeToPascal(t *testing.T) {
	assert.Equal(t, "UrlPatTes", SnakeToPascal("url_pat_tes"))
	assert.Equal(t, "HelWor", SnakeToPascal("hel_wor"))
	assert.Equal(t, "Id", SnakeToPascal("id"))
	assert.Equal(t, "DevId", SnakeToPascal("dev_id"))
	assert.Equal(t, "DevIds", SnakeToPascal("dev_ids"))
	assert.Equal(t, "DevIdStr", SnakeToPascal("dev_id_str"))
	assert.Equal(t, "RawJsoNam", SnakeToPascal("raw_jso_nam"))
	assert.Equal(t, "DefHttErr", SnakeToPascal("def_htt_err"))
	assert.Equal(t, "Bra1Typ", SnakeToPascal("bra_1_typ"))
	assert.Equal(t, "RanAccMem", SnakeToPascal("ran_acc_mem"))
}

func TestSnakeToCamel(t *testing.T) {
	assert.Equal(t, "urlPatTes", SnakeToCamel("url_pat_tes"))
	assert.Equal(t, "helWor", SnakeToCamel("hel_wor"))
	assert.Equal(t, "id", SnakeToCamel("id"))
	assert.Equal(t, "devId", SnakeToCamel("dev_id"))
	assert.Equal(t, "devIds", SnakeToCamel("dev_ids"))
	assert.Equal(t, "devIdStr", SnakeToCamel("dev_id_str"))
	assert.Equal(t, "rawJsoNam", SnakeToCamel("raw_jso_nam"))
	assert.Equal(t, "defHttErr", SnakeToCamel("def_htt_err"))
	assert.Equal(t, "bra1Typ", SnakeToCamel("bra_1_typ"))
	assert.Equal(t, "ranAccMem", SnakeToCamel("ran_acc_mem"))
}

func TestPadLeft(t *testing.T) {
	assert.Equal(t, "     Hello", PadLeft("Hello", 10, " "))
	assert.Equal(t, "Hello", PadLeft("Hello", 5, " "))
	assert.Equal(t, "Hello", PadLeft("Hello", 3, " "))
	assert.Equal(t, "00000Hello", PadLeft("Hello", 10, "0"))
	assert.Equal(t, "Hello", PadLeft("Hello", 5, "0"))
	assert.Equal(t, "Hello", PadLeft("Hello", 3, "0"))
}

func TestPadRight(t *testing.T) {
	assert.Equal(t, "Hello     ", PadRight("Hello", 10, " "))
	assert.Equal(t, "Hello", PadRight("Hello", 5, " "))
	assert.Equal(t, "Hello", PadRight("Hello", 3, " "))
	assert.Equal(t, "Hello00000", PadRight("Hello", 10, "0"))
	assert.Equal(t, "Hello", PadRight("Hello", 5, "0"))
	assert.Equal(t, "Hello", PadRight("Hello", 3, "0"))
}
