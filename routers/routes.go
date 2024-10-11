package routers

import (
	"net/http"
	"strings"

	"github.com/Doreen-Onyango/portfolio.git/controllers"
	"github.com/Doreen-Onyango/portfolio.git/utils"
)

var rts = map[string]bool{
	"/":          true,
	"/contacts":  true,
	"/portfolio": true,
	"/skills":    true,
}

func Routes(mux *http.ServeMux) {
	dir := utils.GetProjectRoot("pages", "static")
	fServer := http.FileServer(http.Dir(dir))
	mux.Handle("/static/", http.StripPrefix("/static/", fServer))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.Welcomecontroller(w, r)
	})

	mux.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		controllers.ContactsController(w, r)
	})

	mux.HandleFunc("/portfolio", func(w http.ResponseWriter, r *http.Request) {
		controllers.PortfolioController(w, r)
	})

	mux.HandleFunc("/skills", func(w http.ResponseWriter, r *http.Request) {
		controllers.SkillController(w, r)
	})
}

func RouteChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/static/") {
			next.ServeHTTP(w, r)
			return
		}

		if _, ok := rts[r.URL.Path]; !ok {
			controllers.NotFoundController(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
