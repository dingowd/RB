package server

import (
	"encoding/json"
	"github.com/dingowd/RB/app"
	"github.com/dingowd/RB/model"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"sync"
)

type Server struct {
	App  *app.App
	Addr string
	Srv  *http.Server
	mu   sync.Mutex
}

func NewServer(app *app.App, addr string) *Server {
	s := &Server{App: app, Addr: addr}
	return s
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id is missing!!!"))
		return
	}
	err := s.App.Store.Delete(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("ok"))
}

func (s *Server) Update(w http.ResponseWriter, r *http.Request) {
	var d model.CacheStudent
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		msg := "Wrong request. Check Data" + err.Error()
		w.Write([]byte(msg))
		return
	}
	err = s.App.Store.Update(d)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("ok"))
}

func (s *Server) GetAll(w http.ResponseWriter, r *http.Request) {
	c := s.App.Cache.ReadFromCache()
	tmpl, _ := template.ParseFiles("./templates/index.html")
	tmpl.Execute(w, c)
}

func (s *Server) Insert(w http.ResponseWriter, r *http.Request) {
	var d model.ForJson
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		msg := "Wrong request. Check Data" + err.Error()
		w.Write([]byte(msg))
		return
	}
	err = s.App.Store.Insert(d)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("ok"))
}

func (s *Server) Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/delete", s.Delete).Methods("DELETE")
	router.HandleFunc("/update", s.Update).Methods("POST")
	router.HandleFunc("/get", s.GetAll).Methods("GET")
	router.HandleFunc("/insert", s.Insert).Methods("POST")

	http.Handle("/", router)
	Srv := &http.Server{Addr: s.Addr, Handler: router}
	s.Srv = Srv
	s.App.Log.Info("http сервер запускается")
	err := s.Srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	s.App.Log.Info("остановка http сервера")
	return s.Srv.Close()
}
