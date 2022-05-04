package middlewares

import (
	"net/http"

	"github.com/mahirB/user_managment_api/services"
)

// CORS ...
func CORS(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	// CORS support for Preflighted requests
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

	next(res, req)
}

// Preflight ...
func Preflight(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	if req.Method == "OPTIONS" {
		services.Renderer.Render(res, http.StatusOK, map[string]string{"status": "OK"})
		return
	}

	next(res, req)
}
