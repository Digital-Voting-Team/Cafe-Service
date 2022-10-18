package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteCafeRequest struct {
	CafeID int64 `url:"-"`
}

func NewDeleteCafeRequest(r *http.Request) (DeleteCafeRequest, error) {
	request := DeleteCafeRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.CafeID = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
