package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config object
type Config struct {
	name        string
	version     string
	mapsEnabled bool
	mapsKey     string
}

var config *Config

// Load config from file
func Load() {
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	viper.SetDefault("application.name", "ifsc")
	viper.SetDefault("application.version", "NotDefined")
	viper.SetDefault("maps.enabled", false)

	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file %s was edited, reloading config\n", e.Name)
		readLatestConfig()
	})

	readLatestConfig()

}

func readLatestConfig() {
	config = &Config{
		name:        viper.GetString("application.name"),
		version:     viper.GetString("application.version"),
		mapsEnabled: viper.GetBool("maps.enabled"),
		mapsKey:     viper.GetString("maps.key"),
	}

}

// Application : Exporting configuration
func Application() *Config {
	return config
}

// Name : Exporting Name
func Name() string {
	return config.name
}

// Version : Export application version
func Version() string {
	return config.version
}

// MapsEnabled : Lookup to google maps enabled?
func MapsEnabled() bool {
	return config.mapsEnabled
}

// MapsKey : Google Maps key
func MapsKey() string {
	return config.mapsKey
}
