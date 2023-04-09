package bootstrap

import (
	"context"

	"git.sofunny.io/data-analysis/device_info/internal/api"
	"git.sofunny.io/data-analysis/device_info/internal/api/router"
	"git.sofunny.io/data-analysis/device_info/internal/config"
	"git.sofunny.io/data-analysis/device_info/internal/database"
	"git.sofunny.io/data-analysis/device_info/internal/job"
	"git.sofunny.io/data-analysis/device_info/internal/model"
	"github.com/gin-gonic/gin"
)

type Bootstrap struct {
	cfg       *config.Config
	db        *database.DB
	jd        *job.JobService
	apiServer *api.APIServer
}

func New(cfg *config.Config) (*Bootstrap, error) {
	strap := &Bootstrap{
		cfg: cfg,
	}

	db, err := database.New(&cfg.Database)
	if err != nil {
		return nil, err
	}
	strap.db = db

	if err := model.Migrate(strap.db); err != nil {
		return nil, err
	}

	strap.jd = job.New(&cfg.Job, strap.db)

	engine := gin.Default()
	router.RegisterRouter(engine, strap.db)
	strap.apiServer = api.New(cfg, engine, strap.db, strap.jd)

	return strap, nil
}

func (b *Bootstrap) Run(ctx context.Context) error {
	go func() {
		b.jd.Run(ctx)
	}()

	go func() {
		b.apiServer.Run(ctx)
	}()

	select {
	case <-ctx.Done():
		return nil
	}
}
