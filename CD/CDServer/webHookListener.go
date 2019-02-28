package CDServer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	var info myInfo
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(bytesBody, &info)
	// fmt.Println(body)
	fmt.Println(info)
	w.WriteHeader(200)
}
func StartServer() {
	http.HandleFunc("/echo/", echo)
	http.HandleFunc("/restart", restartApp)
	if err := http.ListenAndServe(":1337", nil); err != nil {
		panic(err)
	}
}
