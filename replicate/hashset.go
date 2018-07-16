package replicate

type Set interface {
	Add(value string)
	Contains(value string) (bool)
	Length() (int)
	Remove(value string)
	Keys() ([]string)
}

type StringHashSet struct {
	data map[string]bool
}

func (mapSet *StringHashSet) Add(value string) {
	mapSet.data[value] = true
}

func (mapSet *StringHashSet) Contains(value string) (exists bool) {
	_, exists = mapSet.data[value]
	return
}

func (mapSet *StringHashSet) Length() (int) {
	return len(mapSet.data)
}

func(mapSet *StringHashSet) Remove(value string) {
	delete(mapSet.data, value)
}

func (mapSet *StringHashSet) Keys() ([]string) {
	keys := make([]string, mapSet.Length())

	i := 0
	for k := range mapSet.data {
		keys[i] = k
		i++
	}

	return keys
}


func NewStringSet() (Set) {
	return &StringHashSet{make(map[string] bool)}
}