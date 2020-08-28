package config

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestParseConfig(t *testing.T) {
	pretty("configs", viper.ConfigFileUsed())
	c, err := ParseConfig()
	if err != nil {
		t.Error("parse err", err)
		return
	}
	pretty("c", c)
}

func pretty(prefix string, a interface{}) {
	b, _ := json.MarshalIndent(a, "", "  ")
	fmt.Printf("%v: %s\n", prefix, string(b))
}
