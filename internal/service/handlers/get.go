package handlers

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

// GetByIndex godoc
// @Summary      GetByIndex
// @Description  Get record by uuid
// @Accept       json
// @Produce      json
// @Param name  path string true "UUID of record"
// @Success      200  {object}  resources.PersonResponse
// @Failure      400  {object}  jsonapi.ErrorsPayload
// @Failure      500  {object}  jsonapi.ErrorsPayload
// @Router       /get/{name} [get]
func GetByIndex(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	personQ := Person(r)

	person, err := personQ.FilterByName(name).Get()

	if err != nil {
		Log(r).WithError(err).Error("failed to get person")
		ape.Render(w, problems.InternalError())
		return
	}

	if person == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	ape.Render(w, person.Resource())

}
