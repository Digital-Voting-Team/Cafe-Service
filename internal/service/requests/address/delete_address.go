package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteAddressRequest struct {
	AddressId int64 `url:"-"`
}

func NewDeleteAddressRequest(r *http.Request) (DeleteAddressRequest, error) {
	request := DeleteAddressRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.AddressId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
