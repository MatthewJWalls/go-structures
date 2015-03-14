package main

import "testing"

func TestPush(t *testing.T) {

	list := NewLinkedList()

	list.Push("durr")
	list.Push("hurr")

	if list.first.Value.(string) != "hurr" {
		t.Error("First list value was incorrect.")
	}

	if list.first.next.Value.(string) != "durr" {
		t.Error("Second list value was incorrect.")
	}

	if list.first.next.next != nil {
		t.Error("List did not end.")
	}

}

func TestPushBack(t *testing.T) {

	list := NewLinkedList()

	list.PushBack("one")
	list.PushBack("two")
	list.PushBack("three")

	if list.first.Value.(string) != "one" {
		t.Error("First list value was incorrect.")
	}

	if list.first.next.Value.(string) != "two" {
		t.Error("Second list value was incorrect.")
	}

	if list.first.next.next.Value.(string) != "three" {
		t.Error("Third list value was incorrect.")
	}

	if list.first.next.next.next != nil {
		t.Error("List did not end.")
	}

	if list.last.Value != "three" {
		t.Error("List last was incorrect.")
	}

}

func TestLength(t *testing.T) {

	list := NewLinkedList()

	list.Push("durr")
	list.Push("hurr")

	if list.Length() != 2 {
		t.Error("List length did not resolve to 2.")
	}

}

func TestGet(t *testing.T) {

	list := NewLinkedList()

	list.PushBack("one")
	list.PushBack("two")
	list.PushBack("three")

	if list.Get(2).(string) != "three" {
		t.Error("Faild to get.")
	}

	if list.Get(0).(string) != "one" {
		t.Error("Faild to get.")
	}

	if list.Get(1).(string) != "two" {
		t.Error("Faild to get.")
	}

}

func TestDelete(t *testing.T) {

	list := NewLinkedList()

	list.PushBack("one")
	list.PushBack("two")
	list.PushBack("three")

	list.Delete(1)

	// produces {one, three}

	if list.first.Value != "one" {
		t.Error("did not delete properly. First was not one.")
	}

	if list.last.Value != "three" {
		t.Error("Failed to delete from list.")
	}

	if list.first.next.Value != "three" {
		t.Error("did not delete properly. Last was not three.")
	}

	if list.first.next.next != nil {
		t.Error("did not delete properly. Went beyond the veil.")
	}

	list.Delete(1)

	// produces {one}

	if list.last.Value != "one" {
		t.Error("Failed to delete from list. Last was incorrect")
	}

	if list.first.Value != "one" {
		t.Error("did not delete properly. First was not one.")
	}

	if list.first.next != nil {
		t.Error("did not delete properly. First.next was wrong.")
	}

	list.Delete(0)

	// produces empty list

	if list.first != nil {
		t.Error("did not delete properly. Was not empty.")
	}

	if list.last != nil {
		t.Error("did not delete properly. Last was not nil")
	}

}

