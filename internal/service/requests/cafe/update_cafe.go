package requests

import (
	"Cafe-Service/internal/service/helpers"
	"Cafe-Service/resources"
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type UpdateCafeRequest struct {
	CafeId int64 `url:"-" json:"-"`
	Data   resources.Cafe
}

func NewUpdateCafeRequest(r *http.Request) (UpdateCafeRequest, error) {
	request := UpdateCafeRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.CafeId = cast.ToInt64(chi.URLParam(r, "id"))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *UpdateCafeRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/attributes/name": validation.Validate(&r.Data.Attributes.CafeName, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/rating": validation.Validate(&r.Data.Attributes.Rating,
			validation.Length(3, 45)),
		"/data/relationships/address/data/id": validation.Validate(&r.Data.Relationships.Address.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
	}).Filter()
}
