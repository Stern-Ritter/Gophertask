package server

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env"

	config "github.com/Stern-Ritter/gophertask/internal/config/server"
	"github.com/Stern-Ritter/gophertask/internal/utils"
)

type jsonConfig struct {
	URL               string `json:"address,omitempty"`
	DatabaseDSN       string `json:"database_dsn,omitempty"`
	AuthenticationKey string `json:"authentication_key,omitempty"`
	TLSCertPath       string `json:"tls_cert,omitempty"`
	TLSKeyPath        string `json:"tls_key,omitempty"`
	ShutdownTimeout   int    `json:"shutdown_timeout,omitempty"`
	LoggerLvl         string `json:"logger_level,omitempty"`
}

func GetConfig(defaultCfg config.ServerConfig) (config.ServerConfig, error) {
	cfg := config.ServerConfig{}

	parseFlags(&cfg)

	err := env.Parse(&cfg)
	if err != nil {
		return cfg, err
	}

	cfgFile := strings.TrimSpace(cfg.ConfigFile)
	needParseJSONConfig := len(cfgFile) > 0
	if needParseJSONConfig {
		err = parseJSONConfig(&cfg, cfgFile)
		if err != nil {
			return cfg, err
		}
	}

	mergeDefaultConfig(&cfg, defaultCfg)

	trimStringVarsSpaces(&cfg)

	return cfg, nil
}

func parseFlags(cfg *config.ServerConfig) {
	flag.StringVar(&cfg.URL, "h", "", "address and port to run server in format <host>:<port>")
	flag.StringVar(&cfg.DatabaseDSN, "d", "", "database dsn")
	flag.StringVar(&cfg.AuthenticationKey, "a", "", "secret authentication key")
	flag.StringVar(&cfg.TLSCertPath, "tls-cert", "", "path to tls certificate")
	flag.StringVar(&cfg.TLSKeyPath, "tls-key", "", "path to tls key")
	flag.StringVar(&cfg.ConfigFile, "c", "", "path to json config file")
	flag.Parse()
}

func parseJSONConfig(cfg *config.ServerConfig, fPath string) error {
	data, err := os.ReadFile(fPath)
	if err != nil {
		return fmt.Errorf("read config file %s: %w", fPath, err)
	}

	jsonCfg := jsonConfig{}
	err = json.Unmarshal(data, &jsonCfg)
	if err != nil {
		return fmt.Errorf("parse config file %s: %w", fPath, err)
	}

	mergeJSONConfig(cfg, jsonCfg)
	return nil
}

func mergeJSONConfig(cfg *config.ServerConfig, jsonCfg jsonConfig) {
	cfg.URL = utils.Coalesce(cfg.URL, jsonCfg.URL)
	cfg.DatabaseDSN = utils.Coalesce(cfg.DatabaseDSN, jsonCfg.DatabaseDSN)
	cfg.AuthenticationKey = utils.Coalesce(cfg.AuthenticationKey, jsonCfg.AuthenticationKey)
	cfg.TLSCertPath = utils.Coalesce(cfg.TLSCertPath, jsonCfg.TLSCertPath)
	cfg.TLSKeyPath = utils.Coalesce(cfg.TLSKeyPath, jsonCfg.TLSKeyPath)
	cfg.ShutdownTimeout = utils.Coalesce(cfg.ShutdownTimeout, jsonCfg.ShutdownTimeout)
	cfg.LoggerLvl = utils.Coalesce(cfg.LoggerLvl, jsonCfg.LoggerLvl)
}

func mergeDefaultConfig(cfg *config.ServerConfig, defaultCfg config.ServerConfig) {
	cfg.URL = utils.Coalesce(cfg.URL, defaultCfg.URL)
	cfg.DatabaseDSN = utils.Coalesce(cfg.DatabaseDSN, defaultCfg.DatabaseDSN)
	cfg.AuthenticationKey = utils.Coalesce(cfg.AuthenticationKey, defaultCfg.AuthenticationKey)
	cfg.TLSCertPath = utils.Coalesce(cfg.TLSCertPath, defaultCfg.TLSCertPath)
	cfg.TLSKeyPath = utils.Coalesce(cfg.TLSKeyPath, defaultCfg.TLSKeyPath)
	cfg.ConfigFile = utils.Coalesce(cfg.ConfigFile, defaultCfg.ConfigFile)
	cfg.ShutdownTimeout = utils.Coalesce(cfg.ShutdownTimeout, defaultCfg.ShutdownTimeout)
	cfg.LoggerLvl = utils.Coalesce(cfg.LoggerLvl, defaultCfg.LoggerLvl)
}

func trimStringVarsSpaces(cfg *config.ServerConfig) {
	cfg.URL = strings.TrimSpace(cfg.URL)
	cfg.DatabaseDSN = strings.TrimSpace(cfg.DatabaseDSN)
	cfg.AuthenticationKey = strings.TrimSpace(cfg.AuthenticationKey)
	cfg.TLSCertPath = strings.TrimSpace(cfg.TLSCertPath)
	cfg.TLSKeyPath = strings.TrimSpace(cfg.TLSKeyPath)
	cfg.ConfigFile = strings.TrimSpace(cfg.ConfigFile)
	cfg.LoggerLvl = strings.TrimSpace(cfg.LoggerLvl)
}
