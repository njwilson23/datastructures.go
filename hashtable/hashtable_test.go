package hashtable

import (
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
	s := "bees and oats"
	sum := stringToInt(s)
	if sum != 1225 {
		t.Fail()
	}
}

func TestHashTable(t *testing.T) {
	var err error
	ht := InitHashTable(5000)

	err = ht.Insert("colour", "#4682b4")
	if err != nil {
		t.Error()
	}

	err = ht.Insert("age", "unknown")
	if err != nil {
		t.Error()
	}

	err = ht.Insert("size", "large")
	if err != nil {
		t.Error()
	}

	value, err := ht.Get("colour")
	if err != nil {
		t.Error()
	}
	if value.(string) != "#4682b4" {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	var err error
	ht := InitHashTable(5000)

	err = ht.Insert("colour", "#4682b4")
	if err != nil {
		t.Error()
	}

	err = ht.Insert("age", "unknown")
	if err != nil {
		t.Error()
	}

	ht.Delete("colour")
	_, err = ht.Get("colour")
	if err != KEY_NOT_FOUND_ERROR {
		t.Error()
	}
}
