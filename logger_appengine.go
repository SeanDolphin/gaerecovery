// +build appengine appenginevm

package gaerecovery

import (
	"net/http"
	"runtime/debug"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	logger = func(req *http.Request, err interface{}, printstack bool) {
		ctx := appengine.NewContext(req)
		log.Criticalf(ctx, "%s", err)
		if printstack {
			log.Debugf(ctx, "%s", debug.Stack())
		}
	}
}
