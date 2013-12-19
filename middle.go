package gosearch

import (
	"github.com/codegangsta/martini"
	"net/http"
)

func SearcherMiddle(c martini.Context, w http.ResponseWriter, r *http.Request) {
	switch r.URL.Query().Get("type") {
	case "bing":
		c.MapTo(BingSearcher{}, (*Searcher)(nil))
		w.Header().Set("Content-Type", "application/json")
	default:
		c.MapTo(BingSearcher{}, (*Searcher)(nil))
		w.Header().Set("Content-Type", "application/json")
	}
}
