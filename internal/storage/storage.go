package storage

import (
	"context"
	"errors"
	"github.com/dingowd/RB/model"
)

type Storage interface {
	Connect(ctx context.Context, dsn string) error
	Close()
	GetAll() (*model.Students, error)
	Update(m model.CacheStudent) error
	Delete(id string) error
	Insert(s model.ForJson) error
}

var (
	ErrorDocumentExist = errors.New("Document already exist")
)
