# proyect-name

[![go](https://img.shields.io/badge/go-v1.17.X-cyan.svg)](https://golang.org/)
[![make](https://img.shields.io/badge/make-v3.8.X-yellow.svg)](https://www.gnu.org/software/make/)

## Prerequisites

You will need the following things properly installed on your computer.

* [Git](http://git-scm.com/)
* [Go](https://golang.org/)
* [Make](https://www.gnu.org/software/make/)

## Installation

Following you can find the instructions:

* clone this repository
* change into the new directory `cd project-name`
* execute `make build` and `make install`
* you can see the app binary in bin directory of you `$GOPATH`

## Build

Run `make lambda` to build the project. The build artifacts will be stored in the `./function-name/dist/` directory if you ran the second option.

## Run App

Run `go run function-name/cmd/main.go` on the app root directory.

## Running unit tests

Run `make test` to execute the unit tests.

## Running format check

Run `make fmt` to execute the format check.

## Running Linter

Run `make lint` to execute linter verify.