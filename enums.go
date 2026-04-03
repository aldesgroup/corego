package core

import (
	"strings"
)

// ------------------------------------------------------------------------------------------------
// the environment types for our servers
// ------------------------------------------------------------------------------------------------

// EnvType represents the type of environment we're running the app in
type EnvType int

const (
	EnvTypeLOCAL      EnvType = -2
	EnvTypeSANDBOX    EnvType = 1
	EnvTypeSTAGING    EnvType = 2
	EnvTypePRODUCTION EnvType = 3
)

var envTypes = map[int]string{
	int(EnvTypeLOCAL):      "LOCAL",
	int(EnvTypeSANDBOX):    "SANDBOX",
	int(EnvTypeSTAGING):    "STAGING",
	int(EnvTypePRODUCTION): "PRODUCTION",
}

func (thisEnvType EnvType) String() string {
	return envTypes[int(thisEnvType)]
}

// Val helps implement the IEnum interface
func (thisEnvType EnvType) Val() int {
	return int(thisEnvType)
}

// Values helps implement the IEnum interface
func (thisEnvType EnvType) Values() map[int]string {
	return envTypes
}

func EnvTypeValFrom(value string) EnvType {
	for eT, label := range envTypes {
		if label == value {
			return EnvType(eT)
		}
	}

	PanicMsg("There cannot be an 'envtype' named '%s'. Possible values are: %s",
		value, strings.Join(GetSortedValues(envTypes), ", "))
	return 0
}
