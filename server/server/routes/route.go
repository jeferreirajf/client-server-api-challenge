package route

import "net/http"

type Route interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
