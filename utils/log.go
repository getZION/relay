package utils

import (
	"os"

	"github.com/rs/zerolog"
)

var Log = zerolog.New(os.Stderr).Output(zerolog.ConsoleWriter{Out: os.Stdout})
