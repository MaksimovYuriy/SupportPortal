package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/config"
	"github.com/MaksimovYuriy/SupportPortal/internal/database/postgres"
	"github.com/MaksimovYuriy/SupportPortal/internal/engine"
	agentservice "github.com/MaksimovYuriy/SupportPortal/internal/services/agent"
	flowservice "github.com/MaksimovYuriy/SupportPortal/internal/services/flow"
	queueservice "github.com/MaksimovYuriy/SupportPortal/internal/services/queue"
	ticketservice "github.com/MaksimovYuriy/SupportPortal/internal/services/ticket"
	userservice "github.com/MaksimovYuriy/SupportPortal/internal/services/user"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/handlers"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/rest"
)

func Run() error {
	appCtx, stopApp := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopApp()
	logger := slog.Default()

	cfg, err := config.Load()
	if err != nil {
		return err
	}
	db, err := postgres.NewDB(cfg.DB)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("failed to close database connection", "error", err)
		}
	}()

	userRepository := postgres.NewUserRepository(db)
	agentRepository := postgres.NewAgentRepository(db)
	queueRepository := postgres.NewQueueRepository(db)
	agentQueueRepository := postgres.NewAgentQueueRepository(db)
	flowRepository := postgres.NewFlowRepository(db)
	flowStepRepository := postgres.NewFlowStepRepository(db)
	ticketRepository := postgres.NewTicketRepository(db)

	agentService := agentservice.NewAgentService(agentRepository, queueRepository, agentQueueRepository)
	userService := userservice.NewUserService(userRepository)
	queueService := queueservice.NewQueueService(queueRepository)
	flowService := flowservice.NewFlowService(flowRepository, flowStepRepository, queueRepository)
	ticketService := ticketservice.NewTicketService(ticketRepository, flowRepository, flowStepRepository, agentRepository, agentQueueRepository)

	userHandler := handlers.NewUserHandler(userService)
	agentHandler := handlers.NewAgentHandler(agentService)
	queueHandler := handlers.NewQueueHandler(queueService)
	flowHandler := handlers.NewFlowHandler(flowService)
	ticketHandler := handlers.NewTicketHandler(ticketService)

	handlers := &rest.Handlers{
		UserHandler:   userHandler,
		AgentHandler:  agentHandler,
		QueueHandler:  queueHandler,
		FlowHandler:   flowHandler,
		TicketHandler: ticketHandler,
	}

	router := rest.NewRouterWithLogger(handlers, logger)
	if cfg.Engine.Enabled {
		ticketEngine := engine.NewEngine(ticketService, agentService, logger, cfg.Engine.Interval, cfg.Engine.BatchLimit)
		go ticketEngine.Run(appCtx)
	}

	addr := ":8080"
	server := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       time.Minute,
	}

	logger.Info("SupportPortal API started", "addr", addr)
	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-appCtx.Done():
		logger.Info("SupportPortal API shutting down")
	}

	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	return nil
}
