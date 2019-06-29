package config

type (
	// Config is the struct representing informations contained
	// in the given config file.
	// It is used to parse data into go struct.
	Config struct {
		Plugins []*Plugin `mapstructure:"plugins"`
	}

	// Plugin represents informations about a plugin.
	Plugin struct {
		Name string `mapstructure:"name"`
		Url  string `mapstructure:"url"`
	}
)
