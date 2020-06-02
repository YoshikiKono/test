package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type todoList []todo

var db *sql.DB

//構造体一覧
type todo struct {
	ID        string `db:"id" json:"id"`
	Schedule  string `db:"schedule" json:"schedule"`
	TimeLimit string `db:"timelimit" json:"timeLimit"`
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:0111@/todo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.Handle("/", http.FileServer(http.Dir("indexToDo")))
	http.HandleFunc("/register", register)
	http.HandleFunc("/display", display)
	http.HandleFunc("/remove", remove)
	log.Fatal(http.ListenAndServe(":8880", nil))
}

//予定登録時の処理
func register(w http.ResponseWriter, r *http.Request) {
	form := todo{}

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	schedule := form.Schedule
	timeLimit := form.TimeLimit

	stmt, err := db.Prepare("INSERT INTO trn_todo(schedule, timelimit) VALUES(?,?)")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(schedule, timeLimit)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

//DBの内容を一覧表示させる
func display(w http.ResponseWriter, r *http.Request) {
	todoList := todoList{}
	rows, err := db.Query("SELECT id, schedule, timelimit FROM trn_todo")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	for rows.Next() {
		todo := todo{}
		err := rows.Scan(&todo.ID, &todo.Schedule, &todo.TimeLimit)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		todoList = append(todoList, todo)
	}
	defer rows.Close()

	JSONTodoList, err := json.Marshal(todoList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(JSONTodoList)
}

func remove(w http.ResponseWriter, r *http.Request) {
	form := todo{}
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	targetID := form.ID

	stmt, err := db.Prepare("DELETE FROM trn_todo WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(targetID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}