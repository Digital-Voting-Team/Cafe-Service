package handlers

import (
	"Cafe-Service/internal/service/helpers"
	requests "Cafe-Service/internal/service/requests/cafe"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteCafe(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteCafeRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Cafe, err := helpers.CafesQ(r).FilterById(request.CafeID).Get()
	if Cafe == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	err = helpers.CafesQ(r).Delete(request.CafeID)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete cafe")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
