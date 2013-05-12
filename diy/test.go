package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	envs := os.Environ()
	for _, env := range envs {
		fmt.Fprintf(w, "%s\n", env)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(os.Getenv("OPENSHIFT_INTERNAL_IP")+":"+os.Getenv("OPENSHIFT_INTERNAL_PORT"), nil)
}
