// A logger middleware for [rkgo/web](https://github.com/rkgo/web)
//
//  app := web.New()
//  app.Use(logger.Middleware())
//
package logger

import (
	"log"
	"os"
	"time"

	"github.com/rkgo/web"
)

func Middleware(logger *log.Logger) web.Middleware {
	if logger == nil {
		logger = log.New(os.Stdout, "[web] ", 0)
	}

	return func(ctx web.Context, next web.Next) {
		start := time.Now()
		logger.Printf(
			">> %s %s",
			ctx.Req().Method, ctx.Req().URL.Path,
		)

		next(ctx)

		logger.Printf(
			"<< %s %s %v in %v",
			ctx.Req().Method, ctx.Req().URL.Path, ctx.Status(), time.Since(start),
		)
	}
}
