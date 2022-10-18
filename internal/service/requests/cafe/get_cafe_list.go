package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetCafeListRequest struct {
	pgdb.OffsetPageParams
	FilterName   []string `filter:"cafe_name"`
	FilterRating []string `filter:"rating"`
}

func NewGetCafeListRequest(r *http.Request) (GetCafeListRequest, error) {
	var request GetCafeListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
