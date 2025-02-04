package main

import (
	"context"
	"log"

	"github.com/kingxl111/Pinger/backend/internal/handlers"
	"github.com/kingxl111/Pinger/backend/internal/service"
	"github.com/kingxl111/Pinger/backend/internal/storage"

	conf "github.com/kingxl111/Pinger/backend/internal/config"
	logg "github.com/kingxl111/Pinger/backend/internal/logging"
)

func main() {
	ctx := context.Background()

	cfg := conf.MustLoad()

	logger, err := logg.NewLogger("logs.txt")
	if err != nil {
		log.Fatalf("can't initialize logger: %v", err.Error())
	}

	//// wait 7 s before connect to db
	//time.Sleep(7 * time.Second)
	db, err := storage.NewDB(
		ctx,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName,
		cfg.DB.SSLmode)
	if err != nil {
		log.Fatalf("can't initialize database: %v", err.Error())
	}
	defer db.Close()

	st := storage.NewStorage(db)
	services := service.NewService(st)
	router := handlers.NewHandler(services)

	srv := &handlers.Server{}
	log.Printf("server started on %s", cfg.HTTPServer.Address)
	err = srv.Run(router.NewRouter(&ctx, logger.Lg, cfg.Env), cfg)
	if err != nil {
		log.Fatalf("can't start server: %v", err.Error())
	}
}
