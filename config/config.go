package config

import "gopkg.in/yaml.v3"

type LanguageConfig struct {
	Name   string            `yaml:"name"`
	Glyphs map[string]string `yaml:"glyphs"`
}

func NewLanguageConfigFromBytes(bytes []byte) (*LanguageConfig, error) {
	config := &LanguageConfig{}
	err := yaml.Unmarshal(bytes, &config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
