package service

import (
	"cafe-service/internal/data/pg"
	address "cafe-service/internal/service/handlers/address"
	cafe "cafe-service/internal/service/handlers/cafe"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"

	"cafe-service/internal/service/helpers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()
	log := s.log.WithFields(map[string]interface{}{
		"service": "cafe-service",
	})

	r.Use(
		ape.RecoverMiddleware(log),
		ape.LoganMiddleware(log),
		ape.CtxMiddleware(
			helpers.CtxLog(log),
			helpers.CtxAddressesQ(pg.NewAddressesQ(s.db)),
			helpers.CtxCafesQ(pg.NewCafesQ(s.db)),
		),
	)
	r.Route("/integrations/cafe-service", func(r chi.Router) {
		r.Route("/addresses", func(r chi.Router) {
			r.Post("/", address.CreateAddress)
			r.Get("/", address.GetAddressList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", address.GetAddress)
				r.Put("/", address.UpdateAddress)
				r.Delete("/", address.DeleteAddress)
			})
		})
		r.Route("/cafes", func(r chi.Router) {
			r.Post("/", cafe.CreateCafe)
			r.Get("/", cafe.GetCafeList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", cafe.GetCafe)
				r.Put("/", cafe.UpdateCafe)
				r.Delete("/", cafe.DeleteCafe)
			})
		})
	})

	return r
}
