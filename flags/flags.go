package flags

import (
	flag "github.com/spf13/pflag"
)

var (
	env            = flag.String(EnvKey, EnvDefaultValue, EnvUsage)
	port           = flag.Int(PortKey, PortDefaultValue, PortUsage)
	baseConfigPath = flag.String(BaseConfigPathKey, BaseConfigPathDefaultValue,
		BaseConfigPathUsage)
)

func init() {
	flag.Parse()
}

// Env is the application.yml runtime environment
func Env() string {
	return *env
}

// Port is the application.yml port number where the process will be started
func Port() int {
	return *port
}

// BaseConfigPath is the path that holds the configuration files
func BaseConfigPath() string {
	return *baseConfigPath
}
