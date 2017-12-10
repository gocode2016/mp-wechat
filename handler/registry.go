package handler

import (
	"net/http"
)

func Register() {
	http.HandleFunc("/", root)
}
