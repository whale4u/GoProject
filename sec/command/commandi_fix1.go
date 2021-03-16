package main

import (
	"fmt"
	"net"
	"net/http"
	"os/exec"
)

//校验IP格式
func IsIpv4Net(host string) bool {
	return net.ParseIP(host) != nil
}

func commandInjectionPrepareHandler(w http.ResponseWriter, r *http.Request) {
	// vars is map type
	vars := r.URL.Query()
	tmpKey, ok := vars["ip"]
	if !ok {
		fmt.Println(w, "ip is empty")
		return
	}
	ip := tmpKey[0]

	if IsIpv4Net(ip) {
		cmdStr := "ping -c 3 " + ip
		cmd := exec.Command("bash", "-c", cmdStr)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprint(w, err)
		}
		fmt.Fprint(w, string(out))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("IP格式错误！")
	}
}

func main() {
	http.HandleFunc("/command", commandInjectionPrepareHandler)
	http.ListenAndServe(":2333", nil)
}
