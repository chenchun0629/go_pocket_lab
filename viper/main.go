package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

// https://www.liwenzhou.com/posts/Go/viper_tutorial/
// https://darjun.github.io/2020/01/18/godailylib/viper/
// https://tkstorm.com/posts-list/programming/go/library/viper/

type ServeConfig struct {
	Protocols []string
	Ports     []string
	Timeout   string
}

func (c *ServeConfig) Load() {
	err := viper.Sub("data.server").Unmarshal(c)
	fmt.Printf("%#v, %#v", c, err)
}

func main() {

	var sc = new(ServeConfig)
	//viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)

		fmt.Printf("changed sc: %#v \n ", sc)
		sc.Load()
		fmt.Printf("changed sc: %#v \n ", sc)
	})

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read configs failed: %v", err)
	}

	fmt.Println("protocols: ", viper.GetStringSlice("server.protocols"))
	fmt.Println("ports: ", viper.GetIntSlice("server.ports"))
	fmt.Println("timeout: ", viper.GetDuration("server.timeout"))

	fmt.Println("mysql ip: ", viper.GetString("mysql.ip"))
	fmt.Println("mysql port: ", viper.GetInt("mysql.port"))

	if viper.IsSet("redis.port") {
		fmt.Println("redis.port is set")
	} else {
		fmt.Println("redis.port is not set")
	}

	fmt.Println("mysql settings: ", viper.GetStringMap("mysql"))
	fmt.Println("redis settings: ", viper.GetStringMap("redis"))
	fmt.Println("all settings: ", viper.AllSettings())

	sc.Load()

	time.Sleep(time.Hour)
}
