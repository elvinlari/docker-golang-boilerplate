package servid

import (
	"net/http"
	"os"

	dtbase "github.com/elvinlari/docker-golang/internal/platform/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	taskModel "github.com/elvinlari/docker-golang/internal/task/model"
	taskHttp "github.com/elvinlari/docker-golang/internal/task/http"
	companyModel "github.com/elvinlari/docker-golang/internal/company/model"
	companyHttp "github.com/elvinlari/docker-golang/internal/company/http"
	userModel "github.com/elvinlari/docker-golang/internal/user/model"
	userHttp "github.com/elvinlari/docker-golang/internal/user/http"
	inviteModel "github.com/elvinlari/docker-golang/internal/invite/model"
	inviteHttp "github.com/elvinlari/docker-golang/internal/invite/http"
)


func config() {
	logger()
}

func logger() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

// App Instance which contains router and db
type App struct {
	*http.Server
	r  *gin.Engine
}

// NewApp creates new App with db connection pool
func NewApp() *App {
	config()

	// gorm
	db, err := dtbase.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	dtbase.RunMigration(db)

	// Set Gin mode to release
    gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// task domain
	tsDB := &taskModel.TaskService{DB: db}
	tsHTTP := taskHttp.TaskService{
		Service: tsDB,
	}
	taskHttp.RegisterRoutes(router, &tsHTTP)

	// company domain
	cmDB := &companyModel.Service{DB: db}
	cmHTTP := companyHttp.Service{
		Service: cmDB,
	}
	companyHttp.RegisterRoutes(router, &cmHTTP)

	// user domain
	usrDB := &userModel.Service{DB: db}
	usrHTTP := userHttp.Service{
		Service: usrDB,
	}
	userHttp.RegisterRoutes(router, &usrHTTP)

	// invite domain
	intDB := &inviteModel.Service{DB: db}
	intHTTP := inviteHttp.Service{
		Service: intDB,
	}
	userHttp.RegisterRoutes(router, &intHTTP)

	// show all routes
	showRoutes(router)

	server := &App{
		r:  router,
	}
	return server
}

// Start launching the server
func (a *App) Start() {
	log.Fatal(http.ListenAndServe(os.Getenv("GO_PORT"), a.r))
}

func showRoutes(r *gin.Engine) {
    log.Info("registered routes: ")
    
    // Iterate over registered routes
    routes := r.Routes()
    for _, route := range routes {
        log.Infof("%s %s\n", route.Method, route.Path)
    }
}

