package config

type Config struct {
	HTTPServer HTTPServer
	Database   Database
}

type HTTPServer struct {
	HostName string `long:"server hostname" env:"HTTPSERVER_HOSTNAME" default:"localhost" description:"host name"`
	Port     string `long:"server port, default is 8080" env:"HTTPSERVER_PORT" default:"8080" description:"host port"`
}

type Database struct {
	Hostname string `long:"DB hostname" env:"DATABASE_HOSTNAME"`
	Port     string `long:"DB port" env:"DATABASE_PORT"`

	Name     string `long:"DB name" env:"DATABASE_NAME"`
	Username string `long:"DB username" env:"DATABASE_USERNAME"`
	Password string `long:"DB password" env:"DATABASE_PASSWORD"`
}
