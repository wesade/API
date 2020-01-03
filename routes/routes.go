package routes

import (
	"github.com/go-chi/chi"
	"uni/API/handler"
)

func CreateRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/students", handler.GetAllStudents)
	router.Get("/students/{id}", handler.GetSingleStudent)
	router.Post("/", handler.CreateNewStudent)
	router.Delete("/students/{id}", handler.DeleteStudent)
	router.Put("/students/{id}", handler.UpdateStudent)

	return router
}
