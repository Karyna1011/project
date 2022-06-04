package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/subgroup/project/internal/data"
	"gitlab.com/tokend/subgroup/project/internal/service/requests"
	"net/http"
)

// Add godoc
// @Summary      Add
// @Description  Adds record
// @Accept       json
// @Produce      json
// @Param        input body  requests.GatewayRequest true "Gets encoded data and sender's address"
// @Success      200  {int}     http.StatusOK
// @Failure      400  {object}  jsonapi.ErrorsPayload
// @Failure      500  {object}  jsonapi.ErrorsPayload
// @Router       /add [post]
func Add(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewPostGatewayRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	selected, err := Person(r).FilterByName(request.Name).Get()
	if err != nil {
		Log(r).WithError(err).Error("error finding person")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if selected == nil {
		_, err = Person(r).Insert(data.Person{
			Name:      request.Name,
			Completed: request.Completed,
			Duration:  request.Duration,
		})
	}
	if err != nil {
		Log(r).WithError(err).Error("error inserting person")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
