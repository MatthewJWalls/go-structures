package main

import (
	"testing"
)

func TestHashTableBasic(t *testing.T) {

	table := NewHashTable()

	if table.Length() != 0 {
		t.Error("Incorrect length at start.")
	}

	table.Put("hello", 400)
	table.Put("chimp", 999)
	
	if table.Get("hello").(int) != 400 {
		t.Error("hello not found in hash.")
	}

	if table.Get("chimp").(int) != 999 {
		t.Error("chimp not found in hash.")
	}

	if table.Length() != 2 {
		t.Error("Incorrect length.")
	}

}

func TestHashFunction(t *testing.T) {

	if elfHash("derp", 1024) == elfHash("fire", 1024) {
		t.Error("hash(derp) can't equal hash(herp)")
	}

}
