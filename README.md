# logger

A logger middleware that works well (but not exclusively) with [rkusa/web](https://github.com/rkusa/web).

[![GoDoc][godoc]](https://godoc.org/github.com/rkusa/logger)

### Example

```go
app := web.New()
app.Use(logger.Middleware())
```

## License

[MIT](LICENSE)

[godoc]: http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square
