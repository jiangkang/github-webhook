package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type HookConfig struct {
	PathRepo string `json:"path_repo"`
}

func main() {
	parseConfig()
	http.HandleFunc("/payload", handlePayload)
	err := http.ListenAndServe(":5555", nil)
	if err != nil {
		log.Fatal("Server start failed!")
	}
}

func parseConfig() {
	file, err := os.Open("hook_config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := HookConfig{}
	error := decoder.Decode(&conf)
	if error != nil {
		log.Fatal(error)
	}
	pathRepo = conf.PathRepo
}

// 绝对路径
var pathRepo string

func handlePayload(writer http.ResponseWriter, request *http.Request) {
	var event = request.Header.Get("X-Github-Event")
	switch event {
	case "push":
		exec.Command("cd", pathRepo)
		exec.Command("git", "pull")
		fmt.Println("git pull done!")
	}
}

