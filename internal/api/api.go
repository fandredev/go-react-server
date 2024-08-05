package api

import (
	"net/http"

	"github.com/fandredev/go-react-server/internal/store/pgstore"
	"github.com/go-chi/chi/v5"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	api := apiHandler{
		q: q,
	}

	r := chi.NewRouter()

	api.r = r

	return api
}