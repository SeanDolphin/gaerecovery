package gaerecovery

import (
	"net/http"
	"runtime/debug"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type recovery struct {
	PrintStack bool
}

func Recovery() *recovery {
	return &recovery{
		PrintStack: true,
	}
}

func (rec *recovery) ServeHTTP(writer http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)

			ctx := appengine.NewContext(req)
			log.Criticalf(ctx, "%s", err.Error())
			if rec.PrintStack {
				log.Debugf(ctx, "%s", debug.Stack())
			}
		}
	}()

	next(writer, req)
}
