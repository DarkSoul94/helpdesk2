package server

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_support"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	userhttp "github.com/DarkSoul94/helpdesk2/pkg_user/delivery/http"
	userrepo "github.com/DarkSoul94/helpdesk2/pkg_user/repo/mysql"
	userusecase "github.com/DarkSoul94/helpdesk2/pkg_user/usecase"
	"github.com/jmoiron/sqlx"

	"github.com/DarkSoul94/helpdesk2/pkg_ticket"
	tickethttp "github.com/DarkSoul94/helpdesk2/pkg_ticket/delivery/http"
	ticketrepo "github.com/DarkSoul94/helpdesk2/pkg_ticket/repo/mysql"
	ticketucforsupp "github.com/DarkSoul94/helpdesk2/pkg_ticket/tickets_for_support"
	ticketusecase "github.com/DarkSoul94/helpdesk2/pkg_ticket/usecase"

	"github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager"
	catsecrepo "github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager/repo/mysql"
	catsecusecase "github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager/usecase"

	"github.com/DarkSoul94/helpdesk2/pkg_ticket/comment_manager"
	commentrepo "github.com/DarkSoul94/helpdesk2/pkg_ticket/comment_manager/repo/mysql"
	commentusecase "github.com/DarkSoul94/helpdesk2/pkg_ticket/comment_manager/usecase"

	"github.com/DarkSoul94/helpdesk2/pkg_ticket/reg_fil_manager"
	regfilrepo "github.com/DarkSoul94/helpdesk2/pkg_ticket/reg_fil_manager/repo/mysql"
	regfilusecase "github.com/DarkSoul94/helpdesk2/pkg_ticket/reg_fil_manager/usecase"

	"github.com/DarkSoul94/helpdesk2/pkg_ticket/file_manager"
	filerepo "github.com/DarkSoul94/helpdesk2/pkg_ticket/file_manager/repo/mysql"
	fileusecase "github.com/DarkSoul94/helpdesk2/pkg_ticket/file_manager/usecase"

	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
	grouprepo "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager/standart/repo/mysql"
	groupusecase "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager/standart/usecase/group"
	permusecase "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager/standart/usecase/permissions"

	supporthttp "github.com/DarkSoul94/helpdesk2/pkg_support/delivery/http"
	supportrepo "github.com/DarkSoul94/helpdesk2/pkg_support/repo/mysql"
	supportusecase "github.com/DarkSoul94/helpdesk2/pkg_support/usecase"

	"github.com/DarkSoul94/helpdesk2/pkg_scheduler"
	schedulerhttp "github.com/DarkSoul94/helpdesk2/pkg_scheduler/delivery/http"
	schedulerrepo "github.com/DarkSoul94/helpdesk2/pkg_scheduler/repo/mysql"
	schedulerusecase "github.com/DarkSoul94/helpdesk2/pkg_scheduler/usecase"

	"github.com/DarkSoul94/helpdesk2/pkg_consts"
	constshttp "github.com/DarkSoul94/helpdesk2/pkg_consts/delivery/http"
	constsrepo "github.com/DarkSoul94/helpdesk2/pkg_consts/repo/mysql"
	constsusecase "github.com/DarkSoul94/helpdesk2/pkg_consts/usecase"

	"github.com/DarkSoul94/helpdesk2/pkg_reports"
	reportshttp "github.com/DarkSoul94/helpdesk2/pkg_reports/delivery/http"
	reportsrepo "github.com/DarkSoul94/helpdesk2/pkg_reports/repo/mysql"
	reportsusecase "github.com/DarkSoul94/helpdesk2/pkg_reports/usecase"

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
	dbConnect *sql.DB

	groupRepo group_manager.IGroupRepo
	groupUC   group_manager.IGroupUsecase

	ticketUCForSupp pkg_ticket.IUCForSupport
	suppRepo        pkg_support.ISupportRepo
	suppUC          pkg_support.ISupportUsecase
	schedulerSupp   pkg_support.ISuppForScheduler

	permUC group_manager.IPermManager

	userUC   pkg_user.IUserUsecase
	userRepo pkg_user.IUserRepo

	authUC auth.AuthUC

	catSecRepo cat_sec_manager.ICatSecRepo
	catSecUC   cat_sec_manager.ICatSecUsecase

	regFilRepo reg_fil_manager.IRegFilRepo
	regFilUC   reg_fil_manager.IRegFilUsecase

	fileRepo file_manager.IFileRepo
	fileUC   file_manager.IFileUsecase

	commentRepo comment_manager.ICommentRepo
	commentUC   comment_manager.ICommentUsecase

	ticketRepo pkg_ticket.ITicketRepo
	ticketUC   pkg_ticket.ITicketUsecase

	schedulerRepo pkg_scheduler.ISchedulerRepo
	schedulerUC   pkg_scheduler.ISchedulerUsecase
	constsRepo    pkg_consts.IConstsRepo
	constsUC      pkg_consts.IConstsUsecase

	reportsRepo pkg_reports.IReportsRepo
	reportsUC   pkg_reports.IReportsUsecase

	httpServer *http.Server
}

