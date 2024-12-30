package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	tokenList []string
	mutex     sync.Mutex
)

func logRequest(r *http.Request) {
	log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	token := r.URL.Query().Get("token")
	if token != "" {
		mutex.Lock()
		tokenList = append(tokenList, token)
		mutex.Unlock()
		fmt.Fprintln(w, "successful")
	} else {
		http.Error(w, "Token is required", http.StatusBadRequest)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	mutex.Lock()
	defer mutex.Unlock()

	if len(tokenList) > 0 {
		getOne := tokenList[0]
		tokenList = tokenList[1:]
		fmt.Fprintln(w, getOne)
	} else {
		fmt.Fprintln(w, "None")
	}
}

func totalHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	mutex.Lock()
	defer mutex.Unlock()

	fmt.Fprintf(w, "%v\n\nTotal: %d", tokenList, len(tokenList))
}

func main() {
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/total", totalHandler)

	fmt.Println("Server is running on port 5000...")
	http.ListenAndServe(":5000", nil)
}
