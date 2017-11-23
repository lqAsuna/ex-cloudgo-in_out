package server

import (
	"net/http"
	"os"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)
func NewServer() *negroni.Negroni {
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx)
	n.UseHandler(mx)
	return n
}
func initRoutes(mx *mux.Router) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	mx.HandleFunc("/unknown", unknownhandle).Methods("GET")
	mx.HandleFunc("/api", apihandle).Methods("GET")
	mx.HandleFunc("/", formhandle).Methods("POST")
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
}

func unknownhandle(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusNotImplemented, "501 NotImplemented")
	}
}

func apihandle(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Nickname string `json:"username"`
			Password string `json:"content"`
		}{ID: "8675309", Content: "Hello!"})
	}
}

func formhandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		username := req.FormValue("username")
		content := req.FormValue("content")
		formatter.HTML(w, http.StatusOK, "detail", struct {
			Nickname string
			Password string
		}{Nickname: username, Password: content})
	}
}
