package requester

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request)

type MockServer struct {
	server *http.Server
}

func (m *MockServer) Start(port string) {

	router := mux.NewRouter()

	router.HandleFunc("/get", m.get()).Methods(http.MethodGet)
	router.HandleFunc("/post", m.post()).Methods(http.MethodPost)
	router.HandleFunc("/put", m.put()).Methods(http.MethodPut)
	router.HandleFunc("/delete", m.delete()).Methods(http.MethodDelete)
	router.HandleFunc("/upload", m.upload()).Methods(http.MethodPut)

	m.server = &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		Handler:      router,
	}

	go func() {
		err := m.server.ListenAndServe()
		if err != nil {
			if err.Error() != "http: Server closed" {
				fmt.Println("MockServer: the server is not running anymore,", err.Error())
			}
		}
	}()
}

func (m *MockServer) Stop() {
	err := m.server.Shutdown(nil)
	if err != nil {
		panic(err)
	}
}

func (m *MockServer) parseKeyFromQuery(r *http.Request, key string) string {
	value := ""
	if len(r.URL.Query()[key]) > 0 {
		value = r.URL.Query()[key][0]
	}
	return value
}

func (m *MockServer) parseKeyFromVars(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}

func (m *MockServer) writeJSON(w http.ResponseWriter, status int, response string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func (m *MockServer) get() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		m.writeJSON(w, http.StatusOK, `{ "data": { "method": "get" }, "success": "true" }`)
	}
}

func (m *MockServer) post() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), 500)
		}
		m.writeJSON(w, http.StatusOK, fmt.Sprintf(`{ "data": { "method": "post", "name": "%s" }, "success": "true" }`, r.FormValue("name")))

	}
}

func (m *MockServer) put() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		m.writeJSON(w, http.StatusOK, `{ "data": { "method": "put" }, "success": "true" }`)
	}
}

func (m *MockServer) delete() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		m.writeJSON(w, http.StatusOK, `{ "data": { "method": "delete" }, "success": "true" }`)
	}
}

func (m *MockServer) upload() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {

		// https://gist.github.com/aodin/9493190
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		json := fmt.Sprintf(`{ "data": { "method": "upload", "content": "%s" }, "success": "true" }`, string(body))
		m.writeJSON(w, http.StatusOK, json)
	}
}
