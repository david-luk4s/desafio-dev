package web

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/david-luk4s/desafio-dev/application"
)

var (
	PATH_TEMPLATE = "adapters/interfaces/web/templates/"
)

func makeTemlate(base string) *template.Template {
	files := []string{base}
	return template.Must(template.ParseFiles(files...))
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, PATH_TEMPLATE+"index.html")
}

func FormUploadFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, PATH_TEMPLATE+"form.html")
}

func ListOperations(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	items, _ := application.ListTransaction(ctx)

	makeTemlate(PATH_TEMPLATE+"list.html").Execute(w, items)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if r.Method != "POST" {
		fmt.Fprintf(w, "allowed not method")
		return
	}

	file, _, err := r.FormFile("myfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Process full transaction
	application.ProcessTransaction(ctx, file)

	http.ServeFile(w, r, PATH_TEMPLATE+"result.html")
}
