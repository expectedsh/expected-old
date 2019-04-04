package main

import (
	"database/sql"
	"github.com/expectedsh/expected/pkg/accounts"
	"github.com/expectedsh/expected/pkg/containers"
	"github.com/expectedsh/expected/pkg/images"
	"github.com/expectedsh/expected/pkg/registryserver"
	"github.com/expectedsh/expected/pkg/registryserver/auth/token"
	"github.com/expectedsh/expected/pkg/services"
	"github.com/expectedsh/expected/pkg/services/postgres"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Addr  string `envconfig:"addr" default:":3000"`
	Certs struct {
		PublicKey  string `envconfig:"public_key" default:"./certs/server.crt"`
		PrivateKey string `envconfig:"private_key" default:"./certs/server.key"`
	}
}

func initDB(addr string, connMaxLifetime time.Duration, maxIdleConns, maxOpenConns int) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(connMaxLifetime)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	if err = accounts.InitDB(db); err != nil {
		return nil, err
	}
	if err = containers.InitDB(db); err != nil {
		return nil, err
	}
	if err = images.InitDB(db); err != nil {
		return nil, err
	}
	return db, err
}

func main() {
	logrus.Infoln("processing environment configuration")
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		logrus.WithError(err).Fatalln("unable to parse environment variables")
	}

	logrus.SetFormatter(&logrus.TextFormatter{})

	services.Register(postgres.NewFromEnv())
	services.Start()
	defer services.Stop()

	if err := accounts.InitDB(services.Postgres().Client()); err != nil {
		logrus.Fatal(err)
	}
	if err := images.InitDB(services.Postgres().Client()); err != nil {
		logrus.Fatal(err)
	}

	token.Init(config.Certs.PublicKey, config.Certs.PrivateKey)

	logrus.Infoln("starting api server")
	server := registryserver.New(config.Addr)

	logrus.Infof("listening on %v\n", config.Addr)
	if err := server.Start(); err != nil {
		logrus.WithError(err).Fatalln("unable to start api server")
	}
}
