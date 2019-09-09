package instantiate

import (
	"flag"
	"path"

	"github.com/spf13/viper"
)

var service string

// Viper pounce
func Viper() error {
	c := flag.String("c", path.Join(".", "config", "dev.json"), "specify config file")

	flag.Parse()

	viper.AutomaticEnv()
	viper.SetConfigFile(*c)

	return viper.ReadInConfig()
}

// func init() {
// 	// var service string

// 	if service != "" {
// 		fmt.Println("BINGO: I will write to:", service)
// 	} else {
// 		fmt.Println("wa waw")
// 	}
// }
