package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/mahirB/user_managment_api/controllers"
	"github.com/mahirB/user_managment_api/middlewares"

	"github.com/julienschmidt/httprouter"
	"github.com/mahirB/user_managment_api/services"
	"github.com/urfave/negroni"
)

var (
	user  controllers.Users
	perms controllers.Permissions

	deploy = kingpin.Flag("deploy", "Create tables").Short('d').Bool()
)

func init() {
	kingpin.Parse()
}

func main() {

	// Configuration provider
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current folder")
	}

	configPath := path + "/conf.yaml"

	fmt.Println("Reading configuration data from:")
	fmt.Println(configPath)

	services.NewConfigurer(configPath)

	// DB access and renderer
	services.NewAccess(services.Configuration.DB)
	services.NewRenderer()

	// Run it for Table deployment
	if *deploy {
		controllers.Deploy()
		return
	}

	// Multiplexer
	server := negroni.Classic()
	mux := httprouter.New()

	// Middlewares
	server.Use(negroni.HandlerFunc(middlewares.CORS))
	server.Use(negroni.HandlerFunc(middlewares.Preflight))

	// Prevent sending of Call Stack on panic
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	server.Use(recovery)

	server.UseHandler(mux)

	// Health check
	mux.GET("/health", health)

	// User endpoints
	mux.GET("/users", user.List)
	mux.GET("/user/:id", user.Get)
	mux.POST("/user", user.Create)
	mux.DELETE("/user/:id", user.Delete)
	mux.PUT("/user", user.Update)

	// Permissions endpoints
	mux.GET("/permissions", perms.List)
	mux.GET("/permissions/:id", perms.Get)
	mux.POST("/permissions/:id", perms.Update)

	server.Run(services.Configuration.IP + ":" + services.Configuration.Port)
}

func health(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	services.Renderer.Render(res, http.StatusOK, map[string]interface{}{"health": "UP"})
}
