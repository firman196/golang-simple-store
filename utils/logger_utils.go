package utils

import (
	"github.com/rs/zerolog/log"
)

func FatalError(err string, msg string) {
	log.Fatal().Err(err).Msg(msg)
}
