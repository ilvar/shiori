package webserver

import (
	"time"

	"github.com/go-shiori/shiori/internal/config"
	"github.com/go-shiori/shiori/internal/database"
	cch "github.com/patrickmn/go-cache"
)

// Config is parameter that used for starting web server
type Config struct {
	DB            database.DB
	DataDir       string
	OpenAiKey     string
	OpenAiModel   string
	ServerAddress string
	ServerPort    int
	RootPath      string
	Log           bool
}

func GetLegacyHandler(cfg Config, dependencies *config.Dependencies) *Handler {
	return &Handler{
		DB:           cfg.DB,
		DataDir:      cfg.DataDir,
		OpenAiKey:    cfg.OpenAiKey,
		OpenAiModel:  cfg.OpenAiModel,
		UserCache:    cch.New(time.Hour, 1000*time.Hour),
		SessionCache: cch.New(time.Hour, 1000*time.Hour),
		ArchiveCache: cch.New(time.Minute, 5*time.Hour),
		RootPath:     cfg.RootPath,
		Log:          cfg.Log,
		depenencies:  dependencies,
	}
}