// NewApp ...
func NewApp() *App {
	db := initDB()

	grpRepo := grouprepo.NewGroupRepo(db)
	suppRepo := supportrepo.NewSupportRepo(db)
	userRepo := userrepo.NewRepo(db)
	fileRepo := filerepo.NewFileRepo(db)
	catsecRepo := catsecrepo.NewCatSecRepo(db)
	regfilRepo := regfilrepo.NewRegFilRepo(db)
	ticketRepo := ticketrepo.NewTicketRepo(db)
	commentRepo := commentrepo.NewCommentRepo(db)
	schedulerRepo := schedulerrepo.NewShedulerRepo(db)
	constsRepo := constsrepo.NewConstsRepo(db)
	reportsRepo := reportsrepo.NewReportsRepo(db)

	grpUC := groupusecase.NewGroupManager(grpRepo)
	permUC := permusecase.NewPermManager(grpRepo)
	ticketUCForSupp := ticketucforsupp.NewTicketUCForSupport(ticketRepo)
	suppUC := supportusecase.NewSupportUsecase(suppRepo, permUC, ticketUCForSupp)
	schedulerSupp := supportusecase.NewSuppForSchedulerUsecase(suppRepo)
	userUC := userusecase.NewUsecase(userRepo, grpUC, permUC, suppUC)
	authUC := authusecase.NewUsecase(userUC,
		viper.GetString("app.auth.secret_key"),
		[]byte(viper.GetString("app.auth.signing_key")),
		viper.GetDuration("app.auth.ttl"))

	catsecUC := catsecusecase.NewCatSecUsecase(catsecRepo)
	regfilUC := regfilusecase.NewRegFilUsecase(regfilRepo)
	fileUC := fileusecase.NewFileUsecase(fileRepo)
	commentUC := commentusecase.NewCommentUsecase(commentRepo)
	ticketUC := ticketusecase.NewTicketUsecase(ticketRepo, catsecUC, regfilUC, fileUC, permUC, userUC, suppUC, commentUC)
	constsUC := constsusecase.NewConstsUsecase(constsRepo)

	schedulerUC := schedulerusecase.NewSchedulerUsecase(schedulerRepo, schedulerSupp)
	schedReports := schedulerusecase.NewShedulerForReports(schedulerRepo, constsUC, schedulerSupp)

	reportsUC := reportsusecase.NewReportsUsecase(reportsRepo, schedReports)

	return &App{
		dbConnect: db,

		groupRepo: grpRepo,
		groupUC:   grpUC,

		ticketUCForSupp: ticketUCForSupp,
		suppRepo:        suppRepo,
		suppUC:          suppUC,
		schedulerSupp:   schedulerSupp,

		permUC: permUC,

		userRepo: userRepo,
		userUC:   userUC,

		authUC: authUC,

		catSecRepo: catsecRepo,
		catSecUC:   catsecUC,

		regFilRepo: regfilRepo,
		regFilUC:   regfilUC,

		fileRepo: fileRepo,
		fileUC:   fileUC,

		commentRepo: commentRepo,
		commentUC:   commentUC,

		ticketRepo: ticketRepo,
		ticketUC:   ticketUC,

		schedulerRepo: schedulerRepo,
		schedulerUC:   schedulerUC,

		constsRepo: constsRepo,
		constsUC:   constsUC,

		reportsRepo: reportsRepo,
		reportsUC:   reportsUC,
	}
}

