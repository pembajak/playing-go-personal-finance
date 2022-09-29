package cmd

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pembajak/personal-finance/config"
	"github.com/pembajak/personal-finance/internal/app"
	"github.com/pembajak/personal-finance/internal/app/appcontext"
	"github.com/pembajak/personal-finance/internal/app/delivery/rest"
	tokenJWT "github.com/pembajak/personal-finance/internal/pkg/token"
)

func RunRest() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	cfg := config.NewConfig()
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	appContext := appcontext.NewAppContext()

	db, err := appContext.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	token := tokenJWT.New(tokenJWT.WithIssuer(cfg.GetString("app.issuer")), tokenJWT.WithSecretKey(cfg.GetString("app.secret_key")))

	repo := app.WiringRepository(db)
	usecase := app.WiringUsecase(repo, token)

	restServer := rest.NewServer(
		net.JoinHostPort(cfg.GetString("server.host"), cfg.GetString("server.port")),
		rest.NewRestDelivery(usecase),
		time.Duration(cfg.GetInt("server.read_timeout"))*time.Second,
		time.Duration(cfg.GetInt("server.write_timeout"))*time.Second,
		time.Duration(cfg.GetInt("server.idle_timeout"))*time.Second,
	)

	httpServer := restServer.GetHTTPServer()
	go restServer.GracefullShutdown(httpServer, logger, quit, done)

	logger.Println("=> http server started on", net.JoinHostPort(cfg.GetString("server.host"), cfg.GetString("server.port")))
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", cfg.GetString("server.port"), err)
	}

	<-done

	logger.Println("Server stopped")

}
