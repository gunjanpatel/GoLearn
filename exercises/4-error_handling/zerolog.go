package main

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Debug().
		Int("EmployeeId", 100).
		Msg("Getting info")

	time.Sleep(time.Second * 2)

	log.Debug().
		Str("Name", "Gunjan").
		Send()

	log.Error().Msg("Got the error")
}
