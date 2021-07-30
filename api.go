package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// MetricsRedisKey : the Redis Sorted Set key to save the requests metrics.
const MetricsRedisKey = "metrics"

/*
	This endpoint calls the FizzBuzz function, store metrics about the request and send the result to the client.
*/
func fizzBuzz(c echo.Context) error {
	// URL parameters.
	var (
		int1 int
		int2 int
		limit int
		str1 string
		str2 string
	)

	// Get, bind and check for parameters errors.
	if err := echo.QueryParamsBinder(c).
		MustInt("int1", &int1).
		MustInt("int2", &int2).
		MustInt("limit", &limit).
		MustString("str1", &str1).
		MustString("str2", &str2).
		BindError(); err != nil {
		log.Println("Failed to parse URL parameters:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Generates an unique member key for the request parameters.
	memberKey := fmt.Sprintf("%d_%d_%d_%s_%s", int1, int2, limit, str1, str2)

	// Use Redis ZIncrBy to store and sort requests count by score inside a sorted set.
	// More info: https://redis.io/topics/data-types#sorted-sets
	ctx := c.Request().Context()
	if err := RedisClient.ZIncrBy(ctx, MetricsRedisKey, 1, memberKey).Err(); err != nil {
		return err
	}

	// Process fizzbuzz and send back result to client.
	results := FizzBuzz(int1, int2, limit, str1, str2)
	return c.JSON(http.StatusOK, results)
}

/*
	This endpoint returns the parameters corresponding to the most used request, as well as the number of hits for
 	this request.
 */
func metrics(c echo.Context) error {
	ctx := c.Request().Context()

	// Use Redis ZRangeWithScores to fetch the metrics sorted set and send it back to client.
	results, err := RedisClient.ZRangeWithScores(ctx, MetricsRedisKey, 0, -1).Result()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, results)
}

// NewAPI Creates a new Echo API.
func NewAPI() *echo.Echo {
	e := echo.New()

	// We define the API REST endpoints.
	e.GET("/", fizzBuzz)
	e.GET("/metrics", metrics)

	return e
}