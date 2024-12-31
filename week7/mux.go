package main

import (
	"context"
	"net/http"

	"github.com/ket0825/go_todo_app/clock"
	"github.com/ket0825/go_todo_app/config"
	"github.com/ket0825/go_todo_app/handler"
	"github.com/ket0825/go_todo_app/service"
	"github.com/ket0825/go_todo_app/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	// chi.Router의 HandlerFunc은 라우팅 패턴과, 핸들러 함수를 받고, 라우팅 패턴에 해당하는 요청이 들어왔을 때 핸들러 함수를 실행한다.
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg) // connection 생성
	if err != nil {
		return nil, cleanup, err
	}
	r := store.Repository{Clocker: clock.RealClocker{}}
	at := &handler.AddTask{
		Service:   &service.AddTask{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{
		Service: &service.ListTask{DB: db, Repo: &r},
		// list이기에 validator는 사용하지 않음
	}
	mux.Get("/tasks", lt.ServeHTTP)
	ru := &handler.RegisterUser{
		Service:   &service.RegisterUser{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/register", ru.ServeHTTP)
	return mux, cleanup, nil
}
