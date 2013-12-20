package gosearch

import (
	"net/http"
)

func SearchHandler(r *http.Request, w http.ResponseWriter, s Searcher) (int, string) {
	return s.Search(r, w)
}
