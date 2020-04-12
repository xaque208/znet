package znet

import (
	"github.com/xaque208/znet/internal/agent"
	"github.com/xaque208/znet/internal/astro"
	"github.com/xaque208/znet/internal/gitwatch"
	"github.com/xaque208/znet/internal/lights"
	"github.com/xaque208/znet/internal/timer"
)

// Config stores the items that are required to configure this project.
type Config struct {
	Agent        agent.Config        `yaml:"agent,omitempty"`
	Astro        astro.Config        `yaml:"astro,omitempty"`
	Environments []EnvironmentConfig `yaml:"environments,omitempty"`
	GitWatch     gitwatch.Config     `yaml:"gitwatch,omitempty"`
	HTTP         HTTPConfig          `yaml:"http,omitempty"`
	Junos        JunosConfig         `yaml:"junos,omitempty"`
	LDAP         LDAPConfig          `yaml:"ldap,omitempty"`
	Lights       lights.Config       `yaml:"lights,omitempty"`
	Nats         NatsConfig          `yaml:"nats,omitempty"`
	Rooms        []lights.Room       `yaml:"rooms,omitempty"`
	RPC          RPCConfig           `yaml:"rpc,omitempty"`
	Timer        timer.Config        `yaml:"timer,omitempty"`
	Vault        VaultConfig         `yaml:"vault,omitempty"`
}

// EnvironmentConfig is the environment configuration.
type EnvironmentConfig struct {
	Name         string   `yaml:"name,omitempty"`
	SecretValues []string `yaml:"secret_values,omitempty"`
}

// NatsConfig is the NATs configuration.
type NatsConfig struct {
	URL   string
	Topic string
}

// JunosConfig is the configuration for Junos devices.
type JunosConfig struct {
	Hosts      []string
	Username   string
	PrivateKey string
}

// HTTPConfig is the configuration for the listening HTTP server.
type HTTPConfig struct {
	ListenAddress string
}

// RPCConfig is the configuration for the RPC client and server.
type RPCConfig struct {
	ListenAddress string
	ServerAddress string
}

// LDAPConfig is the client configuration for LDAP.
type LDAPConfig struct {
	BaseDN    string `yaml:"basedn"`
	BindDN    string `yaml:"binddn"`
	BindPW    string `yaml:"bindpw"`
	Host      string `yaml:"host"`
	UnknownDN string `yaml:"unknowndn"`
}

// VaultConfig is the client configuration for Vault.
type VaultConfig struct {
	Host      string
	TokenPath string `yaml:"token_path,omitempty"`
	VaultPath string `yaml:"vault_path,omitempty"`
}
