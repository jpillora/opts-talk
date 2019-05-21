package search

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jpillora/opts"
)

type cmd struct {
	Term term
}

func New() opts.Opts {
	return opts.New(&cmd{})
}

type term string

func (term) Complete(s string) []string {
	resp, err := http.Get("https://www.google.com/complete/search?client=psy-ab&q=" + s)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	results := []interface{}{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil
	}
	log.Printf("%+v", results)

	return nil
}
