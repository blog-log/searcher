package config

import "github.com/kelseyhightower/envconfig"

func Load() (*Algolia, error) {
	var cfg Algolia
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
