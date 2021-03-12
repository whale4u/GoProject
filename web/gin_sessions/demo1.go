package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

//初始化一个cookie存储对象
var store = sessions.NewCookieStore([]byte("make-a-difference"))

func main() {
	http.HandleFunc("/save", SaveSession)
	http.HandleFunc("/get", GetSession)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP Server failed, err: ", err)
		return
	}
}

func SaveSession(w http.ResponseWriter, r *http.Request) {
	//获取一个session对象，session-name是ssesion的名字
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//在session中存储值
	session.Values["foo"] = "bar"
	session.Values[42] = 43

	//
	session.Save(r, w)

	fmt.Println(session.Values)
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	foo := session.Values["foo"]
	// 为什么输出nil ？？？？
	fmt.Println(foo)
}

// 删除
// 将session的最大存储时间设置为小于零的数即为删除
//session.Options.MaxAge = -1
//session.Save(r, w)
