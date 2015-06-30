package gaerecovery

import "net/http"

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
			logger(req, err, rec.PrintStack)
		}
	}()

	next(writer, req)
}

var logger = func(req *http.Request, err interface{}, printstack bool) {}
