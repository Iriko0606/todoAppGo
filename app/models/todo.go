package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserId    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `INSERT INTO todo (content, user_id, created_at) VALUES (?,?,?)`
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at from todo where id = ?`
	todo = Todo{}
	err = Db.QueryRow(cmd, id).Scan(&todo.ID,
		&todo.Content,
		&todo.UserId,
		&todo.CreatedAt)

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todo`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserId,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todo where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserId,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, todo)

	}
	rows.Close()

	return todos, err
}

func (t *Todo) UpdateTodo() (err error) {
	cmd := `update todo set content = ?, user_id = ? where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserId, t.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (t *Todo) DeleteTodo() (err error) {
	cmd := `delete from todo where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}