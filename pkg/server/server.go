package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/urfave/negroni"

	"github.com/build-tanker/tanker/pkg/appcontext"
)

// Server provides a way to interact with the server
type Server struct {
	ctx    *appcontext.AppContext
	db     *sqlx.DB
	server *http.Server
}

// NewServer makes a new server object available
func NewServer(ctx *appcontext.AppContext, db *sqlx.DB) *Server {
	return &Server{
		ctx: ctx,
		db:  db,
	}
}

// Start a server
func (s *Server) Start() error {
	config := s.ctx.GetConfig()
	log := s.ctx.GetLogger()

	server := negroni.New()
	server.Use(negroni.NewRecovery())
	server.Use(negroni.NewLogger())

	router := Router(s.ctx, s.db)
	serverURL := fmt.Sprintf(":%s", config.Port())

	server.Use(recover())
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

// Stop a server
func (s *Server) Stop() error {
	s.server.Shutdown(nil)
	return nil
}

func recover() negroni.HandlerFunc {
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
