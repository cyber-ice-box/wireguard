package config

import (
	"errors"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
)

type (
	Config struct {
		GRPC       GRPCConfig       `yaml:"grpc"`
		VPN        VPNConfig        `yaml:"vpn"`
		PostgresDB PostgresDBConfig `yaml:"postgresDB"`
	}

	VPNConfig struct {
		Port       string `yaml:"port" env:"vpn-port" env-default:"51820" env-description:"VPN port"`
		Address    string `yaml:"address" env:"vpn-address" env-default:"10.0.0.1" env-description:"VPN address"`
		PrivateKey string `yaml:"privateKey" env:"vpn-privateKey" env-description:"VPN privateKey"`
	}

	PostgresDBConfig struct {
		Endpoint string `yaml:"endpoint" env:"pg-endpoint" env-description:"Endpoint of Postgres"`
		Username string `yaml:"username" env:"pg-username" env-description:"Username of Postgres"`
		Password string `yaml:"password" env:"pg-password" env-description:"Password of Postgres"`
		DBName   string `yaml:"dbName" env:"pg-dbName" env-description:"Database of Postgres"`
		SSLMode  string `yaml:"sslMode" env:"pg-sslMode" env-default:"verify-full" env-description:"SSL mode of Postgres"`
	}

	GRPCConfig struct {
		Endpoint string     `yaml:"endpoint" env:"gprc-endpoint" env-default:"0.0.0.0:5454" env-description:"Endpoint of GRPC server"`
		TLS      TLSConfig  `yaml:"tls"`
		Auth     AuthConfig `yaml:"auth"`
	}

	TLSConfig struct {
		Enabled  bool   `yaml:"enabled" env:"gprc-tls-enabled" env-default:"false" env-description:"Enabled TLS of GRPC server"`
		CertFile string `yaml:"certFile" env:"gprc-tls-certFile" env-default:"" env-description:"CertFile of GRPC server"`
		CertKey  string `yaml:"certKey" env:"gprc-tls-certKey" env-default:"" env-description:"CertKey of GRPC server"`
		CAFile   string `yaml:"caFile" env:"gprc-tls-caFile" env-default:"" env-description:"CaFile of GRPC server"`
	}

	AuthConfig struct {
		AuthKey string `yaml:"authKey" env:"gprc-authKey" env-description:"Auth key of GRPC server"`
		SignKey string `yaml:"signKey" env:"gprc-signKey" env-description:"Sign key of GRPC server"`
	}
)

type Valid = bool

func GetConfig(path *string) (*Config, Valid) {
	log.Info().Msg("Reading wireguard configuration")
	instance := &Config{}
	header := "Config variables:"
	help, _ := cleanenv.GetDescription(instance, &header)

	var err error

	if path != nil {
		err = cleanenv.ReadConfig(*path, instance)
	} else {
		err = cleanenv.ReadEnv(instance)
	}

	if err != nil {
		log.Error().Err(err).Msg("See the help bellow")
		fmt.Println(help)
		return nil, false
	}

	err = validateConfig(instance)

	if err != nil {
		log.Error().Err(err).Msg("See the help bellow")
		fmt.Println(help)
		return nil, false
	}

	return instance, true
}

func validateConfig(cfg *Config) error {
	if cfg.GRPC.Auth.AuthKey == "" || cfg.GRPC.Auth.SignKey == "" {
		return errors.New("sign and auth keys must be provided")
	}

	if cfg.GRPC.TLS.Enabled {
		if cfg.GRPC.TLS.CertFile == "" || cfg.GRPC.TLS.CertKey == "" || cfg.GRPC.TLS.CAFile == "" {
			return errors.New("if TLS is enabled, all certifications must be provided")
		}
	}

	return nil
}
