package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

//注意：不能用&和&&，否则取IP时会取不到后面拼接的参数!!!
//payload1：curl http://127.0.0.1:2333/commandi\?ip\=1\|whoami
//payload2：curl http://127.0.0.1:2333/commandi\?ip\=1\|\|whoami

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
	fmt.Println(ip)
	cmd := exec.Command("bash", "-c", cmdStr)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, string(out))
}

func main() {
	http.HandleFunc("/commandi", commandInjectionPrepareHandler)
	http.ListenAndServe(":2333", nil)
}
