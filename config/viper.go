package config

import (
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var defaultConfigFile = "./config.yaml"

func readFromFile(v *viper.Viper) {
	// default config path
	configPath := v.GetString("config-path")
	if configPath == "" {
		configPath = defaultConfigFile
	}
	// if file doesn't exist just continue
	_, err := os.Stat(configPath)
	if err != nil {
		return
	}
	// set config type from file extension
	ext := filepath.Ext(configPath)[1:]
	v.SetConfigType(ext)
	// set file path
	v.SetConfigFile(configPath)
	// try read config or fatal
	err = v.ReadInConfig()
	if err != nil {
		logrus.Fatalf("failed to read config (%s): %+v", configPath, err)
	}
}

func NewViper() *viper.Viper {
	v := viper.New()
	// enable parse from environment variable
	v.AutomaticEnv()
	// replace "." and "-" with "_" for envs
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	// try read config where its possible
	readFromFile(v)
	return v
}

type ViperConfigProvider struct {
	v *viper.Viper
}

func NewViperConfigProvider(v *viper.Viper) IConfigProvider {
	return &ViperConfigProvider{v: v}
}

func (cp *ViperConfigProvider) Parse(config IConfig) error {
	return config.ParseFromViper(cp.v)
}

func ViperGetOrDefault(v *viper.Viper, key string, defaultValue string) string {
	if v := v.Get(key); v == nil {
		return defaultValue
	}
	return v.GetString(key)
}

func ViperGetOrDefaultUint64(v *viper.Viper, key string, defaultValue uint64) uint64 {
	if v := v.Get(key); v == nil {
		return defaultValue
	}
	return v.GetUint64(key)
}

func ViperGetOrDefaultUint32(v *viper.Viper, key string, defaultValue uint32) uint32 {
	if v := v.Get(key); v == nil {
		return defaultValue
	}
	return v.GetUint32(key)
}

func ViperGetOrDefaultFloat64(v *viper.Viper, key string, defaultValue float64) float64 {
	if v := v.Get(key); v == nil {
		return defaultValue
	}
	return v.GetFloat64(key)
}

func ViperGetOrDefaultBool(v *viper.Viper, key string, defaultValue bool) bool {
	if v := v.Get(key); v == nil {
		return defaultValue
	}
	return v.GetBool(key)
}

func ViperGetOrDefaultDuration(v *viper.Viper, key string, defaultValue time.Duration) time.Duration {
	if v := v.Get(key); v == nil {
		return defaultValue
	}
	return v.GetDuration(key)
}

func ViperGetStringArray(v *viper.Viper, key string, defaultValue []string) []string {
	if v := v.Get(key); v == nil {
		return defaultValue
	}
	return v.GetStringSlice(key)
}

type rawConfigProvider struct {
	src IConfig
}

func (p *rawConfigProvider) Parse(dst IConfig) error {
	buf := bytes.Buffer{}
	err := gob.NewEncoder(&buf).Encode(p.src)
	if err != nil {
		return err
	}
	err = gob.NewDecoder(&buf).Decode(dst)
	if err != nil {
		return err
	}
	return nil
}

func NewRawConfig(config IConfig) IConfigProvider {
	return &rawConfigProvider{src: config}
}
