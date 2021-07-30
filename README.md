# Leboncoin

Technical test for [Leboncoin](https://leboncoingroupe.com/).

Web API written in Go.

## Objective

Your goal is to implement a web server that will expose a REST API endpoint that:
- [x] Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
- [x] Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:
- [x] Ready for production
- [x] Easy to maintain by other developers

Bonus: add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- [x] Accept no parameter
- [x] Return the parameters corresponding to the most used request, as well as the number of hits for this request


## 💻 Getting started 

Using go cli:
```bash
$ go run .
⇨ http server started on [::]:8080
````

Using Docker:
```bash
$ docker build -t leboncoin .
$ docker run -p 8080:8080 leboncoin
```

This will expose an HTTP web server on the 8080 port.

You can access it with this URL: http://127.0.0.1:8080/

## REST Endpoints

### GET /

Returns a list of strings with numbers from 1 to limit, where:
- all multiples of int1 are replaced by str1 
- all multiples of int2 are replaced by str2
- all multiples of int1 and int2 are replaced by str1str2

#### URL parameters

- `int1` integer
- `int2` integer
- `limit` integer
- `str1` string
- `str2` string

#### Example

```bash
$ export INT1=3 INT2=5 LIMIT=10 STR1=fizz STR2=buzz
$ curl "http://127.0.0.1:8080/?int1=$INT1&int2=$INT2&limit=$LIMIT&str1=$STR1&str2=$STR2"
["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz"]
```

### GET /metrics

Return the parameters corresponding to the most used request, as well as the number of hits for this request.
Member key is constructed as the following: `int1_int2_limit_str1_str2`

### Example

```bash
$ curl "http://127.0.0.1:8080/metrics"
[
  {
    "Score": 1,
    "Member": "3_5_30_fizz_buzz"
  },
  {
    "Score": 2,
    "Member": "3_5_10_fizz_buzz"
  },
  {
    "Score": 3,
    "Member": "32_5_10_fizz_buzz"
  },
  {
    "Score": 11,
    "Member": "32_531_10_fizz_buzz"
  }
]
```

## 🤖 CI

A CI is running on every push thanks to [Github Actions](https://github.com/features/actions).
This CI contains a [lint](.github/workflows/lint.yml) and a [build](.github/workflows/build.yml) workflows.

That follows the best Go lint practices and test the build on every code changes.

## ⚙️ Config

You can specify optional environment variables to change default server variables:

- `PORT` The web endpoint listening port (default: `8080`)
- `HOST` The web endpoint listening address (default: `0.0.0.0`)
- `REDIS_URL` The redis URL to connect to (default: inmemory redis instance)

## Credits

Made by Antoine Ordonez.

Website: https://shellbear.me