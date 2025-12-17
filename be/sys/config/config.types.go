package config

// config struct all configuration in yaml file will be loaded to this struct
type Config struct {
	App    AppConfig    `yaml:"app"`
	Log    LogConfig    `yaml:"log"`
	Server ServerConfig `yaml:"server"`
	Db     DbConfig     `yaml:"db"`
	Cache  CacheConfig  `yaml:"cache"`
	AppDir string       `yaml:"-"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type LogConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DbConfig struct {
	Dsn string `yaml:"dsn"`
}

type CacheConfig struct {
	Enabled  bool              `yaml:"enabled"`
	Provider string            `yaml:"provider"`
	TTL      int               `yaml:"ttl"`
	Server   CacheServerConfig `yaml:"server"`
	User     string            `yaml:"user"`
	Password string            `yaml:"password"`
}

type CacheServerConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
