package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/Doreen-Onyango/portfolio.git/renders"
)

var (
	TemplateCache map[string]*template.Template
	fn            = template.FuncMap{}
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, ok := TemplateCache[tmpl]
	if !ok {
		renders.RenderServerErrorTemplate(w, http.StatusNotFound, fmt.Sprintf("template %s not found", tmpl))
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.Execute(w, data)
	if err != nil {
		renders.RenderServerErrorTemplate(w, http.StatusInternalServerError, fmt.Sprintf("Error rendering template: %v", err))
	}
}
func LoadTemplates() error {
	cache := map[string]*template.Template{}
	baseDir := GetProjectRoot("sections", "templates")

	tempDir := filepath.Join(baseDir, "*.page.html")
	pages, err := filepath.Glob(tempDir)
	if err != nil {
		return fmt.Errorf("error globbing templates: %v", err)
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(fn).ParseFiles(page)
		if err != nil {
			return fmt.Errorf("error parsing: %s %v ", name, err)
		}
		// layoutPath := filepath.Join(baseDir, "*.layout.html")
		// matches, err := filepath.Glob(layoutPath)
		// if err != nil {
		// 	return fmt.Errorf("error finding layout files %v", err)
		// }

		// if len(matches) > 0 {
		// 	ts, err = ts.ParseGlob(layoutPath)
		// 	if err != nil {
		// 		return fmt.Errorf("error parsing layout files: %v", err)
		// 	}
		// }
		cache[name] = ts
	}
	TemplateCache = cache
	return nil
}
