package configs

var client Client

// Init is used to initialize the configs
func Init(directory string, configNames ...string) error {
	var err error
	client, err = New(Options{
		Provider: FileBased,
		Params: map[string]interface{}{
			"configsDirectory": directory,
			"configNames":      configNames,
			"configType":       "yaml",
		},
	})
	if err != nil {
		return err
	}
	return nil
}

// Get is used to get the config client
func Get() Client {
	return client
}
