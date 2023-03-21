package api

import (
	"fmt"
	"log"
	"net/http"
)

func Handler() {
	fmt.Println("server listen in *:8080...")

	//Api routers
	http.HandleFunc("/api/upload", upload)
	http.HandleFunc("/api/list", list)
	log.Panic(http.ListenAndServe("0.0.0.0:8080", nil))
}
