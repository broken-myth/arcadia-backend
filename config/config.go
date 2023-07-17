package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Name     string
	Username string
	Password string
	Port     uint
	Host     string
	LogFile  string
}

type RedisConfig struct {
	Host     string
	Port     uint
	Password string
}

type LoggerConfig struct {
	FileName string
	MaxSize  int
	Level    string
}

type AuthConfig struct {
	// OAuth2
	OAuth2Key    string
	OAuth2Secret string
	RedirectURL  string

	// JWT
	JWTSecret         string
	TokenHourLifeSpan string

	// Admin
	AdminHeader string
}

type Config struct {
	// DOCKER, DEV
	AppEnv string

	// Server Config
	Host string
	Port uint

	//Allowed Origins
	AllowedOrigins string

	// OpenAPI docs URL
	SwaggerURL string

	// Logger Config
	Log LoggerConfig

	// Database Config
	Db DatabaseConfig

	RedisDb RedisConfig

	// Auth Config
	Auth AuthConfig

	// Rate limit Config
	Ratelimit uint
}

// All configurations
var allConfigurations = struct {

	// Configuration for app environment : DEV
	Dev Config

	// Configuration for app environment : DOCKER
	Docker Config
}{}

// Current config
var currentConfig Config

// Function to get current config
func GetConfig() Config {
	return currentConfig
}

// Initialize config
func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(color.RedString("Error loading .env file"))
	}

	appEnv := os.Getenv("APP_ENV")

	// Load JSON config
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(color.RedString("Error loading config.json file"))
		panic(err)
	}
	defer configFile.Close()

	// Parse JSON config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&allConfigurations)

	if err != nil {
		fmt.Println(color.RedString("Error decoding config.json file"))
	}

	// Set current config
	switch appEnv {
	case "DEV":
		currentConfig = allConfigurations.Dev
	case "DOCKER":
		currentConfig = allConfigurations.Docker
	default:
		panic(color.RedString("Error setting current config"))
	}

	fmt.Print("\n")
	fmt.Println(color.GreenString("Config loaded successfully"))
}
