package cache

/*type CacheInterface interface {
	Init()
	ReadFromCache() error
	WriteToCache(o model.Order) int
}

type Cache struct {
	Log  logger.Logger
	Body model.CacheOrderList
	Tick int
	Store storage.Storage
}

func NewCache(log logger.Logger, store storage.Storage, t int) *Cache {
	return &Cache{
		Log:  log,
		Body: make(model.CacheOrderList, 0),
		Tick: t,
		Store: store,
	}
}

func (c *Cache) Init() {
	c.Store.GetAll()
}

func (c *Cache) ReadFromCache() error {

	return nil
}

func (c *Cache) WriteToCache(o model.Order) int {

	return 1
}*/
