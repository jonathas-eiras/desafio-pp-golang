package config

type Application struct {
	Version string  `yaml:"version"`
	Name    string  `yaml:"name"`
	Service Service `yaml:"service"`
}

type Service struct {
	Users Users `yaml:"users"`
}
type Users struct {
	BasePath string `yaml:"basePath"`
	Token    string `yaml:"token"`
}
