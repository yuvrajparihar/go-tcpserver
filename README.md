Multi-threaded TCP Server In Go
===

This codebase demonstrates multi-threaded tcp server in Golang.

- ref video: https://youtu.be/f9gUFy-9uCM

## How to Start Server

```
$ go run main.go
```

Fire the following commands on another terminal to simulate
multiple concurrent requests.

```
curl http://localhost:1729 &
curl http://localhost:1729 &
curl http://localhost:1729 &
```