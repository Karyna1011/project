package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type GatewayRequest struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	Duration  int64  `json:"duration"`
}

func NewPostGatewayRequest(r *http.Request) (GatewayRequest, error) {
	var request GatewayRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return GatewayRequest{}, errors.Wrap(err, "error requesting")
	}

	return request, nil
}
