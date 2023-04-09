package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"git.sofunny.io/data-analysis/device_info/internal/config"
	"git.sofunny.io/data-analysis/device_info/internal/database"
	"git.sofunny.io/data-analysis/device_info/internal/job"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type APIServer struct {
	cfg    *config.Config
	db     *database.DB
	jd     *job.JobService
	engine *gin.Engine
	server *http.Server
}

func New(cfg *config.Config, engine *gin.Engine, db *database.DB, jd *job.JobService) *APIServer {
	mux := http.NewServeMux()
	mux.Handle("/", engine)

	server := &http.Server{
		Addr:         cfg.APIServer.Listen,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	return &APIServer{
		cfg:    cfg,
		db:     db,
		jd:     jd,
		server: server,
		engine: engine,
	}
}

func (s *APIServer) Run(ctx context.Context) error {
	log.Info().Str("m", "api-server").Str("addr", s.cfg.APIServer.Listen).Msg("listen on")

	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
