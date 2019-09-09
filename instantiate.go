package instantiate

import (
	"flag"
	"path"

	"github.com/spf13/viper"
)

// Viper pounce
func Viper() error {
	c := flag.String("c", path.Join(".", "config", "dev.json"), "specify config file")

	flag.Parse()

	viper.AutomaticEnv()
	viper.SetConfigFile(*c)

	return viper.ReadInConfig()
}
