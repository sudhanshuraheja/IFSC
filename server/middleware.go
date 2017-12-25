package server

import (
	"net/http"

	"github.com/sudhanshuraheja/ifsc/logger"
	"github.com/urfave/negroni"
)

// Recover : middleware for recovering after panic
func Recover() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorrf(r, "Recovered from panic: %+v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}()
		next(w, r)
	})
}
