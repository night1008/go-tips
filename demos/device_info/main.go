package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"git.sofunny.io/data-analysis/device_info/internal/bootstrap"
	"git.sofunny.io/data-analysis/device_info/internal/config"
)

func main() {
	configFile := flag.String("config", "config.yaml", "")

	flag.Parse()

	c := config.NewConfig("DEVICE_INFO_SERVER")
	if *configFile != "" {
		if err := c.ReadConfigFile(*configFile); err != nil {
			panic("read config: " + err.Error())
		}
	}

	strap, err := bootstrap.New(c)
	if err != nil {
		log.Fatal().Err(err).Msg("server init")
	}

	fmt.Println(strap)

	// setupLogging(c.Log)
	// utils.SetupSentry(c.Sentry)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	log.Info().Msg("server starting")
	if err := strap.Run(ctx); err != nil && err != context.Canceled {
		log.Fatal().Err(err).Msg("server crashed")
	}
	log.Info().Msg("server terminated")
}

// func setupLogging(c config.LogConfig) {
// 	var logger zerolog.Logger
// 	switch c.Format {
// 	case config.LogFormatConsole:
// 		logger = log.Output(zerolog.ConsoleWriter{
// 			Out:        os.Stderr,
// 			TimeFormat: time.RFC3339,
// 			NoColor:    !isatty.IsTerminal(os.Stderr.Fd()),
// 		})
// 	case config.LogFormatJson:
// 		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
// 	}

// 	log.Logger = logger.With().Logger().Level(c.Level)
// }

// func main() {
// 	resp, err := http.Get("https://browser.geekbench.com/v5/cpu/21020281")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := os.WriteFile("a.html", body, 0644); err != nil {
// 		panic(err)
// 	}
// }
