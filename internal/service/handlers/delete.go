package handlers

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"net/http"
)

// Delete godoc
// @Summary      Delete
// @Description  Get record by uuid
// @Accept       json
// @Produce      json
// @Param name  path string true "UUID of record"
// @Failure      400  {object}  jsonapi.ErrorsPayload
// @Failure      500  {object}  jsonapi.ErrorsPayload
// @Router       /get/{name} [delete]
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
