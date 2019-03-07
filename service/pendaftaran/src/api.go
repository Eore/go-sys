package pendaftaran

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ApiRoute() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "miawasaa")
	})
	return router
}
