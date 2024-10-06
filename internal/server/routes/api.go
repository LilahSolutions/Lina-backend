package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func New() http.Handler {
	r := chi.NewRouter()

	chatRouter := ChatRouter{}

	r.Mount("/chat", chatRouter.Routes())

	return r
}
