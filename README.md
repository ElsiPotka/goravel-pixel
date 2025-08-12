<div align="center">

<img src="https://www.goravel.dev/logo.png?v=1.14.x" width="300" alt="Logo">

[![Doc](https://pkg.go.dev/badge/github.com/goravel/framework)](https://pkg.go.dev/github.com/goravel/framework)
[![Go](https://img.shields.io/github/go-mod/go-version/goravel/framework)](https://go.dev/)
[![Release](https://img.shields.io/github/release/goravel/framework.svg)](https://github.com/goravel/framework/releases)
[![Test](https://github.com/goravel/framework/actions/workflows/test.yml/badge.svg)](https://github.com/goravel/framework/actions)
[![Report Card](https://goreportcard.com/badge/github.com/goravel/framework)](https://goreportcard.com/report/github.com/goravel/framework)
[![Codecov](https://codecov.io/gh/goravel/framework/branch/master/graph/badge.svg)](https://codecov.io/gh/goravel/framework)
![License](https://img.shields.io/github/license/goravel/framework)

</div>


## About Goravel

Goravel is a web application framework with complete functions and good scalability. As a starting scaffolding to help
Gopher quickly build their own applications.

The framework style is consistent with [Laravel](https://github.com/laravel/laravel), let Php developer don't need to learn a
new framework, but also happy to play around Golang! In tribute to Laravel!

Welcome to star, PR and issues！

## Getting started

```
// Copy .env from .env.example

// Generate APP_KEY
go run . artisan key:generate

// Generate JWT key
go run . artisan jwt:secret

Install cosmtrek/air, Goravel has a built-in configuration file that can be used directly:

air

If you are using Windows system, you need to modify the .air.toml file in the root directory, and add the .exe suffix to the following two lines:

[build]
  bin = "./storage/temp/main.exe"
  cmd = "go build -o ./storage/temp/main.exe ."

  go run . artisan jwt:secret

Encrypt and decrypt env file

You may want to add the production environment env file to version control, but you don't want to expose sensitive information, you can use the env:encrypt command to encrypt the env file:

go run . artisan env:encrypt

// Specify the file name and key
go run . artisan env.encrypt --name .env.safe --key BgcELROHL8sAV568T7Fiki7krjLHOkUc

Then use the env:decrypt command to decrypt the env file in the production environment:

GORAVEL_ENV_ENCRYPTION_KEY=BgcELROHL8sAV568T7Fiki7krjLHOkUc go run . artisan env:decrypt

// or
go run . artisan env.decrypt --name .env.safe --key BgcELROHL8sAV568T7Fiki7krjLHOkUc
```

## Documentation

Online documentation [https://www.goravel.dev](https://www.goravel.dev)

Example [https://github.com/goravel/example](https://github.com/goravel/example)


## License

The Goravel framework is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
