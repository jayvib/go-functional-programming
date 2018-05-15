package collection

const (
	INVALID_INT_VAL = -1
	INVALID_STRING_VAL = "<nil>"
)

type CarIterator interface {
	Next() (value string, ok bool)
}

type Collection struct {
	index int
	List []string
}

func NewCollection(s []string) *Collection {
	return &Collection{INVALID_INT_VAL, s}
}

func (collection *Collection) Next() (value string, ok bool) {
	collection.index++
	if collection.index >= len(collection.List) {
		return INVALID_STRING_VAL, false
	}
	return collection.List[collection.index], true
}

func (collection *Collection) Map(f func(string) string) *Collection {
	for i := 0; i < collection.len(); i++ {
		collection.List[i] = f(collection.List[i])
	}
	return collection
}

func (collection *Collection) len() int {
	return len(collection.List)
}