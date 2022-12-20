package service

import (
	"github.com/Digital-Voting-Team/cafe-service/internal/data/pg"
	address "github.com/Digital-Voting-Team/cafe-service/internal/service/handlers/address"
	cafe "github.com/Digital-Voting-Team/cafe-service/internal/service/handlers/cafe"
	"github.com/Digital-Voting-Team/cafe-service/internal/service/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"gitlab.com/distributed_lab/ape"

	"github.com/Digital-Voting-Team/cafe-service/internal/service/helpers"
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
		middleware.BasicAuth(s.endpoints),
	)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/integrations/cafe-service", func(r chi.Router) {
		r.Use(middleware.CheckManagerPosition())
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
