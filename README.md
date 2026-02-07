# Refinery

## Docker

Start up postgres: `docker run -d -e POSTGRES_PASSWORD=password -e POSTGRES_USER=craig -p 5432:5432 postgres:latest`

## Startup

Create `env.yaml` config file

Start up service: `go run main.go`

## Monitoring

Monitor Rabbit at `http://localhost:15672/#/`

Default user is `guest` and password is `guest`

# Architecture Overview



