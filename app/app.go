package app

import (
	"github.com/dingowd/RB/internal/cache"
	"github.com/dingowd/RB/internal/logger"
	"github.com/dingowd/RB/internal/storage"
)

type App struct {
	Log   logger.Logger
	Store storage.Storage
	Cache cache.CacheInterface
}

func New(logger logger.Logger, storage storage.Storage, cache cache.CacheInterface) *App {
	return &App{
		Log:   logger,
		Store: storage,
		Cache: cache,
	}
}
