package icache

type Cache interface {
	Get() (interface{}, error)
	Set() error
}
