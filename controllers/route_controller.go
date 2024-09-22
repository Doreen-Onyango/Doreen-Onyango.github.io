package controllers

import (
	"net/http"

	"github.com/Doreen-Onyango/portfolio.git/utils"
)

func Welcomecontroller(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "welcome_page.html", nil)
}
