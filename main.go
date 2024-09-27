package main

import (
	_ "embed"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

type Config struct {
	S3 struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Url      string `yaml:"url"`
	}
	Db struct {
		Driver   string `yaml:"driver"`
		Filepath string `yaml:"filepath"`
	}
}

//go:embed config.yaml
var configBytes []byte

func main() {

	config := Config{}
	yaml.Unmarshal(configBytes, &config)
	fmt.Printf("config = %+v\n", config)

	s3Client := GetS3Client(config)
	s3Client.UploadPackage([]byte{123, 123, 123, 123}, "bucket", "config")

	db := CreateDbRepo(config)
	db.CreatePackageTable()

}
