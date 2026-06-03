//go:build darwin
// +build darwin

package core

func IsWindows() bool {
	return false
}

func IsLinux() bool {
	return false
}

func CopyCmd() string {
	return "cp -r"
}

func RemoveCmd() string {
	return "rm -fr"
}

func MoveCmd() string {
	return "mv"
}
