package pconfig

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Port         int    `yaml:"port" json:"port"`
	DbUrl        string `yaml:"db_url" json:"dburl"`
	JaegerUrl    string `yaml:"jaeger_url" json:"jaegerurl"`
	SentryUrl    string `yaml:"sentry_url" json:"sentryurl"`
	KafkaBrocker string `yaml:"kafka_broker" json:"kafkabroker"`
	AppId        int    `yaml:"some_app_id" json:"someappid"`
	AppKey       string `yaml:"some_app_key" json:"someappkey"`
}

// func main() {
// 	cfg, err := Configure()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(cfg.Port)
// }

func Configure() (Config, error) {
	var cfg = Config{}
	var filename = ""
	flag.IntVar(&cfg.Port, "port", 8080, "port 0-65535")
	flag.StringVar(&cfg.DbUrl, "dburl", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "database url")
	flag.StringVar(&cfg.JaegerUrl, "jaeger", "http://jaeger:16686", "jaeger url")
	flag.StringVar(&cfg.SentryUrl, "sentry", "http://sentry:9000", "sentry url")
	flag.StringVar(&cfg.KafkaBrocker, "kafka", "http://kafka:9092", "kafka broker")
	flag.IntVar(&cfg.AppId, "appid", 10000001, "app id")
	flag.StringVar(&cfg.AppKey, "appkey", "Dagsdhgeyr273rrn23jurh2u3rh", "app key")
	flag.StringVar(&filename, "filename", "", "json config file")
	flag.Parse()

	if len(filename) > 5 {
		if !(filepath.Ext(filename) == ".json" || filepath.Ext(filename) == ".yaml") {
			return Config{}, fmt.Errorf("file format must be json or yaml. %s", filename)
		}

		file, err := ioutil.ReadFile(filename)
		if err != nil {
			return Config{}, err
		}

		if filepath.Ext(filename) == ".json" {
			err = json.Unmarshal(file, &cfg)
			if err != nil {
				return Config{}, fmt.Errorf("invalid format configuration file: %v", err)
			}
			fmt.Println(cfg.Port)
		} else if filepath.Ext(filename) == ".yaml" {
			err = yaml.Unmarshal(file, &cfg)
			if err != nil {
				return Config{}, fmt.Errorf("invalid format configuration file: %v", err)
			}
		}
	}
	cfg, err := validate(cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func validate(cfg Config) (Config, error) {
	if cfg.Port > 65535 || cfg.Port < 0 {
		return Config{}, fmt.Errorf("incorrect Port")
	}
	if !isUrl(cfg.DbUrl) {
		return Config{}, fmt.Errorf("incorrect DbUrl")
	}
	if !isUrl(cfg.JaegerUrl) {
		return Config{}, fmt.Errorf("incorrect jaegerUrl")
	}
	if !isUrl(cfg.SentryUrl) {
		return Config{}, fmt.Errorf("incorrect sentryUrl")
	}
	if !isUrl(cfg.KafkaBrocker) {
		return Config{}, fmt.Errorf("incorrect kafkaBrocker")
	}
	if cfg.AppId < 1000000 {
		return Config{}, fmt.Errorf("incorrect appId")
	}
	return cfg, nil
}

func isUrl(str string) bool {
	parsedUrl, err := url.Parse(str)
	return err == nil && parsedUrl.Scheme != "" && parsedUrl.Host != ""
}
