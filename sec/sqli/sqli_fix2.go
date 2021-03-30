package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string
	Score    int
}

func sqlInjectionStringHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	selectTemplate := `select id, username, score from users where username=?`
	checkErr(err)

	vars := r.URL.Query()
	tmp, ok := vars["username"]
	if !ok {
		fmt.Fprint(w, "username is empty")
	}
	username := tmp[0]
	stmt, err := db.Prepare(selectTemplate)
	if err != nil {
		fmt.Printf("prepare sql failed, err:%v\n", err)
		return
	}
	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	//sql := fmt.Sprintf(selectTemplate, username)
	//fmt.Println(sql)
	//if db == nil {
	//	fmt.Fprint(w, "db is nil")
	//}
	//rows, err := db.Query(sql)
	//if err != nil {
	//	fmt.Fprint(w, err)
	//	return
	//}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		var id int
		var username string
		var score int
		err = rows.Scan(&id, &username, &score)
		if err != nil {
			panic(err)
		}
		users = append(users, User{
			Username: username,
			Score:    score,
		})
	}
	fmt.Fprint(w, users)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/sqli/string", sqlInjectionStringHandler)
	// 端口签名需要有冒号！
	http.ListenAndServe(":2333", nil)
}
