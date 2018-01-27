package server

import (
	"fmt"
	"net/http"

	"github.com/jeffbmartinez/delay"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"

	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
)

// StartAPIServer : setup routes and start the server
func StartAPIServer() error {
	server := negroni.New()
	router := Router()

	server.Use(negroni.NewRecovery())
	server.Use(negroni.NewLogger())

	if config.EnableDelayMiddleware() {
		server.Use(delay.Middleware{})
	}

	if config.EnableGzipCompression() {
		server.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	if config.EnableStaticFileServer() {
		server.Use(negroni.NewStatic(http.Dir("data")))
	}

	server.Use(Recover())
	server.UseHandler(router)

	serverURL := fmt.Sprintf(":%s", config.Port())
	logger.Infoln("The server is now running at", serverURL)
	return http.ListenAndServe(serverURL, server)
}

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
