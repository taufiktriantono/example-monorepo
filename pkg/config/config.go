package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/taufiktriantono/api-first-monorepo/pkg/security"
	"go.uber.org/fx"
)

var (
	config = viper.New()
)

var Module = fx.Module("config", fx.Provide(LoadConfig))

type Config struct {
	RootDomain   string `mapstructure:"ROOT_DOMAIN"`
	AppEnv       string `mapstructure:"APP_ENV"`
	AppName      string `mapstructure:"APP_NAME"`
	AppVersion   string `mapstructure:"APP_VERSION"`
	AppNamespace string `mapstructure:"APP_NAMESPACE"`
	Pyroscope    struct {
		Addr string `mapstructure:"ADDR"`
	} `mapstructure:"PYROSCOPE"`
	Server struct {
		Addr           string        `mapstructure:"ADDR"`
		ReadTimeout    time.Duration `mapstructure:"READ_TIMEOUT"`
		WriteTimeout   time.Duration `mapstructure:"WRITE_TIMEOUT"`
		IdleTimeout    time.Duration `mapstructure:"IDLE_TIMEOUT"`
		UseUnixSocket  bool          `mapstructure:"USE_UNIX_SOCKET"`
		UnixSocketPath string        `mapstructure:"UNIX_SOCKET_PATH"`
		TLS            struct {
			Enable   bool   `mapstructure:"ENABLE"`
			CertPath string `mapstructure:"CERT_PATH"`
			KeyPath  string `mapstructure:"KEY_PATH"`
		} `mapstructure:"TLS"`
	} `mapstructure:"HTTP_SERVER"`
	Session struct {
		Type   string `mapstructure:"TYPE"`
		Name   string `mapstructure:"NAME"`
		Secret string `mapstructure:"SECRET"`
	} `mapstructure:"SESSION"`
	Database struct {
		Type     string `mapstructure:"TYPE"`
		Host     string `mapstructure:"HOST"`
		Port     string `mapstructure:"PORT"`
		DBNAME   string `mapstructure:"DBNAME"`
		User     string `mapstructure:"USER"`
		Password string `mapstructure:"PASSWORD"`
		SSLMode  string `mapstructure:"SSLMODE"`
		Timezone string `mapstructure:"TIMEZONE"`
	} `mapstructure:"DATABASE"`
	Redis struct {
		Addr     string `mapstructure:"ADDR"`
		Password string `mapstructure:"PASSWORD"`
		DB       int    `mapstructure:"DB"`
	} `mapstructure:"REDIS"`
	Kafka struct {
		Addrs string `mapstructure:"ADDR"`
	}
	ConnectionPool struct {
		MaxIdleConn     int           `mapstructure:"MAX_IDLE_CONN"`
		MaxOpenConn     int           `mapstructure:"MAX_OPEN_CONN"`
		MaxOpenConns    int           `mapstructure:"MAX_OPEN_CONNS"`
		ConnMaxLifetime time.Duration `mapstructure:"CONN_MAX_LIFETIME"`
		ConnMaxIdleTime time.Duration `mapstructure:"CONN_MAX_IDLE_TIME"`
	} `mapstructure:"CONNECTION_POOL"`
	AccessControl struct {
		Model  string `mapstructure:"MODEL"`
		Policy string `mapstructure:"POLICY"`
	} `mapstructure:"ACCESS_CONTROL"`
	SecretAES string `mapstructure:"SECRET_AES"`
}

func GetString(target string) string {
	return config.GetString(target)
}

func GetStringMap(target string) map[string]any {
	return config.GetStringMap(target)
}

func GetStringMapString(target string) map[string]string {
	return config.GetStringMapString(target)
}

func GetStringMapStringSlice(target string) map[string][]string {
	return config.GetStringMapStringSlice(target)
}

func GetInt(target string) int {
	return config.GetInt(target)
}

func GetInt32(target string) int32 {
	return config.GetInt32(target)
}

func GetInt64(target string) int64 {
	return config.GetInt64(target)
}

func LoadConfig(services ...string) (*Config, error) {
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./configs")

	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}

	secret, err := security.ValidateBase64Secret(config.GetString("SECRET_AES"))
	if err != nil {
		return nil, err
	}

	if err := decryptRecursive("", secret, config.AllSettings()); err != nil {
		return nil, err
	}

	var cfg Config
	if err := config.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	for _, service := range services {
		newCfg := viper.New()
		newCfg.SetConfigName(service)
		newCfg.SetConfigType("yaml")
		newCfg.AddConfigPath("./conf")

		if err := newCfg.ReadInConfig(); err != nil {
			return nil, err
		}

		if err := decryptRecursive("", secret, newCfg.AllSettings()); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}

func decryptRecursive(prefix string, secret []byte, data map[string]interface{}) error {
	for k, v := range data {
		fullKey := buildFullKey(prefix, k)

		switch val := v.(type) {
		case string:
			if err := handleEncryptedString(fullKey, secret, val); err != nil {
				return err
			}
		case map[string]interface{}:
			if err := decryptRecursive(fullKey, secret, val); err != nil {
				return err
			}
		case map[interface{}]interface{}:
			converted := convertMap(val)
			if err := decryptRecursive(fullKey, secret, converted); err != nil {
				return err
			}
		default:
			// skip non-string/map types
		}
	}
	return nil
}

func buildFullKey(prefix, key string) string {
	if prefix == "" {
		return key
	}
	return prefix + "." + key
}

func handleEncryptedString(fullKey string, secret []byte, val string) error {
	if strings.HasPrefix(val, "ENC(") && strings.HasSuffix(val, ")") {
		encValue := val[4 : len(val)-1]
		decrypted, err := security.Decrypt(encValue, []byte(secret))
		if err != nil {
			return fmt.Errorf("failed to decrypt key %s: %w", fullKey, err)
		}
		viper.Set(fullKey, decrypted)
	}
	return nil
}

func convertMap(input map[interface{}]interface{}) map[string]interface{} {
	converted := make(map[string]interface{})
	for ik, iv := range input {
		if ks, ok := ik.(string); ok {
			converted[ks] = iv
		}
	}
	return converted
}
