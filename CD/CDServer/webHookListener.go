package CDServer

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/echo/")
	message = message
	w.Write([]byte(message))
}

func restartApp(w http.ResponseWriter, r *http.Request) {
	// out, err := exec.Command("bash", "-c", "ps -e | grep -w 'go'| awk '{print $1}'").Output()
	out, err := exec.Command("bash", "-c", "kill $(lsof -t -i :9000)").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output is %s\n", out)
	w.WriteHeader(200)
}
func StartServer() {
	http.HandleFunc("/echo/", echo)
	http.HandleFunc("/restart", restartApp)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
