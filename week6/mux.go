package main

import (
	"net/http"

	"week6/handler"
	"week6/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux() http.Handler {
	mux := chi.NewRouter()
	// chi.Router의 HandlerFunc은 라우팅 패턴과, 핸들러 함수를 받고, 라우팅 패턴에 해당하는 요청이 들어왔을 때 핸들러 함수를 실행한다.
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	v := validator.New()
	mux.Handle("/tasks", &handler.AddTask{Store: store.Tasks, Validator: v})
	// AddTask를 handler로 사용해도 되는 이유: 내부적으로 메서드가 ServeHTTP를 구현하고 있기 때문에, HandlerFunc로 사용할 수 있다.
	at := &handler.AddTask{Store: store.Tasks, Validator: v}
	mux.Post("/tasks", at.ServeHTTP) // 그냥 위의 Handle에서 가능하지만 Post에 대한 핸들러를 따로 만들어서 사용하여 메서드 별 핸들러를 분리할 수 있다.
	lt := &handler.ListTask{Store: store.Tasks}
	mux.Get("/tasks", lt.ServeHTTP)
	return mux
}
