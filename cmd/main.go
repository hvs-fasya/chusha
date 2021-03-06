package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/jessevdk/go-flags"
	"github.com/kabukky/httpscerts"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/hvs-fasya/chusha/internal/api"
	"github.com/hvs-fasya/chusha/internal/api/handlers/front"
	"github.com/hvs-fasya/chusha/internal/engine"
	"github.com/hvs-fasya/chusha/internal/migrate"
	"github.com/hvs-fasya/chusha/internal/redis-client"
)

var opts struct {
	DatabaseName    string `long:"db-name" env:"DB_NAME" default:"chusha" description:"database name"`
	DatabaseUser    string `long:"db-user" env:"DB_USER" default:"chusha" description:"user name of database"`
	DatabasePass    string `long:"db-pass" env:"DB_PASS" default:"chusha" description:"user password of database"`
	DatabaseDialect string `long:"db-dialect" env:"DB_DIALECT" default:"postgres" description:"database engine"`
	DatabaseHOST    string `long:"db-host" env:"DB_HOST" default:"127.0.0.1:5432" description:"database host name"`

	RedisHOST string `long:"redis-host" env:"REDIS_HOST" default:"localhost:6379" description:"redis-client host name"`
	RedisPass string `long:"redis-pass" env:"REDIS_PASS" default:"" description:"redis-client password"`
	RedisDB   int    `long:"redis-db" env:"REDIS_DB" default:"0" description:"redis-client database"`

	APIPort string `long:"api-port" env:"API_PORT" default:"8080" description:"api server port"`
	//FrontPort  string `long:"front-port" env:"FRONT_PORT" default:"8081" description:"front server port"`
	StaticPath string `long:"static-path" env:"STATIC_PATH" default:"./front/stage/" description:"static files path"`

	AppEnv      string `long:"environment" env:"APP_ENV" default:"dev"  description:"app environment"`
	LogLevel    int    `long:"log-level" env:"LOG_LEVEL" default:"0" description:"log level debug:0, info: 1, warn: 2, error: 3, fatal: 4, panic:5"`
	MigrateDown bool   `long:"migrate-down" env:"MIGRATE_DOWN" description:"migrate down and exit"`
	Timezone    string `long:"timezone" env:"TIMEZONE" default:"Europe/Moscow"  description:"app timezone"`
}

func main() {
	p := flags.NewParser(&opts, flags.Default)
	if _, e := p.ParseArgs(os.Args[1:]); e != nil {
		os.Exit(0)
	}
	setLogger(opts.LogLevel)
	_ = os.Setenv("TZ", opts.Timezone)
	//подключение к базе даннх
	dsn := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
		opts.DatabaseDialect, opts.DatabaseUser, opts.DatabasePass, opts.DatabaseHOST, opts.DatabaseName)
	dataBase, err := engine.NewPgDB(dsn)
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	engine.DB = dataBase
	log.Info().Msgf("Connect to database - success. DB host: %s", opts.DatabaseHOST)
	//подключение к редис
	redisOptions := &redis.Options{
		Addr:     opts.RedisHOST,
		Password: opts.RedisPass,
		DB:       opts.RedisDB,
	}
	err = redis_client.NewRedis(redisOptions)
	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(0)
	}

	//накатываем миграции
	migrationService := migrate.NewMigrationService(dataBase.Conn, opts.DatabaseDialect)
	if opts.MigrateDown {
		n, err := migrationService.MigrateDown()
		if err != nil {
			log.Panic().Msg(err.Error())
		}
		log.Info().Msgf("DB migrations rollback - %d migrations", n)
		os.Exit(0)
	}
	n, err := migrationService.MigrateUP()
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	log.Info().Msgf("DB migrate - %d migrations", n)
	//инициализация сервера
	err = httpscerts.Check("cert.pem", "key.pem")
	if err != nil {
		fmt.Println(err)
		if opts.AppEnv == "dev" {
			err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:8080")
			if err != nil {
				log.Fatal().Msgf("Self-signed certs generate error: %s", err)
			}
		} else {
			log.Fatal().Msg("No https certs")
		}
	}
	srv := api.Server{}
	connstr := ":" + opts.APIPort
	front.InitFront(opts.StaticPath)
	srv.Run(connstr)
}

func setLogger(level int) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	switch level {
	case 0:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case 1:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case 2:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case 3:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case 4:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case 5:
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
