package znet

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// GetEnvironmentConfig receives a slice of environment configurations and returns the one that matches the given name.
func GetEnvironmentConfig(environments []EnvironmentConfig, envName string) (EnvironmentConfig, error) {

	log.Debugf("looking for environment %s: %+v", envName, environments)

	for _, e := range environments {
		if e.Name == envName {
			return e, nil
		}
	}

	return EnvironmentConfig{}, fmt.Errorf("no environment with name %s found", envName)
}

// LoadEnvironment reads reads environment variables out of vault for return.
func LoadEnvironment(config VaultConfig, e EnvironmentConfig) (map[string]string, error) {

	environment := make(map[string]string)
	if config.Host == "" || config.VaultPath == "" {
		return map[string]string{}, fmt.Errorf("incomplete vault configuration, unable to load Environment")
	}

	s, err := NewSecretClient(config)
	if err != nil {
		return map[string]string{}, err
	}

	for _, k := range e.SecretValues {
		path := fmt.Sprintf("%s/%s", config.VaultPath, k)
		// log.Debugf("Reading vault path: %s", path)
		secret, err := s.Logical().Read(path)
		if err != nil {
			log.Error(err)
		}

		if secret != nil {
			environment[k] = secret.Data["value"].(string)
		}
	}

	return environment, nil
}
