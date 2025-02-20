package cmd

import (
	"errors"
	"market/src/api/transport"
	"market/src/repository/postgres"
	"market/src/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logger.NewLogger()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pg, err := postgres.New(cfg.PostgresConfig, log)
	if err != nil {
		log.Panicf("Failed to connect to database: %v", err)
	}

	app := service.NewMarketService(pg, log)
	r := transport.NewRouter(log, app)
	go func() {
		if err = r.Run(ctx, addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("Error starting server: %v", err)
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	<-sigCh
	cancel()
	pg.Close()
	log.Info("Shutting down")
}
