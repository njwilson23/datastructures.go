package hashtable

import (
	"errors"
	"math"

	"github.com/njwilson23/datastructures/linkedlist"
)

var KEY_NOT_FOUND_ERROR = errors.New("key not found")

type HashTable struct {
	Size     int
	array    []linkedlist.LinkedList
	hashFunc func(string) int
}

type KeyValuePair struct {
	key   string
	value interface{}
}

func sumRune(summable []rune) int {
	sum := 0
	for i := range summable {
		sum = sum + int(summable[i])
	}
	return sum
}

func stringToInt(s string) int {
	r := []rune(s)
	return sumRune(r)
}

func divisionHash(val, size int) int {
	return val - val/size
}

func multiplicationHash(val, size int) int {
	c := 0.5*math.Sqrt(5) - 0.5 // suggested by Knuth
	return int(math.Floor(float64(size) * float64(val) * math.Mod(c, 1.0)))
}

func InitHashTable(size int) *HashTable {
	array := make([]linkedlist.LinkedList, size)
	ht := HashTable{size, array, stringToInt}
	return &ht
}

func (ht *HashTable) Insert(key string, value interface{}) error {
	hashInt := stringToInt(key)
	arrayPos := divisionHash(hashInt, ht.Size)
	ht.array[arrayPos].Append(KeyValuePair{key, value})
	return nil
}

func (ht *HashTable) Get(key string) (interface{}, error) {
	var kv KeyValuePair
	hashInt := stringToInt(key)
	arrayPos := divisionHash(hashInt, ht.Size)
	list := ht.array[arrayPos]
	node := list.Head
	for node != nil {
		kv = node.Value.(KeyValuePair)
		if kv.key == key {
			return kv.value, nil
		}
		node = node.Next
	}
	return nil, KEY_NOT_FOUND_ERROR
}

func (ht *HashTable) Delete(key string) error {
	var kv KeyValuePair
	hashInt := stringToInt(key)
	arrayPos := divisionHash(hashInt, ht.Size)

	list := ht.array[arrayPos]
	node := list.Head
	index := 0
	for node != nil {
		kv = node.Value.(KeyValuePair)
		if kv.key == key {
			list.Delete(index)
			ht.array[arrayPos] = list
			return nil
		}
		index++
		node = node.Next
	}
	return KEY_NOT_FOUND_ERROR
}
