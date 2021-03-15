package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func commandInjectionPrepareHandler(w http.ResponseWriter, r *http.Request) {
	// vars is map type
	vars := r.URL.Query()
	tmpKey, ok := vars["ip"]
	if !ok {
		fmt.Println(w, "ip is empty")
		return
	}
	ip := tmpKey[0]

	cmdStr := "ping -c 3 " + ip
	cmd := exec.Command("bash", "-c", cmdStr)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, string(out))
}

func main() {
	http.HandleFunc("/command", commandInjectionPrepareHandler)
	http.ListenAndServe(":8080", nil)
}
