package sample1

import "net/http"

func init() {
	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/admin/", adminHandler)
}
