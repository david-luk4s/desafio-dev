package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/david-luk4s/desafio-dev/adapters/interfaces/web"
)

func Handler() {
	fmt.Println("server listen in *:8080...")

	//Web routers
	http.HandleFunc("/", web.Home)
	http.HandleFunc("/web/form", web.FormUploadFile)
	http.HandleFunc("/web/upload", web.UploadFile)
	http.HandleFunc("/web/list", web.ListOperations)

	//Api routers
	http.HandleFunc("/api/upload", upload)
	http.HandleFunc("/api/list", list)
	log.Panic(http.ListenAndServe("0.0.0.0:8080", nil))
}
