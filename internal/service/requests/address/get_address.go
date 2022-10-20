package requests

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetAddressRequest struct {
	AddressId int64 `url:"-"`
}

func NewGetAddressRequest(r *http.Request) (GetAddressRequest, error) {
	request := GetAddressRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.AddressId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
