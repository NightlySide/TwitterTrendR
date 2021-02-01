package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/trends", fetchTrends)
    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func fetchTrends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: /trends")
	
	client, err := FromDotEnv()
	if err != nil {
		panic(err)
	}
	
	//search(client, "Kitty")
	trends := trendList(client)[:20]
	/*wg := sync.WaitGroup{}
	wg.Add(len(trends))
	for _, trend := range trends {
		go trendAnalysis(client, trend, &wg)
	}
	wg.Wait()*/

    json.NewEncoder(w).Encode(trends)
}