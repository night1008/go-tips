package job

import (
	"context"
	"fmt"
	"time"

	"git.sofunny.io/data-analysis/device_info/internal/config"
	"git.sofunny.io/data-analysis/device_info/internal/database"
	"github.com/rs/zerolog/log"
	"github.com/simpleframeworks/jobsd"
)

type JobService struct {
	*jobsd.JobsD
}

func New(cfg *config.JobCfg, db *database.DB) *JobService {
	jd := jobsd.New(db.DB).
		WorkerNum(cfg.WorkerNum).
		PollInterval(time.Duration(cfg.PollInterval) * time.Second).
		PollLimit(cfg.PollLimit).
		RetriesTimeoutLimit(cfg.RetriesOnTimeoutLimit).
		RetriesErrorLimit(cfg.RetriesOnErrorLimit)

	return &JobService{
		JobsD: jd,
	}
}

func (s *JobService) Run(ctx context.Context) error {
	log.Info().Str("m", "job").Msg("job server start")
	if err := s.Up(); err != nil {
		return fmt.Errorf("failed to start jobsd: %s", err)
	}

	<-ctx.Done()
	if err := s.Down(); err != nil {
		return fmt.Errorf("failed to down jobsd: %s", err)
	}
	return nil
}
