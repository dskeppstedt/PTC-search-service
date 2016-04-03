package ptc

import (
	"encoding/json"
	"net/http"
	"fmt"
    "gopkg.in/olivere/elastic.v3"
    "reflect"
)

//apiVersion is a simple endpoint handler that
//just writes a simple struct as json as the HTTP response.
func apiVersion(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		}{"PTC-Search-Service", "0.1"})
}

func getTweetsFromUserID(w http.ResponseWriter, r *http.Request) {
	var searchResult *elastic.SearchResult
	searchResult = elasticSearch.SearchTweetsFromID("100004471")
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
        if t, ok := item.(Tweet); ok {
            fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
            json.NewEncoder(w).Encode(
	            struct {
				UserID    string `json:"user_id"`
				Text 	string `json:"text"`
			}{t.User, t.Message})
        }
    }
	
}
