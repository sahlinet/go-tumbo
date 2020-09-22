package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/sahlinet/go-tumbo/pkg/app/handler"
	"github.com/sahlinet/go-tumbo/pkg/config"
	"github.com/sahlinet/go-tumbo/pkg/models"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	// Database
	db, err := gorm.Open(sqlite.Open(config.DB.Name), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = gin.Default()
	a.setRouters()
}

// Wrap the router for GET method
/*func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
*/

/*// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.Use(handler.BasicAuth)
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}
*/

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Router.GET("/workers", a.GetAllWorkers)
	//a.Router.GET("/workers", a.CreateWorker)
}

// Handlers to manage workers
func (a *App) GetAllWorkers(c *gin.Context) func(*gin.Context) {
	return handler.GetAllWorkers(a.DB, c)
}

func (a *App) CreateWorker(w http.ResponseWriter, r *http.Request) {
	handler.CreateWorker(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
