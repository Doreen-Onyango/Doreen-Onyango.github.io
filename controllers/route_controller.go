package controllers

import (
	"net/http"

	"github.com/Doreen-Onyango/portfolio.git/utils"
)

func Welcomecontroller(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "index.page.html", nil)
}

func NotFoundController(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "404.page.html", nil)
}
