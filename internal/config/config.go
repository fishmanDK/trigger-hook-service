package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env            string         `yaml:"env" env-default:"local"`
	Postgres       PostgresConfig `yaml:"postgres" env-required:"true"`
	GRPC           GRPCConfig     `yaml:"grpc"`
	MigrationsPath string
	TokenTTL       time.Duration `yaml:"token_ttl" env-default:"1h"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type PostgresConfig struct {
	PgUser     string `yaml:"pg_user"`
	PgDatabase string `yaml:"pg_database"`
	PgHost     string `yaml:"pg_host"`
	PgPort     string `yaml:"pg_port"`
	PgSslmode  string `yaml:"pg_sslmode"`
	PgPassword string `yaml:"pg_password"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

func (pg PostgresConfig) String() string {
	s := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", pg.PgHost, pg.PgPort, pg.PgUser, pg.PgDatabase, pg.PgPassword, pg.PgSslmode)
	return s
}
