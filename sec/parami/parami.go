package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

func paramInjectionPrepareHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Body1: ", r.Body)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, "parse form error")
	}

	fmt.Println("Body2: ", r.Body)
	argMap := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&argMap)
	fmt.Println("argMap: ", argMap)

	command := "curl"
	cmdArgs := []string{"-v", "--url", "http://www.baidu.com"}

	for argKey, argValue := range argMap {
		arg := "--" + argKey
		cmdArgs = append(cmdArgs, arg, argValue.(string))
		fmt.Println(argKey, "-->", argValue)
	}

	cmd := exec.Command(command, cmdArgs...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, string(output))
	}
	return
}

func main() {
	http.HandleFunc("/parami/", paramInjectionPrepareHandler)
	// 端口签名需要有冒号！
	http.ListenAndServe(":2333", nil)
}
