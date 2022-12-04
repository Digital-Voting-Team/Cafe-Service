package requests

import (
	"cafe-service/internal/service/helpers"
	"cafe-service/resources"
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateCafeRequest struct {
	Data resources.Cafe
}

func NewCreateCafeRequest(r *http.Request) (CreateCafeRequest, error) {
	var request CreateCafeRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreateCafeRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/attributes/name": validation.Validate(&r.Data.Attributes.CafeName, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/rating": validation.Validate(&r.Data.Attributes.Rating,
			validation.Length(3, 30)),
		"/data/relationships/address/data/id": validation.Validate(&r.Data.Relationships.Address.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
	}).Filter()
}
