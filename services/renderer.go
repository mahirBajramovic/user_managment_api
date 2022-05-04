package services

import (
	"log"
	"net/http"

	"github.com/unrolled/render"
)

// Renderer ->
var Renderer *RendererCtrl

// RendererCtrl ->
type RendererCtrl struct {
	r *render.Render
}

// NewRenderer ...
func NewRenderer() *RendererCtrl {
	rend := new(RendererCtrl)

	rend.r = render.New(render.Options{
		IndentJSON:    true,
		UnEscapeHTML:  true,
		IsDevelopment: true,
	})

	Renderer = rend
	return rend
}

// HTML ->
func (rend *RendererCtrl) HTML(res http.ResponseWriter, name string, v interface{}) {
	rend.r.HTML(res, http.StatusOK, name, v)
}

// Render ->
func (rend *RendererCtrl) Render(res http.ResponseWriter, status int, v interface{}) {
	res.Header().Set("Access-Control-Allow-Origin", "*")

	if rend == nil {
		log.Println("Renderer control not inicialized")
		return
	}

	if rend.r == nil {
		log.Println("Renderer not inicialized")
		return
	}

	rend.r.JSON(res, status, v)
}
