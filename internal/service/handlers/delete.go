package handlers

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	err := Person(r).Delete(name)
	if err != nil {
		Log(r).WithError(err).Error("error deleting person")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ape.Render(w, "Deleted")

}
