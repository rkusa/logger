# logger

A logger middleware for [rkgo/web](https://github.com/rkgo/web)

[![GoDoc][godoc]](https://godoc.org/github.com/rkgo/logger)

### Example

```go
app := web.New()
app.Use(logger.Middleware())
```

[godoc]: http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square