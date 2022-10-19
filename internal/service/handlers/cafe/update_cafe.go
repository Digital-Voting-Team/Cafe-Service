package handlers

import (
	"Cafe-Service/internal/data"
	"Cafe-Service/internal/service/helpers"
	requests "Cafe-Service/internal/service/requests/cafe"
	"Cafe-Service/resources"
	"github.com/spf13/cast"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateCafe(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateCafeRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	cafe, err := helpers.CafesQ(r).FilterById(request.CafeID).Get()
	if cafe == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	newCafe := data.Cafe{
		CafeName:  request.Data.Attributes.CafeName,
		Rating:    request.Data.Attributes.Rating,
		AddressId: cast.ToInt64(request.Data.Relationships.Address.Data.ID),
	}

	relateAddress, err := helpers.AddressesQ(r).FilterById(newCafe.AddressId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get new address")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	var resultCafe data.Cafe
	resultCafe, err = helpers.CafesQ(r).FilterById(cafe.Id).Update(newCafe)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to update cafe")
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
