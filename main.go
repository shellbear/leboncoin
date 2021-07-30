package main

import (
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v8"
)

// RedisClient is a shared Redis client.
var RedisClient *redis.Client

// Config stores server config variables.
type Config struct {
	Host string `env:"HOST"`
	Port int `env:"PORT" envDefault:"8080"`
	RedisURL string `env:"REDIS_URL"`
}

// Creates a new redis client instance based on the redisURL parameter.

func newRedisClient(redisURL string) (*redis.Client, error) {
	// If user specified a REDIS_URL parameter, use it to connect to redis instance.
	if redisURL != "" {
		options, err := redis.ParseURL(redisURL)
		if err != nil {
			return nil, fmt.Errorf("failed to parse REDIS_URL: %v", err)
		}

		return redis.NewClient(options), nil
	}

	// If no REDIS_URL parameter has been specified, spawn a inmemory redis instance and use it to store metrics.
	s, err := miniredis.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to started inmemory redis instance: %v", err)
	}

	return redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	}), nil
}

func main() {
	var conf Config

	// Parse environment variables and store them inside config.
	if err := env.Parse(&conf); err != nil {
		log.Fatalln("Failed to parse environment variables:", err)
	}

	var err error
	RedisClient, err = newRedisClient(conf.RedisURL)
	if err != nil {
		log.Fatalln("Failed to init redis client:", err)
	}

	e := NewAPI()
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", conf.Host, conf.Port)))
}
