package Routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/my_todo/Handlers"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func Route() {

	r := chi.NewRouter()
	r.Method("POST", "/Register/", Handler(Handlers.Register))

	r.Method("POST", "/Show/", Handler(Handlers.Show))
	r.Method("POST", "/Create/", Handler(Handlers.CreateUser))
	r.Method("POST", "/Test/", Handler(Handlers.Test))
	//	r.Method("POST", "/Login", Handler(Handlers.Login))
	//	r.Method("GET", "/Signup", Handler(Handlers.Signup))
	//	r.Method("GET", "/", Handler(Handlers.Custom))
	//	r.Method("GET", "/", Handler(Handlers.Custom))
	//	r.Method("GET", "/", Handler(Handlers.Custom))
	//	r.Method("GET", "/", Handler(Handlers.Custom))

	http.ListenAndServe("3333", r)
}
