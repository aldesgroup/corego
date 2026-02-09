package core

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

// ------------------------------------------------------------------------------------------------
// Files
// ------------------------------------------------------------------------------------------------

// FileExists checks if a file exists at the specified path.
func FileExists(pathParts ...string) bool {
	fullpath := path.Join(pathParts...)
	info, err := os.Stat(fullpath)
	return !os.IsNotExist(err) && !info.IsDir()
}

// DirExists checks if a file exists at the specified path.
func DirExists(pathParts ...string) bool {
	fullpath := path.Join(pathParts...)
	info, err := os.Stat(fullpath)
	return !os.IsNotExist(err) && info.IsDir()
}

// Returns the given file's modification time, or panics
func EnsureModTime(filePath string) time.Time {
	info, err := os.Stat(filePath)
	PanicMsgIfErr(err, "Could not get modification time for file '%s'", filePath)
	return info.ModTime()
}

// EnsureDir makes sure the directory with the given path elements exists
func EnsureDir(pathElem ...string) string {
	dirname := path.Join(pathElem...)

	PanicMsgIfErr(os.MkdirAll(dirname, 0o777), "Could not create directory '%s'", dirname)

	return dirname
}

// EnsureReadDir is a convenient shortcut to list a directory's files and folders
func EnsureReadDir(pathElem ...string) []os.DirEntry {
	dirPath := path.Join(pathElem...)
	entries, errRead := os.ReadDir(dirPath)
	PanicMsgIfErr(errRead, "Could not read the given folder '%s'", dirPath)
	return entries
}

// WriteToFile writes the given content to the file with the given path
func WriteToFile(content string, filepaths ...string) {
	// creating the file
	fileName := path.Join(filepaths...)

	// creating the missing directory if needed
	if dir := path.Dir(fileName); path.Base(fileName) != fileName {
		EnsureDir(dir)
	}

	// creating the file
	file, errCreate := os.Create(fileName)
	PanicMsgIfErr(errCreate, "Could not create file %s", fileName)

	// ensuring we've got no leak
	defer func() {
		if errClose := file.Close(); errClose != nil {
			slog.Error(fmt.Sprintf("Could not properly close file %s; cause: %s", fileName, errClose))
			os.Exit(1)
		}
	}()

	// writing to file
	if _, errWrite := file.WriteString(content); errWrite != nil {
		PanicMsgIfErr(errWrite, "Could not write file '%s'", fileName)
	}
}

// EnsureNoDir removes the directory with the given path elements
func EnsureNoDir(pathElem ...string) string {
	dirname := path.Join(pathElem...)
	PanicMsgIfErr(os.RemoveAll(dirname), "Could not remove folder '%s'", dirname)

	return dirname
}

// WriteBytesToFile writes the given bytes to the file with the given path
func WriteBytesToFile(filename string, bytes []byte) {
	if filename != path.Base(filename) {
		EnsureDir(path.Dir(filename))
	}

	PanicMsgIfErr(os.WriteFile(filename, bytes, 0o644), "Could not write to file '%s'", filename)
}

// WriteStringToFile writes the given string to the file with the given path
func WriteStringToFile(filename string, content string, params ...any) {
	WriteBytesToFile(filename, []byte(fmt.Sprintf(content, params...)))
}

// WriteJsonObjToFile writes the given JSON object to the file with the given path
func WriteJsonObjToFile(filename string, obj any) {
	jsonBytes, errMarshal := json.MarshalIndent(obj, "", "\t")
	PanicMsgIfErr(errMarshal, "Could not JSON-marshal to file '%s'", filename)
	WriteBytesToFile(filename, jsonBytes)
}

// ReadFile reads the file with the given path and returns the bytes
func ReadFile(filename string, failIfNotExist bool) []byte {
	if !FileExists(filename) {
		PanicMsgIf(failIfNotExist, "File '%s' cannot be found!", filename)
		return nil
	}

	fileBytes, errRead := os.ReadFile(filename)
	PanicMsgIfErr(errRead, "Could not read file '%s'", filename)
	return fileBytes
}

// ReadFileFromJSON reads the file with the given path and unmarshals the JSON object
func ReadFileFromJSON[T any, Y *T](filename string, obj Y, failIfNotExist bool) Y {
	if fileBytes := ReadFile(filename, failIfNotExist); fileBytes != nil {
		PanicMsgIfErr(json.Unmarshal(fileBytes, obj), "Could not JSON-unmarshal file '%s'", filename)
	}
	return obj
}

// ReadFileFromYAML reads the file with the given path and unmarshals the YAML object
func ReadFileFromYAML[T any, Y *T](filename string, obj Y, failIfNotExist bool) Y {
	if fileBytes := ReadFile(filename, failIfNotExist); fileBytes != nil {
		PanicMsgIfErr(yaml.Unmarshal(fileBytes, obj), "Could not YAML-unmarshal file '%s'", filename)
	}
	return obj
}

// ReplaceInFile performs the given replacements in the file with the given path
func ReplaceInFile(filename string, replacements map[string]string) {
	fileContent := string(ReadFile(filename, true))
	for replace, by := range replacements {
		fileContent = strings.ReplaceAll(fileContent, replace, by)
	}
	WriteStringToFile(filename, "%s", fileContent)
}

// ReplaceInFolder performs the given replacements in the files within the given folder, with the given extension
// Example:
//
//	   ("src", ".go", map[string]{"foo":"bar", "toto":"titi"}) =>
//		  replaces "boo" with "bar", and "toto" with "titi", in all the *.go files found in "./src".
func ReplaceInFolder(folder string, extension string, replacements map[string]string) {
	for _, entry := range EnsureReadDir(folder) {
		if entry.IsDir() {
			ReplaceInFolder(path.Join(folder, entry.Name()), extension, replacements)
		} else if strings.HasSuffix(entry.Name(), extension) {
			ReplaceInFile(path.Join(folder, entry.Name()), replacements)
		}
	}
}
