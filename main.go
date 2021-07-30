package main

import (
	"fmt"
	"log"

	"github.com/alicebob/miniredis/v2"
	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v8"
)

// RedisClient is a shared Redis client.
var RedisClient *redis.Client

// Config stores server config variables.
type Config struct {
	Host string `env:"HOST"`
	Port int `env:"PORT" envDefault:"8080"`
}

func main() {
	var conf Config

	// Parse environment variables and store them inside config.
	if err := env.Parse(&conf); err != nil {
		log.Fatalln("Failed to parse environment variables:", err)
	}

	// Run a inmemory redis instance to prevent external dependencies.
	s, err := miniredis.Run()
	if err != nil {
		log.Fatalln("Failed to started inmemory redis instance:", err)
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	e := NewAPI()
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", conf.Host, conf.Port)))
}
