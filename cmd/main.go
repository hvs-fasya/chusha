package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/hvs-fasya/chusha/internal/api"
	"github.com/hvs-fasya/chusha/internal/api/handlers/front"
	"github.com/hvs-fasya/chusha/internal/engine"
	"github.com/hvs-fasya/chusha/internal/migrate"
)

var opts struct {
	DatabaseName    string `long:"db-name" env:"DB_NAME" default:"chusha" description:"database name"`
	DatabaseUser    string `long:"db-user" env:"DB_USER" default:"chusha" description:"user name of database"`
	DatabasePass    string `long:"db-pass" env:"DB_PASS" default:"chusha" description:"user password of database"`
	DatabaseDialect string `long:"db-dialect" env:"DB_DIALECT" default:"postgres" description:"database engine"`
	DatabaseHOST    string `long:"db-host" env:"DB_HOST" default:"127.0.0.1:5432" description:"database host name"`

	APIPort string `long:"api-port" env:"API_PORT" default:"8080" description:"api server port"`
	//FrontPort  string `long:"front-port" env:"FRONT_PORT" default:"8081" description:"front server port"`
	StaticPath string `long:"static-path" env:"STATIC_PATH" default:"./front/stage/" description:"static files path"`

	LogLevel    int    `long:"log-level" env:"LOG_LEVEL" default:"0" description:"log level debug:0, info: 1, warn: 2, error: 3, fatal: 4, panic:5"`
	MigrateDown bool   `long:"migrate-down" env:"MIGRATE_DOWN" description:"migrate down and exit"`
	Timezone    string `long:"timezone" env:"TIMEZONE" default:"Europe/Moscow"  description:"app timezone"`
}

func main() {

	p := flags.NewParser(&opts, flags.Default)
	if _, e := p.ParseArgs(os.Args[1:]); e != nil {
		os.Exit(1)
	}
	setLogger()
	_ = os.Setenv("TZ", opts.Timezone)
	//подключение к базе даннх
	dsn := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
		opts.DatabaseDialect, opts.DatabaseUser, opts.DatabasePass, opts.DatabaseHOST, opts.DatabaseName)
	dataBase, err := engine.NewPgDB(dsn)
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	engine.DB = dataBase
	log.Info().Msgf("Подключение к БД - успешно. БД хост: %s", opts.DatabaseHOST)
	//накатываем миграции
	migrationService := migrate.NewMigrationService(dataBase.Conn, opts.DatabaseDialect)
	if opts.MigrateDown {
		n, err := migrationService.MigrateDown()
		if err != nil {
			log.Panic().Msg(err.Error())
		}
		log.Info().Msgf("Откат миграций БД - %d миграций успешно", n)
		os.Exit(0)
	}
	n, err := migrationService.MigrateUP()
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	log.Info().Msgf("Миграции БД - %d миграций успешно", n)
	//инициализация сервера
	srv := api.Server{}
	connstr := ":" + opts.APIPort
	front.InitFront(opts.StaticPath)
	srv.Run(connstr)
}

func setLogger() {
	zerolog.SetGlobalLevel(zerolog.Level(opts.LogLevel))
}
