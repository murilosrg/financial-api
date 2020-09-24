# Financial API

[![license](https://img.shields.io/github/license/murilosrg/financial-api)](https://opensource.org/licenses/MIT)
[![go-mod](https://img.shields.io/github/go-mod/go-version/murilosrg/financial-api)](https://github.com/murilosrg/financial-api)

A simple api application written in Go

## Quick Start

```sh
$ docker-compose up --build
```

## Development

### 1.Clone Source Code

```shell
$ git clone https://github.com/murilosrg/financial-api

$ cd financial-api
```

### 2.Download Requirements

```shell
$ go mod download
```

### 3.Create Configuration

```shell
$ touch ./configuration.yml
```

```yml
# configuration for dev
db:
  driver: sqlite3
  address: ./financial.db
  # driver: mysql
  # addr: user:password@/dbname?charset=utf8&parseTime=True&loc=Local
  # driver: postgres
  # addr: host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword
address: :5000
```

### 4.Init and Run

```shell
$ go run ./cmd/financial -init
```

### 5. Visit Website

visit [localhost:5000](http://localhost:5000)

## Run Test

```shell
$ go test ./...
```

## Source

Repository: [financial-api](https://github.com/murilosrg/financial-api)

Author: [murilosrg](https://github.com/murilosrg)

## License

[MIT](https://opensource.org/licenses/MIT)
