package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

func paramInjectionPrepareHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, "parse form error")
	}

	argMap := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&argMap)

	command := "curl"
	cmdArgs := []string{"-v", "--url", "http://www.baidu.com"}

	for argKey, argValue := range argMap {
		arg := "--" + argKey
		cmdArgs = append(cmdArgs, arg, argValue.(string))
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
	http.HandleFunc("/param/", paramInjectionPrepareHandler)
	http.ListenAndServe("2333", nil)
}
