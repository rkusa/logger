// Package logger provides a logging middleware that works well (but not
// exclusively) with [rkusa/web](https://github.com/rkusa/web).
//
//  app := web.New()
//  app.Use(logger.Middleware())
//
package logger

import (
	"log"
	"net/http"
	"os"
	"time"
)

type withStatus interface {
	http.ResponseWriter
	Status() int
}

func Middleware(logger *log.Logger) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	if logger == nil {
		logger = log.New(os.Stdout, "[web] ", 0)
	}

	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		start := time.Now()
		logger.Printf(
			">> %s %s",
			r.Method, r.URL.Path,
		)

		next(rw, r)

		if rw, ok := rw.(withStatus); ok {
			logger.Printf(
				"<< %s %s %d in %v",
				r.Method, r.URL.Path, rw.Status(), time.Since(start),
			)
		} else {
			logger.Printf(
				"<< %s %s in %v",
				r.Method, r.URL.Path, time.Since(start),
			)
		}
	}
}
