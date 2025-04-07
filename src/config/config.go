package config

type Config struct {
	Port      int    `yaml:"port"`
	StaticDir string `yaml:"static_dir"`
}
