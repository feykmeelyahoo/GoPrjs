package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	podHostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Println("hostname:", podHostname)

	myKubeNode := getEnv("KUBENODE", "defaulthost")

	fmt.Println(myKubeNode)
	// for _, e := range os.Environ() {
	// 	pair := strings.Split(e, "=")
	// 	fmt.Println(pair[0])
	// }
	fmt.Fprintf(w, "Hostname of The NODE %s, Hostname of the Pod %s!", myKubeNode, podHostname)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
