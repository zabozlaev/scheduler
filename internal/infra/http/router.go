package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/graphql-go/handler"
	"net/http"
	"scheduler/internal/infra/api"
)

func (a *adapter) newRouter(gqlAdapter *api.Adapter) (http.Handler, error) {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{a.config.FrontendUrl},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(cors.Handler, a.loggerMiddleware)

	h := handler.New(&handler.Config{
		Schema:           gqlAdapter.GqlSchema,
		Pretty:           true,
		Playground:       a.config.Mode == "development",
	})

	r.Handle("/graphql", h)
	return r, nil
}

