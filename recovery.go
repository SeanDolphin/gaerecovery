package gaerecovery

import (
	"net/http"
	"runtime/debug"

	"appengine"
)

type recovery struct {
	PrintStack bool
}

func Recovery() *recovery {
	return &recovery{
		PrintStack: true,
	}
}

func (rec *recovery) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)

			context := appengine.NewContext(r)

			context.Criticalf("%s", err)

			if rec.PrintStack {
				stack := debug.Stack()
				context.Debugf("%s", stack)
			}
		}
	}()

	next(rw, r)
}
