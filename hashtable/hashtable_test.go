package hashtable

import (
	"math"
	"testing"
)

func TestSumRune(t *testing.T) {
	intArray := []int{71, 32, 89, 123, 32, 14}
	runeArray := []rune{}
	for _, i := range intArray {
		runeArray = append(runeArray, rune(i))
	}
	sum := sumRune(runeArray)
	if sum != 361 {
		t.Fail()
	}
}

func TestStringToInt(t *testing.T) {
	s := HashString("bees and oats")
	sum := s.Hash()
	if sum != 1225 {
		t.Fail()
	}
}

func TestHashTable(t *testing.T) {
	var err error
	ht := InitHashTable(int(math.Pow(2, 14)))

	err = ht.Insert(HashString("colour"), "#4682b4")
	if err != nil {
		t.Error()
	}

	err = ht.Insert(HashString("age"), "unknown")
	if err != nil {
		t.Error()
	}

	err = ht.Insert(HashString("size"), "large")
	if err != nil {
		t.Error()
	}

	value, err := ht.Get(HashString("colour"))
	if err != nil {
		t.Error()
	}
	if value.(string) != "#4682b4" {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	var err error
	ht := InitHashTable(int(math.Pow(2, 14)))

	err = ht.Insert(HashString("colour"), "#4682b4")
	if err != nil {
		t.Error()
	}

	err = ht.Insert(HashString("age"), "unknown")
	if err != nil {
		t.Error()
	}

	ht.Delete(HashString("colour"))
	_, err = ht.Get(HashString("colour"))
	if err != KEY_ERROR {
		t.Error()
	}
}
