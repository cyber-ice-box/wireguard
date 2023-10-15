package app

import (
	postgresConfig "github.com/cyber-ice-box/agent/pkg/postgres"
	"github.com/cyber-ice-box/wireguard/internal/config"
	"github.com/cyber-ice-box/wireguard/internal/delivery/controller/grpc"
	"github.com/cyber-ice-box/wireguard/internal/delivery/repository/postgres"
	"github.com/cyber-ice-box/wireguard/internal/service"
	"github.com/rs/zerolog/log"
	"os/exec"

	"net"
	"os"
	"os/signal"
	"syscall"
)

// Run initializes whole application.
func Run(configPath *string) {
	cfg, valid := config.GetConfig(configPath)
	if !valid {
		os.Exit(1)
	}

	db, err := postgresConfig.NewPostgresDB(postgresConfig.Config(cfg.PostgresDB))
	if err != nil {
		log.Fatal().Err(err).Msg("Can not connect to postgres DB")
	}

	wgService := service.NewService()

	if err = wgService.InitServer(cfg.VPN.Address, cfg.VPN.Port, cfg.VPN.PrivateKey); err != nil {
		log.Fatal().Err(err).Msg("failed to configure wireguard server")
	}

	if err = wgService.InitServerClients(postgres.NewPostgresRepository(db)); err != nil {
		log.Fatal().Err(err).Msg("failed to configure wireguard server clients")
	}

	lis, err := net.Listen("tcp", cfg.GRPC.Endpoint)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	grpcServer, err := grpc.New(&cfg.GRPC, wgService)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to setup grpc server")
	}

	if err = exec.Command("/bin/sh", "-c", "sudo sysctl -p").Run(); err != nil {
		log.Fatal().Err(err).Msg("failed to set forward rules")
	}

	if _, err = os.Create("/ready"); err != nil {
		log.Fatal().Err(err).Msg("failed to create ready file")
	}

	go func() {

		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("failed to serve")
		}
	}()

	log.Printf("wireguard gRPC server is running at %s...\n", cfg.GRPC.Endpoint)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	grpcServer.GracefulStop()
}
