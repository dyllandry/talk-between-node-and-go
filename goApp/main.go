package main

import (
	"net/http"
  "fmt"
	"log"
	"github.com/joho/godotenv"
	"os"
	"io/ioutil"
	"time"
)

func main() {
	loadEnvFile()
	go runEvery2Seconds(hitNodeServer)
	startServer()
}

func loadEnvFile() {
	godotenv.Load()
}

func runEvery2Seconds(f func())  {
	for {
		time.Sleep(2 * time.Second)
		f()
	}
}

func hitNodeServer() {
	nodePort, _ := os.LookupEnv("nodePort")
	nodeURL := "http://127.0.0.1:" + nodePort

	resp, err := http.Get(nodeURL)
	log.Println("Sent request to node.")
	if (err != nil) {
		log.Printf("Request to node server (%v) failed.\n", nodeURL)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if (err != nil) {
		log.Printf("Could not read response body: %v\n", nodeURL)
	}

	log.Println("Recieved from node: " + string(body))
}

func startServer() {
	goPort, _ := os.LookupEnv("goPort")

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from go!")
	})

	log.Printf("Go server listening on http://localhost:%v", goPort)
	http.ListenAndServe(":" + goPort, nil)
}