// Run run helpdesklication
func (a *App) Run(port string) error {
	if _, err := os.Stat(viper.GetString("app.store.path")); os.IsNotExist(err) {
		os.Mkdir(viper.GetString("app.store.path"), 0777)
	}

	defer a.close()

	router := gin.New()
	if viper.GetBool("app.release") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		router.Use(gin.Logger())
	}
	router.Use(
		gin.Recovery(),
		//gin.RecoveryWithWriter(logger.GetOutFile()),
	)

	apiRouter := router.Group("/helpdesk")

	authMiddlware := authhttp.NewAuthMiddleware(a.authUC)
	authhttp.RegisterHTTPEndpoints(apiRouter, a.authUC)

	userMiddleware := userhttp.NewPermissionMiddleware(a.permUC)
	userhttp.RegisterHTTPEndpoints(apiRouter, a.userUC, authMiddlware, userMiddleware)

	supportMiddleware := supporthttp.NewPermissionMiddleware(a.permUC)
	supporthttp.RegisterHTTPEndpoints(apiRouter, a.suppUC, authMiddlware, supportMiddleware)

	schedulerMiddleware := schedulerhttp.NewPermissionMiddleware(a.permUC)
	schedulerhttp.RegisterHTTPEndpoints(apiRouter, a.schedulerUC, authMiddlware, schedulerMiddleware)

	ticketMiddleware := tickethttp.NewPermissionMiddleware(a.permUC)
	tickethttp.RegisterHTTPEndpoints(apiRouter, a.ticketUC, authMiddlware, ticketMiddleware)

	constsMiddleware := constshttp.NewPermissionMiddleware(a.permUC)
	constshttp.RegisterHTTPEndpoints(apiRouter, a.constsUC, authMiddlware, constsMiddleware)

	reportshttp.RegisterHTTPEndpoints(apiRouter, a.reportsUC, authMiddlware)

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

	go moveFileFromDBtoFolder(a.dbConnect)

	ctx1, shutdown1 := context.WithCancel(context.Background())
	defer shutdown1()
	go a.ticketUC.DistributeTicket(ctx1)

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
	a.fileRepo.Close()
	a.commentRepo.Close()
	a.ticketRepo.Close()
	a.schedulerRepo.Close()
}

func moveFileFromDBtoFolder(db *sql.DB) {
	type dbFile struct {
		FileID        uint64         `db:"file_id"`
		FileName      string         `db:"file_name"`
		FileExtension string         `db:"file_extension"`
		TicketId      uint64         `db:"ticket_id"`
		FileData      sql.NullString `db:"file_data"`
		FileDate      time.Time      `db:"file_date"`
		Path          sql.NullString `db:"path"`
	}

	var (
		files []dbFile
		query string
	)

	dbx := sqlx.NewDb(db, "mysql")
	defaultPath := viper.GetString("app.store.path")

	query = `SELECT * FROM files 
				WHERE path IS NULL`
	dbx.Select(&files, query)

	for _, f := range files {
		year, month, day := f.FileDate.Date()
		pathToFolder := fmt.Sprintf("%s/%d-%d-%d", defaultPath, day, month, year)
		if _, err := os.Stat(pathToFolder); os.IsNotExist(err) {
			os.Mkdir(pathToFolder, 0777)
		}

		f.Path.String = fmt.Sprintf("%s/%s", pathToFolder, f.FileName)
		f.Path.Valid = true

		newFile, err := os.Create(f.Path.String)
		if err != nil {
			logger.LogError("Failed create file", "server/app", f.FileName, err)
			continue
		}
		defer newFile.Close()
		split := strings.Split(f.FileData.String, ",")
		data, err := base64.StdEncoding.DecodeString(split[1])

		newFile.Write(data)

		f.FileData.String = split[0]

		query = `UPDATE files SET 
					file_data = :file_data,
					path = :path
					WHERE file_id = :file_id`
		_, err = dbx.NamedExec(query, f)
		if err != nil {
			logger.LogError("Failed update file", "server/app", f.FileName, err)
			continue
		}
	}
}
