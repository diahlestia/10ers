# 10xers
> **The backend API for getting token from MagicEden API.**

## Table of Contents
- [10xers](#10xers)
  - [Table of Contents](#table-of-contents)
  - [Running in Your Local Environment](#running-in-your-local-environment)


## Running in Your Local Environment

Do these steps firstly:
1. Clone this repository to your local environment
2. Copy `.env.example` to `.env` and set the environment variables in `.env` adjust the value with your local env
3. Run `go run migrate/migrate.go` on the terminal. This will create the table to your local database.
4. Run `go run main.go` Open http://localhost:8080. If it shows `{"success": true}`, congratulations! Your setup is successful.