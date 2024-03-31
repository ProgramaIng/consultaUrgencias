package server

import (
	"net/http"

	"github.com/consultaUrgencias/internal/handlers"
	"github.com/gin-gonic/gin"
)

type App struct {
	ping handlers.Interface
}

// BuildApplication ...
func BuildApplication() *App {
	return &App{ping: handlers.New()}

}

// URLMapping inicializa las rutas para consultar en el navegador
func mapRoutes(a *App) *gin.Engine {

	router := gin.Default()
	mainRouter := router.Group("/")
	mainRouter.GET("/ping", a.ping.PingHandler)

	return router
}

// Start metodo para que en el main levante el servidor web local
func Start() {

	constructor := BuildApplication()
	Router := mapRoutes(constructor)
	StartServer(Router)
}

// StartServer Inicialización y configuracion del servidor
func StartServer(engine *gin.Engine) {

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	srv.ListenAndServe()
}
