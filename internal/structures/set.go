package structures

type StringSet interface {
	Insert(string)
	Contains(string) bool
}

type HashSet struct {
	values map[string]interface{}
}

func NewHashSet() *HashSet {
	return &HashSet{map[string]interface{}{}}
}

func HashSetFromSlice(inputs []string) *HashSet {
	h := NewHashSet()
	for _, item := range inputs {
		h.Insert(item)
	}
	return h
}

func (h *HashSet) Insert(value string) {
	h.values[value] = nil
}

func (h *HashSet) Contains(value string) bool {
	_, ok := h.values[value]
	return ok
}