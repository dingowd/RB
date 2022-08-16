package cache

import (
	"github.com/dingowd/RB/internal/logger"
	"github.com/dingowd/RB/internal/storage"
	"github.com/dingowd/RB/model"
	"sync"
	"time"
)

type CacheInterface interface {
	ReadFromCache() *model.CacheStudents
	WriteToCache(stop chan struct{})
}

type Cache struct {
	Log   logger.Logger
	Body  model.CacheStudents
	Tick  int
	Store storage.Storage
}

func NewCache(log logger.Logger, store storage.Storage, t int) *Cache {
	return &Cache{
		Log:   log,
		Body:  make(model.CacheStudents, 0),
		Tick:  t,
		Store: store,
	}
}

func (c *Cache) ReadFromCache() *model.CacheStudents {
	cache := make(model.CacheStudents, 0)
	if len(c.Body) > 0 {
		cache = c.Body
	}
	return &cache
}

func (c *Cache) WriteToCache(stop chan struct{}) {
	var mu sync.Mutex
	empty := make(model.CacheStudents, 0)
	for {
		select {
		case <-stop:
			return
		default:
			d := new(model.Students)
			d, _ = c.Store.GetAll()
			mu.Lock()
			c.Body = empty
			for _, v := range *d {
				var e model.CacheStudent
				e.Id = v.Id.Hex()
				e.FirstName = v.FirstName
				e.SecondName = v.SecondName
				e.Faculty = v.Faculty
				tm := time.Unix(v.BirthDate, 0)
				e.BirthDate = tm.Format("02.01.2006")
				c.Body = append(c.Body, e)
			}
			mu.Unlock()
			time.Sleep(time.Duration(c.Tick) * time.Second)
		}
	}
}
