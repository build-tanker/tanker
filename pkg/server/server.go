package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/jeffbmartinez/delay"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"

	"github.com/gojekfarm/tanker/pkg/appcontext"
)

type Server struct {
	ctx    *appcontext.AppContext
	db     *sqlx.DB
	server *http.Server
}

func NewServer(ctx *appcontext.AppContext, db *sqlx.DB) *Server {
	return &Server{
		ctx: ctx,
		db:  db,
	}
}

func (s *Server) Start() error {
	config := s.ctx.GetConfig()
	log := s.ctx.GetLogger()

	server := negroni.New()
	server.Use(negroni.NewRecovery())
	server.Use(negroni.NewLogger())

	router := Router(s.ctx, s.db)

	if config.EnableDelayMiddleware() {
		server.Use(delay.Middleware{})
	}

	if config.EnableGzipCompression() {
		server.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	if config.EnableStaticFileServer() {
		server.Use(negroni.NewStatic(http.Dir("data")))
	}

	serverURL := fmt.Sprintf(":%s", config.Port())

	server.Use(Recover())
	server.UseHandler(router)
	server.Run(serverURL)

	s.server = &http.Server{
		Addr:         serverURL,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		Handler:      server,
	}

	log.Infoln("[negroni] Listening on ", serverURL)
	go func() {
		err := s.server.ListenAndServe()
		if err != nil {
			if err.Error() != "http: Server closed" {
				fmt.Println("Server: the server is not running anymore,", err.Error())
			}
		}
	}()

	http.ListenAndServe(serverURL, server)

	return nil
}

func (s *Server) Stop() error {
	// Not sure how to stop a server
	s.server.Shutdown(nil)
	return nil
}

func Recover() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Recovered from panic: %+v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}()
		next(w, r)
	})
}
