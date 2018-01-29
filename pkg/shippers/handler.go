package shippers

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sudhanshuraheja/tanker/pkg/appcontext"

	"github.com/gorilla/mux"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request)

type ShipperHandler struct{}

func (s *ShipperHandler) Add(ctx *appcontext.AppContext, db *sqlx.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("In the right function")

		vars := mux.Vars(r)
		fmt.Println("Vars", vars)
		fmt.Println("Query", r.URL.Query())
		fmt.Println("Header", r.Header)

		// data, err := json.Marshal({})
		// if err != nil {
		// 	logger.Error(err)
		// }
		w.Write([]byte("{\"success\":\"pong\"}"))

	}
}
