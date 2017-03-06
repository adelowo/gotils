package bag

type Bag interface {
	Add(key, value string)
	Delete(key string)
	Has(key string) bool
	Get(key string) (string, error)
	Reset()
	Count() int
}
