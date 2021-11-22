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
	"github.com/DarkSoul94/helpdesk2/pkg_support"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	userhttp "github.com/DarkSoul94/helpdesk2/pkg_user/delivery/http"
	userrepo "github.com/DarkSoul94/helpdesk2/pkg_user/repo/mysql"
	userusecase "github.com/DarkSoul94/helpdesk2/pkg_user/usecase"

	"github.com/DarkSoul94/helpdesk2/pkg_ticket"
	tickethttp "github.com/DarkSoul94/helpdesk2/pkg_ticket/delivery/http"
	ticketrepo "github.com/DarkSoul94/helpdesk2/pkg_ticket/repo/mysql"
	ticketusecase "github.com/DarkSoul94/helpdesk2/pkg_ticket/usecase"

	"github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager"
	catsecrepo "github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager/repo/mysql"
	catsecusecase "github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager/usecase"

	"github.com/DarkSoul94/helpdesk2/pkg_ticket/reg_fil_manager"
	regfilrepo "github.com/DarkSoul94/helpdesk2/pkg_ticket/reg_fil_manager/repo/mysql"
	regfilusecase "github.com/DarkSoul94/helpdesk2/pkg_ticket/reg_fil_manager/usecase"

	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
	grouprepo "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager/standart/repo/mysql"
	groupusecase "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager/standart/usecase/group"
	permusecase "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager/standart/usecase/permissions"

	supporthttp "github.com/DarkSoul94/helpdesk2/pkg_support/delivery/http"
	supportrepo "github.com/DarkSoul94/helpdesk2/pkg_support/repo/mysql"
	supportusecase "github.com/DarkSoul94/helpdesk2/pkg_support/usecase"

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
	groupRepo group_manager.IGroupRepo
	groupUC   group_manager.IGroupUsecase

	suppRepo pkg_support.ISupportRepo
	suppUC   pkg_support.ISupportUsecase

	permUC group_manager.IPermManager

	userUC   pkg_user.IUserUsecase
	userRepo pkg_user.IUserRepo

	authUC auth.AuthUC

	catSecRepo cat_sec_manager.ICatSecRepo
	catSecUC   cat_sec_manager.ICatSecUsecase

	regFilRepo reg_fil_manager.IRegFilRepo
	regFilUC   reg_fil_manager.IRegFilUsecase

	ticketRepo pkg_ticket.ITicketRepo
	ticketUC   pkg_ticket.ITicketUsecase

	httpServer *http.Server
}

// NewApp ...
func NewApp() *App {
	db := initDB()

	grpRepo := grouprepo.NewGroupRepo(db)
	grpUC := groupusecase.NewGroupManager(grpRepo)

	permUC := permusecase.NewPermManager(grpRepo)

	suppRepo := supportrepo.NewSupportRepo(db)
	suppUC := supportusecase.NewSupportUsecase(suppRepo, permUC)

	userRepo := userrepo.NewRepo(db)
	userUC := userusecase.NewUsecase(userRepo, grpUC, permUC, suppUC)

	authUC := authusecase.NewUsecase(userUC,
		viper.GetString("app.auth.secret_key"),
		[]byte(viper.GetString("app.auth.signing_key")),
		viper.GetDuration("app.auth.ttl"))

	catsecRepo := catsecrepo.NewCatSecRepo(db)
	catsecUC := catsecusecase.NewCatSecUsecase(catsecRepo)

	regfilRepo := regfilrepo.NewRegFilRepo(db)
	regfilUC := regfilusecase.NewRegFilUsecase(regfilRepo)

	ticketRepo := ticketrepo.NewTicketRepo(db)
	ticketUC := ticketusecase.NewTicketUsecase(ticketRepo, catsecUC, regfilUC, permUC)

	return &App{
		groupRepo: grpRepo,
		groupUC:   grpUC,

		suppRepo: suppRepo,
		suppUC:   suppUC,

		permUC: permUC,

		userRepo: userRepo,
		userUC:   userUC,

		authUC: authUC,

		catSecRepo: catsecRepo,
		catSecUC:   catsecUC,

		ticketRepo: ticketRepo,
		ticketUC:   ticketUC,
	}
}

// Run run helpdesklication
func (a *App) Run(port string) error {
	defer a.close()

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

	authMiddlware := authhttp.NewAuthMiddleware(a.authUC)
	authhttp.RegisterHTTPEndpoints(apiRouter, a.authUC)

	userMiddleware := userhttp.NewPermissionMiddleware(a.permUC)
	userhttp.RegisterHTTPEndpoints(apiRouter, a.userUC, authMiddlware, userMiddleware)

	supportMiddleware := supporthttp.NewPermissionMiddleware(a.permUC)
	supporthttp.RegisterHTTPEndpoints(apiRouter, a.suppUC, authMiddlware, supportMiddleware)

	ticketMiddleware := tickethttp.NewPermissionMiddleware(a.permUC)
	tickethttp.RegisterHTTPEndpoints(apiRouter, a.ticketUC, authMiddlware, ticketMiddleware)

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

func (a *App) close() {
	a.userRepo.Close()
	a.groupRepo.Close()
	a.suppRepo.Close()
	a.catSecRepo.Close()
	a.regFilRepo.Close()
	a.ticketRepo.Close()
}
