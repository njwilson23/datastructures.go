package linkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	lst := New()
	if lst.Length() != 0 {
		t.Fail()
	}

	if lst == nil {
		t.Fail()
	}

	if lst.head != nil {
		t.Fail()
	}

	_, err := lst.Get(0)
	if err == nil {
		t.Fail()
	}
	err = lst.Set(0, 0)
	if err == nil {
		t.Fail()
	}
}

func TestAppend(t *testing.T) {
	lst := New()
	lst.Append(42)
	if lst.head.value.(int) != 42 {
		t.Fail()
	}

	lst.Append(63)
	if lst.head.next.value.(int) != 63 {
		t.Fail()
	}

	lst.Append(100)
	if lst.Length() != 3 {
		t.Fail()
	}
}

func TestPrepend(t *testing.T) {
	lst := New()
	lst.Prepend(42)
	if lst.head.value.(int) != 42 {
		t.Fail()
	}

	lst.Prepend(63)
	if lst.head.value.(int) != 63 {
		t.Fail()
	}

	if lst.Length() != 2 {
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	lst := New()
	lst.Prepend(42)
	lst.Append(63)

	var label interface{}
	label, err := lst.Get(0)
	if err != nil {
		t.Error()
	}
	if label.(int) != 42 {
		t.Fail()
	}

	label, err = lst.Get(1)
	if err != nil {
		t.Error()
	}
	if label.(int) != 63 {
		t.Fail()
	}

	_, err = lst.Get(2)
	if err == nil {
		t.Fail()
	}

	_, err = lst.Get(-1)
	if err == nil {
		t.Fail()
	}
}

func TestSet(t *testing.T) {
	lst := New()
	lst.Prepend(42)
	lst.Append(63)

	lst.Set(0, -42)
	lst.Set(1, 17)
	if lst.head.value.(int) != -42 {
		t.Fail()
	}

	if lst.head.next.value.(int) != 17 {
		t.Fail()
	}

	err := lst.Set(2, 0)
	if err == nil {
		t.Fail()
	}

	err = lst.Set(-1, 0)
	if err == nil {
		t.Fail()
	}
}

func TestInsert(t *testing.T) {
	lst := New()
	err := lst.Insert(0, 0)
	if err == nil {
		t.Fail()
	}

	lst.Append(41)
	lst.Append(42)
	lst.Append(44)
	err = lst.Insert(2, 43)
	if err != nil {
		t.Error()
	}
	if lst.head.next.next.value.(int) != 43 {
		t.Fail()
	}
	if lst.Length() != 4 {
		t.Fail()
	}

	err = lst.Insert(-1, 0)
	if err == nil {
		t.Fail()
	}
	err = lst.Insert(4, 0)
	if err == nil {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	lst := New()

	_, err := lst.Delete(0)
	if err == nil {
		t.Fail()
	}

	lst.Prepend(42)
	lst.Append(63)
	lst.Append(100)

	label, err := lst.Delete(1)
	if err != nil {
		t.Error()
	}
	if label.(int) != 63 {
		t.Fail()
	}
	if lst.Length() != 2 {
		t.Fail()
	}

	_, err = lst.Delete(-1)
	if err == nil {
		t.Fail()
	}

	_, err = lst.Delete(3)
	if err == nil {
		t.Fail()
	}
}
