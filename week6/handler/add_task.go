// `AddTask` handler is responsible for adding a new task to the store.
// early return

package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"week6/entity"
	"week6/store"

	"github.com/go-playground/validator/v10"
)

type AddTask struct {
	Store     *store.TaskStore
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// 참고: https://www.joinc.co.kr/w/man/12/golang/tag
	var b struct {
		Title string `json:"title" validate:"required"` // tag를 통해 struct의 metadata를 정의할 수 있다.
	}
	// reflect 패키지를 통해 struct의 metadata를 읽어 프로그램의 런타임에 변수와 값을 검사할 수 있다. (ex. Tag 형태)

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {

		// json.NewDecoder는 request 파싱을 위한 buffer를 생성.
		// Decode는 parameter인 struct를 받아서 Decoder의 reader r 값을 읽고, 넣어준다.
		// b는 reference type이므로, Decode 후에도 값이 유지된다. (Decode parameter는 any지만 내부적으로 reflect를 사용하여 pointer를 받아서 값을 변경)
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	err := validator.New().Struct(b) // Error would occur here
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	t := &entity.Task{
		Title:   b.Title,
		Status:  "todo",
		Created: time.Now(),
	}
	id, err := store.Tasks.Add(t)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID int `json:"id"`
	}{ID: int(id)}
	RespondJSON(ctx, w, rsp, http.StatusOK)

}
