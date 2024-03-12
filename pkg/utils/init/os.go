package init

import (
	"os"

	"github.com/rs/zerolog/log"
)

func MustLoadEnv(key string) string {
	res := os.Getenv(key)
	if res == "" {
		log.Fatal().Any("env_key", key).Msg("cannot get env variable")
	}

	return res
}
