package configloader

import (
	"fmt"
	"os"
)

type SQLServerConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Config struct {
	SQLServer SQLServerConfig `json:"sqlserver"`
}

// LoadConfig reads the config file
func LoadConfig() (Config, error) {
	config := Config{
		SQLServer: SQLServerConfig{
			Host:     os.Getenv("SQLSERVER_HOST"),
			Port:     os.Getenv("SQLSERVER_PORT"),
			User:     os.Getenv("SQLSERVER_USER"),
			Password: os.Getenv("SQLSERVER_PASSWORD"),
			Database: os.Getenv("SQLSERVER_DATABASE"),
		},
	}

	if config.SQLServer.Host == "" || config.SQLServer.Port == "" || config.SQLServer.User == "" || config.SQLServer.Password == "" || config.SQLServer.Database == "" {
		return Config{}, fmt.Errorf("missing required environment variables")
	}

	return config, nil
	// // Get the absolute path of the current Go file (main.go)
	// _, currentFile, _, _ := runtime.Caller(0)
	// dir := filepath.Dir(currentFile)

	// // Read the contents of the directory where the current file is located
	// files, err := os.ReadDir(dir)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Print the contents of the directory
	// fmt.Println("Contents of the current directory:", dir)
	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }

	// file, err := os.Open("../config.json")
	// if err != nil {
	// 	return Config{}, err
	// }
	// defer file.Close()

	// var config Config
	// decoder := json.NewDecoder(file)
	// err = decoder.Decode(&config)
	// return config, err
}
