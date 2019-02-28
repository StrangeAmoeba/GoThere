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
	_, err := exec.Command("deploy-script").Output()
	// fmt.Println("2")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(body)
	w.WriteHeader(200)
}
func StartServer() {
	http.HandleFunc("/echo/", echo)
	http.HandleFunc("/restart", restartApp)
	if err := http.ListenAndServe(":1337", nil); err != nil {
		panic(err)
	}
}
