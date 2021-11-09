package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	userhttp "github.com/DarkSoul94/helpdesk2/pkg_user/delivery/http"
	grouprepo "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager/standart/repo/mysql"
	userrepo "github.com/DarkSoul94/helpdesk2/pkg_user/repo/mysql"
	userusecase "github.com/DarkSoul94/helpdesk2/pkg_user/usecase"

	"github.com/DarkSoul94/helpdesk2/auth"
	authhttp "github.com/DarkSoul94/helpdesk2/auth/delivery/http"
	authusecase "github.com/DarkSoul94/helpdesk2/auth/usecase"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required
	"github.com/spf13/viper"
	gomrmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// App ...
type App struct {
	userUC   pkg_user.UserManagerUC
	userRepo pkg_user.UserManagerRepo

	authUC auth.AuthUC

	httpServer *http.Server
}

// NewApp ...
func NewApp() *App {
	db := initDB()

	userRepo := userrepo.NewRepo(db)
	grpRepo := grouprepo.NewGroupRepo(db)
	userUC := userusecase.NewUsecase(userRepo, grpRepo)

	authUC := authusecase.NewUsecase(userUC, viper.GetString("app.auth.secret_key"), []byte(viper.GetString("app.auth.signing_key")), viper.GetDuration("app.auth.ttl"))

	return &App{
		userRepo: userRepo,
		userUC:   userUC,

		authUC: authUC,
	}
}

// Run run helpdesklication
func (a *App) Run(port string) error {
	defer a.userRepo.Close()

	router := gin.New()
	if viper.GetBool("app.release") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		router.Use(gin.Logger())
	}
	router.Use(
		gin.RecoveryWithWriter(logger.GetOutFile()),
	)

	apiRouter := router.Group("/helpdesk")
	authhttp.RegisterHTTPEndpoints(apiRouter, a.authUC)
	authMiddlware := authhttp.NewAuthMiddleware(a.authUC)

	userhttp.RegisterHTTPEndpoints(apiRouter, a.userUC, authMiddlware)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	var l net.Listener
	var err error
	if viper.GetBool("helpdesk.sock_mode") {
		sockName := viper.GetString("helpdesk.sock_name")
		os.Remove(sockName)
		l, err = net.Listen("unix", sockName)
		if err != nil {
			panic(err)
		}
		defer l.Close()
		os.Chmod(sockName, 0664)
	} else {
		l, err = net.Listen("tcp", a.httpServer.Addr)
		if err != nil {
			panic(err)
		}
	}
	go func(l net.Listener) {
		if err := a.httpServer.Serve(l); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}(l)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *sql.DB {
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		viper.GetString("app.db.login"),
		viper.GetString("app.db.pass"),
		viper.GetString("app.db.host"),
		viper.GetString("app.db.port"),
		viper.GetString("app.db.name"),
		viper.GetString("app.db.args"),
	)
	db, err := sql.Open(
		"mysql",
		dbString,
	)
	if err != nil {
		panic(err)
	}
	runMigrations(db)
	return db
}

func runMigrations(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		viper.GetString("helpdesk.db.name"),
		driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange && err != migrate.ErrNilVersion {
		fmt.Println(err)
	}
}

func initGormDB() *gorm.DB {
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		viper.GetString("app.db.login"),
		viper.GetString("app.db.pass"),
		viper.GetString("app.db.host"),
		viper.GetString("app.db.port"),
		viper.GetString("app.db.name"),
		viper.GetString("app.db.args"),
	)
	db, err := gorm.Open(gomrmysql.Open(dbString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func runGormMigrations(db *gorm.DB) {
	// Migrate the schema
	// Add links to needed models
	db.AutoMigrate()
}
