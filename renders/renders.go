package renders

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderServerErrorTemplate(w http.ResponseWriter, statusCode int, errMsg string) {
	tmpl := `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>error page</title>
</head>
<body>
    <h1>this is an error</h1>
    
</body>
</html>`

	t, err := template.New("error").Parse(tmpl)
	if err != nil {
		http.Error(w, fmt.Sprintf("internal error %v", err), http.StatusInsufficientStorage)
	}
	data := struct {
		StatusCode int
		Error      string
	}{
		StatusCode: statusCode,
		Error:      errMsg,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)
	t.Execute(w, data)

}
