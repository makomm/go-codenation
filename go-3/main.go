package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"encoding/json"
	"strings"
	"math/rand"
)

//PORT port to be used
const PORT = "8080"
var db, err = sql.Open("sqlite3", "./database.sqlite")



type response struct {
	Actor string `json:"actor"`
	Detail string `json:"quote"`
}
func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.Handle("/v1/quote", quote()).Methods("GET", "OPTIONS")
	r.Handle("/v1/quote/{actor}", quoteByActor()).Methods("GET", "OPTIONS")
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + PORT,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func quote() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var res response
		randomN := rand.Intn(5)
		rows, err := db.Query("SELECT actor, detail FROM scripts WHERE actor IS NOT NULL AND detail IS NOT NULL")
		if err != nil {
			panic(err)
		}
		rows.Next()
		for i:=0; i < randomN; i++{
			rows.Next()
		}
		rows.Scan(&res.Actor, &res.Detail)
		fmt.Println(res, randomN)
		response, _ := json.Marshal(res)
		w.Write(response)
	})
}

func quoteByActor() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var res response
		randomN := rand.Intn(5)
		actor := strings.Replace(params["actor"], "+", " ", 1)
		rows, err := db.Query("SELECT actor, detail FROM scripts WHERE actor = $1 AND detail IS NOT NULL" , actor)
		if err != nil {
			panic(err)
		}
		rows.Next()
		for i:=0; i < randomN; i++{
			rows.Next()
		}
		rows.Scan(&res.Actor, &res.Detail)
		response, _ := json.Marshal(res)
		w.Write(response)
	})
}
