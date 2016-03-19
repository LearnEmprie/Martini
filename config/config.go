package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"bytes"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func GetParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func GetConfig()  {

	var buffer bytes.Buffer

	config_path := GetCurrentDirectory()

	buffer.WriteString(config_path)
	buffer.WriteString("/config.go")

	file, err := os.Open(buffer.String()) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("read %d bytes: %q\n", count, data[:count])
}