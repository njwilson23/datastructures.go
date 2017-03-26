package hashtable

import (
	"errors"
	"math"

	"github.com/njwilson23/datastructures/linkedlist"
)

var KEY_ERROR = errors.New("key not found")

type Hashable interface {
	Hash() int
}

type HashTable struct {
	Size     int
	array    []*linkedlist.LinkedList
	hashFunc func(int) int
}

type KeyValuePair struct {
	key   Hashable
	value interface{}
}

func sumRune(summable []rune) int {
	sum := 0
	for i := range summable {
		sum = sum + int(summable[i])
	}
	return sum
}

type HashString string

func (hs HashString) Hash() int { return sumRune([]rune(string(hs))) }

func divisionHash(val, size int) int { return val - val/size }

func multiplicationHash(val, size int, c float64) int {
	return int(math.Floor(float64(size) * math.Mod(float64(val)*c, 1.0)))
}

func InitHashTable(size int) *HashTable {
	array := make([]*linkedlist.LinkedList, size)
	for i := range array {
		array[i] = linkedlist.New()
	}
	c := 0.5*math.Sqrt(5) - 0.5 // suggested by Knuth
	ht := HashTable{size, array, func(v int) int { return multiplicationHash(v, size, c) }}
	return &ht
}

func (ht *HashTable) Insert(key Hashable, value interface{}) error {
	arrayPos := ht.hashFunc(key.Hash())
	lst := ht.array[arrayPos]
	lst.Append(KeyValuePair{key, value})
	return nil
}

func (ht *HashTable) Get(key Hashable) (interface{}, error) {
	var kv KeyValuePair
	arrayPos := ht.hashFunc(key.Hash())
	lst := ht.array[arrayPos]
	node := lst.Head
	for node != nil {
		kv = node.Value.(KeyValuePair)
		if kv.key == key {
			return kv.value, nil
		}
		node = node.Next
	}
	return nil, KEY_ERROR
}

func (ht *HashTable) Delete(key Hashable) error {
	var kv KeyValuePair
	arrayPos := ht.hashFunc(key.Hash())

	lst := ht.array[arrayPos]
	node := lst.Head
	index := 0
	for node != nil {
		kv = node.Value.(KeyValuePair)
		if kv.key == key {
			lst.Delete(index)
			return nil
		}
		index++
		node = node.Next
	}
	return KEY_ERROR
}
