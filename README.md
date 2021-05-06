# Wish List

## Dependencies

0. **Golang v1.16 or newer**
1. JWT-GO v3.2.0
2. Gin CORS v1.3.1
3. Go Playground Validator v10.4.1
4. Go SQL Driver v1.6
5. Crypto v0.0.0-20210421170649-83a5a9bb288b

## Build Setup

### Production Environment

1. Set at environment variables: `GIN_MODE=release`
2. `go get -d -v`
3. `go install -v`
4. `./tthk-wish-list`
5. Server should run on port 8080.

### Development Environment

1. `go get -d -v`
2. `go run .` **OR** `go build . && ./tthk-wish-list`

**Using Dockerfile, you just need to define `GIN_MODE` environment variable.**
