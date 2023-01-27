package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/go-wonk/si"
	"github.com/go-wonk/si/sigorm"
	"github.com/go-wonk/si/sihttp"
	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	"github.com/w-woong/common/configs"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	productadapter "github.com/w-woong/product/adapter"
	productport "github.com/w-woong/product/port"
	"github.com/w-woong/woong/adapter"
	"github.com/w-woong/woong/cmd/route"
	"github.com/w-woong/woong/entity"
	"github.com/w-woong/woong/port"
	"github.com/w-woong/woong/usecase"
	"gorm.io/gorm"
)

var (
	Version = "undefined"

	printVersion     bool
	tickIntervalSec  int = 30
	addr             string
	certPem, certKey string
	readTimeout      int
	writeTimeout     int
	configName       string
	maxProc          int

	usePprof    = false
	pprofAddr   = ":56060"
	autoMigrate = false
)

func init() {
	flag.StringVar(&addr, "addr", ":49001", "listen address")
	flag.BoolVar(&printVersion, "version", false, "print version")
	flag.IntVar(&tickIntervalSec, "tick", 30, "tick interval in second")
	flag.StringVar(&certKey, "key", "./certs/key.pem", "server key")
	flag.StringVar(&certPem, "pem", "./certs/cert.pem", "server pem")
	flag.IntVar(&readTimeout, "readTimeout", 30, "read timeout")
	flag.IntVar(&writeTimeout, "writeTimeout", 30, "write timeout")
	flag.StringVar(&configName, "config", "./configs/server.yml", "config file name")
	flag.IntVar(&maxProc, "mp", runtime.NumCPU(), "GOMAXPROCS")

	flag.BoolVar(&usePprof, "pprof", false, "use pprof")
	flag.StringVar(&pprofAddr, "pprof_addr", ":56060", "pprof listen address")
	flag.BoolVar(&autoMigrate, "autoMigrate", false, "auto migrate")

	flag.Parse()
}

func main() {
	var err error

	if printVersion {
		fmt.Printf("version \"%v\"\n", Version)
		return
	}
	runtime.GOMAXPROCS(maxProc)

	// config
	conf := common.Config{}
	if err := configs.ReadConfigInto(configName, &conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// logger
	logger.Open(conf.Logger.Level, conf.Logger.Stdout,
		conf.Logger.File.Name, conf.Logger.File.MaxSize, conf.Logger.File.MaxBackup,
		conf.Logger.File.MaxAge, conf.Logger.File.Compressed)
	defer logger.Close()

	// db
	sqlDB, err := si.OpenSqlDB(conf.Server.Repo.Driver, conf.Server.Repo.ConnStr,
		conf.Server.Repo.MaxIdleConns, conf.Server.Repo.MaxOpenConns,
		time.Duration(conf.Server.Repo.ConnMaxLifetimeMinutes)*time.Minute)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer sqlDB.Close()

	// gorm
	var gormDB *gorm.DB
	switch conf.Server.Repo.Driver {
	case "pgx":
		gormDB, err = sigorm.OpenPostgres(sqlDB)
	default:
		logger.Error(conf.Server.Repo.Driver + " is not allowed")
		os.Exit(1)
	}
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// repo
	var beginner common.TxBeginner
	var appConfigRepo port.AppConfigRepo
	var homeRepo port.HomeRepo
	var shortNoticeRepo port.ShortNoticeRepo
	var homeGroupProductRepo port.HomeGroupProductRepo
	switch conf.Server.Repo.Driver {
	case "pgx":
		beginner = txcom.NewGormTxBeginner(gormDB)
		appConfigRepo = adapter.NewPgAppconfig(gormDB)
		homeRepo = adapter.NewHomePg(gormDB)
		shortNoticeRepo = adapter.NewShortNoticePg(gormDB)
		homeGroupProductRepo = adapter.NewHomeGroupProductPg(gormDB)
	default:
		logger.Error(conf.Server.Repo.Driver + " is not allowed")
		os.Exit(1)
	}

	if autoMigrate {
		gormDB.AutoMigrate(&entity.AppConfig{}, &entity.Home{}, &entity.ShortNotice{},
			&entity.MainPromotion{}, &entity.Tag{}, &entity.HomeGroupProduct{})
	}

	var groupSvc productport.GroupSvc
	groupSvc = productadapter.NewGroupHttp(sihttp.DefaultInsecureClient(),
		conf.Client.Http.Url, conf.Client.Http.BearerToken)

	// usecase
	usc := usecase.NewAppConfigUsc(beginner, appConfigRepo)
	homeUsc := usecase.NewHomeUsc(beginner, homeRepo, shortNoticeRepo, groupSvc)
	homeGroupProductUsc := usecase.NewHomeGroupProductUsc(beginner, homeGroupProductRepo)

	// router
	router := mux.NewRouter()
	route.AppConfigRoute(router, conf.Server.Http, usc)
	route.HomeRoute(router, conf.Server.Http, homeUsc, homeGroupProductUsc)

	// http server
	tlsConfig := sihttp.CreateTLSConfigMinTls(tls.VersionTLS12)
	httpServer := sihttp.NewServerCors(router, tlsConfig, addr,
		time.Duration(writeTimeout)*time.Second, time.Duration(readTimeout)*time.Second,
		certPem, certKey,
		strings.Split(conf.Server.Http.AllowedOrigins, ","),
		strings.Split(conf.Server.Http.AllowedHeaders, ","),
		strings.Split(conf.Server.Http.AllowedMethods, ","),
	)

	// ticker
	ticker := time.NewTicker(time.Duration(tickIntervalSec) * time.Second)
	tickerDone := make(chan bool)
	common.StartTicker(tickerDone, ticker, func(t time.Time) {
		logger.Info(fmt.Sprintf("NoOfGR:%v, %v", runtime.NumGoroutine(), t))
	})

	// signal, wait for it to shutdown http server.
	common.StartSignalStopper(httpServer, syscall.SIGINT, syscall.SIGTERM)

	// start
	logger.Info("start listening on " + addr)
	if err = httpServer.Start(); err != nil {
		logger.Error(err.Error())
	}

	// finish
	ticker.Stop()
	tickerDone <- true
	logger.Info("finished")
}
