package requests

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetCafeRequest struct {
	CafeId int64 `url:"-"`
}

func NewGetCafeRequest(r *http.Request) (GetCafeRequest, error) {
	request := GetCafeRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.CafeId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
