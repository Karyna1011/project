package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/subgroup/project/internal/data"
	"gitlab.com/tokend/subgroup/project/internal/service/requests"
	"gitlab.com/tokend/subgroup/project/resources"
	"net/http"
)

// List godoc
// @Summary      List
// @Description  Get list of records
// @Accept       json
// @Produce      json
// @Success      200  {object}  resources.PersonListResponse
// @Failure      400  {object}  jsonapi.ErrorsPayload
// @Failure      500  {object}  jsonapi.ErrorsPayload
// @Router       /list [get]
func List(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetPersonListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	people, err := Person(r).Select(req.OffsetPageParams)
	if err != nil {
		Log(r).WithError(err).Error("error selecting persons")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := newPeopleList(people, GetOffsetLinks(r, req.OffsetPageParams))

	ape.Render(w, response)
}

func newPeopleList(models []data.Person, links *resources.Links) resources.PersonListResponse {
	result := resources.PersonListResponse{
		Data: make([]resources.Person, len(models)),
	}

	for i, item := range models {
		result.Data[i] = newPersonModel(item).Data
	}

	result.Links = links
	return result
}

func newPersonModel(model data.Person) resources.PersonResponse {
	return resources.PersonResponse{
		Data: resources.Person{
			Key: resources.NewKeyInt64(model.Id, resources.PERSON),
			Attributes: resources.PersonAttributes{
				Name:      model.Name,
				Completed: model.Completed,
				Duration:  model.Duration,
			},
		},
	}
}
