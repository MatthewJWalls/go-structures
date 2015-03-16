package main

import "fmt"

const (
	INITIAL_SIZE = 1024
)

type KeyVal struct {
	key string
	val interface{}
}

type HashTable struct {
	buckets []*SingleList
}

func (this *HashTable) Put(key string, value interface{}) {
	bucket := this.buckets[elfHash(key, INITIAL_SIZE)]
	bucket.Push(&KeyVal{key, value})
}

func (this *HashTable) Get(key string) (interface{}) {

	bucket := this.buckets[elfHash(key, INITIAL_SIZE)]

	for i := 0; i < bucket.Length(); i++ {
		keyval := bucket.Get(i).(*KeyVal)
		if keyval.key == key {
			return keyval.val
		}
	}

	return nil

}

func (this *HashTable) Stats() string {

	out := "Hash Table Buckets:"

	for i := 0; i < INITIAL_SIZE; i++ {
		if this.buckets[i].Length() > 0 {
			out = fmt.Sprintf("%s %d(%d)", out, i, this.buckets[i].Length())
		}
	}

	return out

}

func (this *HashTable) Length() int {

	length := 0

	for i := 0; i < INITIAL_SIZE; i++ {
		if this.buckets[i].Length() > 0 {
			length += 1
		}
	}

	return length

}

func NewHashTable() *HashTable {
	table := HashTable{make([]*SingleList, INITIAL_SIZE, INITIAL_SIZE)}

	for i := 0; i < INITIAL_SIZE; i++ {
		table.buckets[i] = NewLinkedList()
	}

	return &table
}


// this is a hash function that was (is?) used in unix.

func elfHash(toHash string, max int) int {

	hashValue := 0

	for i := 0; i < len(toHash); i++ {
		
		hashValue = (hashValue << 4) + int(toHash[i])
		hiBits := hashValue & 0xF0000000

		if hiBits != 0 {
			hashValue ^= hiBits >> 24
		}

		hashValue &= ^hiBits

	}

	return hashValue % max

}

