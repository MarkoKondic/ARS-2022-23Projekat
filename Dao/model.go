package Dao

// swagger:model Config
type Config struct {

	// Id of the config
	// in: string
	Id string `json:"id"`

	Labels string `json:"labels"`

	// Version of the config
	// in: string
	Version string `json:"version"`

	Entries map[string]string `json:"entries"`
}

// swagger: model ConfigGroup
type ConfigGroup struct {

	// Id of the configGroup
	// in: string
	Id string `json:"id"`

	// Configs of the configGroup
	// in map[string]map[string]*Config
	Configs []*Config `json:"configs"`
}
