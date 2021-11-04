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

	"github.com/DarkSoul94/helpdesk2/helpdesk"
	helpdeskhttp "github.com/DarkSoul94/helpdesk2/helpdesk/delivery/http"
	helpdeskrepo "github.com/DarkSoul94/helpdesk2/helpdesk/repo/mock"
	helpdeskusecase "github.com/DarkSoul94/helpdesk2/helpdesk/usecase"
	"github.com/DarkSoul94/helpdesk2/pkg/logger"
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
	helpdeskUC      helpdesk.Usecase
	helpdeskRepo    helpdesk.Repository
	httpServer *http.Server
}

// NewApp ...
func NewApp() *App {
	repo := helpdeskrepo.NewRepo()
	uc := helpdeskusecase.NewUsecase(repo)
	return &App{
		helpdeskUC:   uc,
		helpdeskRepo: repo,
	}
}

// Run run helpdesklication
func (a *App) Run(port string) error {
	defer a.helpdeskRepo.Close()
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(
		gin.RecoveryWithWriter(logger.GetOutFile()),
	)

	helpdeskhttp.RegisterHTTPEndpoints(router, a.helpdeskUC)

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
		viper.GetString("helpdesk.db.login"),
		viper.GetString("helpdesk.db.pass"),
		viper.GetString("helpdesk.db.host"),
		viper.GetString("helpdesk.db.port"),
		viper.GetString("helpdesk.db.name"),
		viper.GetString("helpdesk.db.args"),
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
		viper.GetString("helpdesk.db.login"),
		viper.GetString("helpdesk.db.pass"),
		viper.GetString("helpdesk.db.host"),
		viper.GetString("helpdesk.db.port"),
		viper.GetString("helpdesk.db.name"),
		viper.GetString("helpdesk.db.args"),
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
