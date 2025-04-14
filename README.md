# Bluebase

A custom extension of regular [Pocketbase](https://github.com/pocketbase/pocketbase) with [bluemonday](https://github.com/microcosm-cc/bluemonday) sanitization:

- Default sanitization behaviour for maximum user content preservation by using [UGCPolicy](https://pkg.go.dev/github.com/microcosm-cc/bluemonday#UGCPolicy)
- Extends the JSVM binding to enable using sanitization with JS extension, so that you can make use of bluemonday sanitization in your JS code.

# How to use

1. To init the dependencies, run `go mod init bluebase && go mod tidy`.

2. To start the application, run `go run main.go serve`.

3. To build an executable, you can run `CGO_ENABLED=0 go build` and then start the created executable with `./bluebase serve`.

# Contributing

Bluebase is licensed under the [MIT License](LICENSE.md).
