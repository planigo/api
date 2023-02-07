# PlanigoAPI

This is the API for Planigo using Go with [Go Fiber](https://gofiber.io/)

## Pre-requisites

- Go 1.19
- Docker with Docker Engine version at least 19.03.0

## Setup

- Create an account on [Mailgun](https://www.mailgun.com/) and get the private API key and domain
- Create a .env file in the root directory of the project and add the environment variables from `.env.dist`
- Run `go mod download` to download the dependencies
- Run `make install-deps` to install the global dependencies
- Run `docker compose up -d` to start the database
- Run `make config` to setup git hooks
- Run `make dev` to start the server with hot reloading