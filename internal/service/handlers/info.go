package handlers

import (
	"encoding/json"
	"gitlab.com/distributed_lab/ape"
	"net/http"
)

// Info godoc
// @Summary      Info
// @Description  Print info
// @Accept       json
// @Produce      json
// @Failure      500  {object}  jsonapi.ErrorsPayload
// @Router       /info [get]
func Info(w http.ResponseWriter, r *http.Request) {
	message := "It's our database service"
	if err := json.NewEncoder(w).Encode(message); err != nil {
		Log(r).WithError(err).Error("error during writing message")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ape.Render(w, message)
}
