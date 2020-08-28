package config

import (
	"github.com/spf13/viper"
	"k8s.io/klog"
)

type Config struct {
	Creds []struct {
		Name    string `yaml:"name"`
		User    string `yaml:"user"`
		Pass    string `yaml:"pass"`
		Keypath string `yaml:"keypath"`
	} `yaml:"creds"`
	Hosts []struct {
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Cred string `yaml:"cred"`
		Port string `yaml:"port"`
	} `yaml:"hosts"`
}

func init() {
	// setting config properties
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // default type
	viper.AddConfigPath("/etc/goterm/")
	viper.AddConfigPath("$HOME/.goterm")
	viper.AddConfigPath(".")
}

func ParseConfig() (*Config, error) {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return nil, err
	}
	klog.V(2).Info("read from config: ", viper.ConfigFileUsed())
	c := &Config{}
	err = viper.Unmarshal(c)
	return c, err
}

func (c *Config) GetHost(name string) (host, port, cred string) {
	for _, v := range c.Hosts {
		if v.Name == name {
			return v.Host, v.Port, v.Cred
		}
	}
	return
}

// if name empty, get from default cred
func (c *Config) GetCred(name string) (user, pass, keypath string) {
	if len(name) == 0 {
		name = "default"
	}
	for _, v := range c.Creds {
		if v.Name == name {
			return v.User, v.Pass, v.Keypath
		}
	}
	return
}
