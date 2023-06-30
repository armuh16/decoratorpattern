package main

import (
	"fmt"
	"net/http"
)

type DB interface {
	Store(string) error
}

type Store struct{}

func (s *Store) Store(value string) error {
	fmt.Println("storing to db", value)
	return nil
}

func myExecuteFunc(db DB) Executefn {
	return func(s string) {
		fmt.Println("INI FILE TEST -- ", s)
		db.Store(s)
	}
}

func makeHTTPFunc(db DB, fn httpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(db, w, r); err != nil {

		}
		// where is my db??
	}
}

func main() {
	s := &Store{}
	http.HandleFunc("/", makeHTTPFunc(s, handler))
	Execute(myExecuteFunc(s))
}

func handler(db DB, w http.ResponseWriter, r *http.Request) error {
	return nil
}

type httpFunc func(db DB, w http.ResponseWriter, r *http.Request) error

// From third party
type Executefn func(string)

func Execute(fn Executefn) {
	fn("TEST")
}
