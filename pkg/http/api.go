package http

import (
	"goat/pkg/ai"
	"io"
	"net/http"
	"time"
)

type app struct {
	generator ai.Generator
	router *http.ServeMux
}

func NewApp(generator ai.Generator) http.Handler {
	app := &app{
		generator: generator,
		router: http.NewServeMux(),
	}

	app.router.Handle("/fakenews", app.fakenews())

	return app.router
}

func (a *app) fakenews() http.HandlerFunc {
	// one time init...
	return func (w http.ResponseWriter, r *http.Request) {
		title := a.generator.RandomTitle()
		if title == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, title)
	}
}

func (a *app) logExecutionTime(f http.HandlerFunc) http.HandlerFunc {
	// one time init...
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		println(time.Now().Sub(start).Nanoseconds())
	}
}