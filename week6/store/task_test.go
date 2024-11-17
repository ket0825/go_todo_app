// 실제 RDBMS를 사용해 ListTasks 메서드를 테스트하는 함수
// prepareTask 함수는 task 테이블의 상태를 관리하는 헬퍼 함수

package store

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	"github.com/ket0825/go_todo_app/clock"
	"github.com/ket0825/go_todo_app/entity"
	"github.com/ket0825/go_todo_app/testutil"
)

// DELETE문 사용하기에 주의!!!
// Test ListTasks (READ) by real DB
func TestRepository_ListTasks(t *testing.T) {
	ctx := context.Background()
	// entity.Task를 작성하는 다른 test case와 섞이면 테스트가 실패함.
	// 이를 방지하기 위해 트랜잭션으로 test case 내로 한정된 테이블 상태를 관리
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	wants := prepareTasks(ctx, t, tx)

	sut := &Repository{}
	gots, err := sut.ListTasks(ctx, tx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d := cmp.Diff(gots, wants); len(d) != 0 {
		t.Errorf("differs: (-got +want)\n%s", d)
	}
}

func prepareTasks(ctx context.Context, t *testing.T, con Execer) entity.Tasks {
	t.Helper()

	if _, err := con.ExecContext(ctx, "DELETE FROM task;"); err != nil {
		t.Logf("failed to initialize task: %v", err)
	}
	c := clock.FixedClocker{}
	wants := entity.Tasks{
		{
			Title:    "want task 1",
			Status:   "todo",
			Created:  c.Now(),
			Modified: c.Now(),
		},
		{
			Title:    "want task 2",
			Status:   "todo",
			Created:  c.Now(),
			Modified: c.Now(),
		},
		{
			Title:    "want task 3",
			Status:   "done",
			Created:  c.Now(),
			Modified: c.Now(),
		},
	}
	result, err := con.ExecContext(ctx,
		`INSERT INTO task (title, status, created, modified) 
		VALUES 
		(?, ?, ?, ?),
		(?, ?, ?, ?),
		(?, ?, ?, ?);`,
		wants[0].Title, wants[0].Status, wants[0].Created, wants[0].Modified,
		wants[1].Title, wants[1].Status, wants[1].Created, wants[1].Modified,
		wants[2].Title, wants[2].Status, wants[2].Created, wants[2].Modified,
	)
	if err != nil {
		t.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	// ? 왜 id가 1부터시작하지? LastInsertId()는 가장 최근에 추가된 row의 id를 반환하는데?
	// MySQL에서 다중 insert문을 사용했기 때문에 첫 id가 LastInsertId()로 반환된 것이다.
	wants[0].ID = entity.TaskID(id)
	wants[1].ID = entity.TaskID(id + 1)
	wants[2].ID = entity.TaskID(id + 2)
	return wants
}

// Test AddTask (CREATE) by mock
func TestRepository_AddTask(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	c := clock.FixedClocker{}
	var wantID int64 = 20
	okTask := &entity.Task{
		Title:    "ok task",
		Status:   "todo",
		Created:  c.Now(),
		Modified: c.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close() })
	mock.ExpectExec(
		// escape 문자 필요
		`INSERT INTO task \(title, status, created, modified\) 
		VALUES 
		\(\?, \?, \?, \?\);`,
	).WithArgs(okTask.Title, okTask.Status, okTask.Created, okTask.Modified).
		WillReturnResult(sqlmock.NewResult(wantID, 1))

	xdb := sqlx.NewDb(db, "mysql")
	r := &Repository{Clocker: c}
	if err := r.AddTask(ctx, xdb, okTask); err != nil {
		t.Fatalf("want no error, but got %v", err)
	}
}
