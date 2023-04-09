package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/viper"
)

//go:embed defaults.yaml
var defaultConfig []byte

type Config struct {
	raw *viper.Viper

	APIServer APIServerCfg `yaml:"api_server" mapstructure:"api_server"`
	Database  DatabaseCfg  `yaml:"database" mapstructure:"database"`
	Job       JobCfg       `yaml:"job" mapstructure:"job"`
}

func NewConfig(envPrefix string) *Config {
	raw := viper.New()
	raw.SetConfigType("yaml")
	raw.SetEnvPrefix(envPrefix)
	raw.AutomaticEnv()
	raw.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	if err := raw.MergeConfig(bytes.NewBuffer(defaultConfig)); err != nil {
		panic("cannot load default config")
	}

	c := &Config{
		raw: raw,
	}

	if err := c.mapConfig(); err != nil {
		panic("cannot map default config")
	}
	return c
}

func (c *Config) mapConfig() error {
	if err := c.raw.Unmarshal(&c); err != nil {
		return err
	}

	return nil
}

func (c *Config) mergeConfig(r io.Reader) error {
	if err := c.raw.MergeConfig(r); err != nil {
		return err
	}
	if err := c.mapConfig(); err != nil {
		return err
	}
	return nil
}

func (c *Config) ReadConfigFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file filename=%s, err=%s", filename, err)
	}
	if err := c.mergeConfig(f); err != nil {
		return fmt.Errorf("failed to read config filename=%s, err=%s", filename, err)
	}

	return nil
}

type APIServerCfg struct {
	RunMode string `yaml:"run_mode" mapstructure:"run_mode"`
	Listen  string `yaml:"listen" mapstructure:"listen"`
}

type DatabaseCfg struct {
	Debug    bool              `yaml:"debug" mapstructure:"debug"`
	Type     string            `yaml:"type" mapstructure:"type"`
	Host     string            `yaml:"host" mapstructure:"host"`
	Port     int64             `yaml:"port" mapstructure:"port"`
	DBName   string            `yaml:"dbname" mapstructure:"dbname"`
	User     string            `yaml:"user" mapstructure:"user"`
	Password string            `yaml:"password" mapstructure:"password"`
	Opts     map[string]string `yaml:"opts" mapstructure:"opts"`
	Pool     struct {
		ConnMaxIdleTime int `yaml:"conn_max_idle_time" mapstructure:"conn_max_idle_time"`
		ConnMaxLifeTime int `yaml:"conn_max_life_time" mapstructure:"conn_max_life_time"`
		MaxOpenConns    int `yaml:"max_open_conns" mapstructure:"max_open_conns"`
		MaxIdleConns    int `yaml:"max_idle_conns" mapstructure:"max_idle_conns"`
	}
}

type JobCfg struct {
	WorkerNum             int   `yaml:"worker_num" mapstructure:"worker_num"`
	PollInterval          int64 `yaml:"poll_interval" mapstructure:"poll_interval"`
	PollLimit             int   `yaml:"poll_limit" mapstructure:"poll_limit"`
	RetriesOnTimeoutLimit int   `yaml:"retries_on_timeout_limit" mapstructure:"retries_on_timeout_limit"`
	RetriesOnErrorLimit   int   `yaml:"retries_on_error_limit" mapstructure:"retries_on_error_limit"`
}
