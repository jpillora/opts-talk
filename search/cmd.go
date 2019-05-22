package search

import (
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"

	"github.com/jpillora/opts"
)

type cmd struct {
	Term term
}

func New() opts.Opts {
	return opts.New(&cmd{}).Summary("search google")
}

type term string

func (term) Complete(s string) []string {
	if s == "" {
		return nil
	}
	v := url.Values{"client": []string{"client=psy-ab"}, "q": []string{s}}
	resp, err := http.Get("https://www.google.com/complete/search?" + v.Encode())
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	results := []interface{}{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil
	}
	if len(results) < 2 {
		return nil
	}
	list, ok := results[1].([]interface{})
	if !ok {
		return nil
	}
	bold := regexp.MustCompile(`</?b>`)
	strs := []string{}
	for _, item := range list {
		if m, ok := item.([]interface{}); ok {
			s, _ := m[0].(string)
			strs = append(strs, bold.ReplaceAllString(s, ""))
		}
	}
	return strs
}
