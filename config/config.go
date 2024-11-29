package config

import (
	"strings"

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
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Cred     string `yaml:"cred"`
		Port     string `yaml:"port"`
		Label    string `yaml:"label"`
		InitCmds string `yaml:"initcmds"`
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

func (c *Config) GetHost(expr string) (host, port, cred, cmds string) {
	klog.V(2).Infof("checking host: %v", expr)
	for _, v := range c.Hosts {
		if v.Name == expr || strings.HasSuffix(v.Host, expr) {
			klog.V(2).Infof("got cred: %+v", v)
			return v.Host, v.Port, v.Cred, v.InitCmds
		}
	}
	return expr, "22", "", ""
}

// if name empty, get from default cred
func (c *Config) GetCred(name string) (user, pass, keypath string) {
	klog.V(2).Infof("checking cred: %v", name)
	for _, v := range c.Creds {
		if v.Name == name {
			klog.V(2).Infof("found cred for: %v", name)
			return v.User, v.Pass, v.Keypath
		}
	}
	klog.V(2).Infof("cred not found for: %v", name)
	return
}
