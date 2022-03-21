package flags

//define all flags constants
const (
	EnvKey                     = "env"
	PortKey                    = "port"
	BaseConfigPathKey          = "base-config-path"
	BaseConfigPathUsage        = "path to folder that stores your configurations"
	EnvUsage                   = "application.yml env"
	PortUsage                  = "application.yml port"
	BaseConfigPathDefaultValue = "."
	EnvDefaultValue            = "DEV"
	PortDefaultValue           = 8080
)
