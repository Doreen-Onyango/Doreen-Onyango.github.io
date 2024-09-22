package routers

import (
	"net/http"
	"strings"

	"github.com/Doreen-Onyango/portfolio.git/controllers"
	"github.com/Doreen-Onyango/portfolio.git/utils"
)

func Routes(mux *http.ServeMux) {
	dir := utils.GetProjectRoot("sections", "static")
	fServer := http.FileServer(http.Dir(dir))
	mux.Handle("/static", http.StripPrefix("/static", fServer))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.Welcomecontroller(w, r)
	})
}
func RouteChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/static/") {
			next.ServeHTTP(w, r)
		}
		next.ServeHTTP(w, r)
	})

}
