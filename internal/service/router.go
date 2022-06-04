package service

import (
	"github.com/go-chi/chi"
	"github.com/swaggo/http-swagger"
	"gitlab.com/distributed_lab/ape"
	_ "gitlab.com/tokend/subgroup/project/docs"
	"gitlab.com/tokend/subgroup/project/internal/config"
	"gitlab.com/tokend/subgroup/project/internal/data/postgres"
	"gitlab.com/tokend/subgroup/project/internal/service/handlers"
	"net/http"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxPerson(postgres.NewPersonQ(cfg.DB())),
		),
	)

	r.Route("/project", func(r chi.Router) {
		r.Post("/add", handlers.Add)
		r.Get("/list", handlers.List)
		r.Get("/get/{name}", handlers.GetByIndex)
		r.Delete("/delete/{name}", handlers.Delete)
		r.Get("/info", handlers.Info)
	})
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:9000/swagger/doc.json"), //The url pointing to API definition
	))

	http.ListenAndServe(":9000", r)

	return r
}
