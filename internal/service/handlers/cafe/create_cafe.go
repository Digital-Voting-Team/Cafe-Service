package handlers

import (
	"cafe-service/internal/data"
	"cafe-service/internal/service/helpers"
	requests "cafe-service/internal/service/requests/cafe"
	"cafe-service/resources"
	"github.com/spf13/cast"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateCafe(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateCafeRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Cafe := data.Cafe{
		CafeName:  request.Data.Attributes.CafeName,
		Rating:    request.Data.Attributes.Rating,
		AddressId: cast.ToInt64(request.Data.Relationships.Address.Data.ID),
	}

	var resultCafe data.Cafe
	relateAddress, err := helpers.AddressesQ(r).FilterById(Cafe.AddressId).Get()
	if err != nil {
		// TODO ask how to send address not found instead of internal error
		helpers.Log(r).WithError(err).Error("failed to get address")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	resultCafe, err = helpers.CafesQ(r).Insert(Cafe)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create cafe")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var includes resources.Included
	includes.Add(&resources.Address{
		Key: resources.NewKeyInt64(relateAddress.Id, resources.ADDRESS),
		Attributes: resources.AddressAttributes{
			BuildingNumber: relateAddress.BuildingNumber,
			Street:         relateAddress.Street,
			City:           relateAddress.City,
			District:       relateAddress.District,
			Region:         relateAddress.Region,
			PostalCode:     relateAddress.PostalCode,
		},
	})

	result := resources.CafeResponse{
		Data: resources.Cafe{
			Key: resources.NewKeyInt64(resultCafe.Id, resources.CAFE),
			Attributes: resources.CafeAttributes{
				CafeName: resultCafe.CafeName,
				Rating:   resultCafe.Rating,
			},
			Relationships: resources.CafeRelationships{
				Address: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultCafe.AddressId, 10),
						Type: resources.ADDRESS,
					},
				},
			},
		},
		Included: includes,
	}
	ape.Render(w, result)
}
