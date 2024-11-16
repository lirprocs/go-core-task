package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	m := NewStringIntMap()
	key := "key1"
	m.Add(key, 10)

	if m.val[key] != 10 {
		t.Errorf("testAdd() = %d; expected %d", m.val[key], 10)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	key := "key1"
	m.Add(key, 10)
	m.Remove(key)

	if _, exists := m.val[key]; exists {
		t.Errorf("testRemove(): Key %s dont removed", key)
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 10)
	m.Add("key2", 20)

	copyMap := m.Copy()

	if !reflect.DeepEqual(copyMap, m.val) {
		t.Errorf("testCopy() = %v, expected %v", copyMap, m.val)
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	key := "key1"
	m.Add(key, 10)

	if !m.Exists(key) {
		t.Errorf("exists() = %v, expected %t", m.Exists(key), true)
	}

	key2 := "key2"
	if m.Exists(key2) {
		t.Errorf("exists() = %v, expected %t", m.Exists(key2), false)
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 10)

	value, exists := m.Get("key1")
	if !exists || value != 10 {
		t.Errorf("testGet() =  %d, expected %d", value, 10)
	}

	_, exists = m.Get("key2")
	if exists {
		t.Error("Get method returned true for a non-existent key 'key2'")
	}
}
