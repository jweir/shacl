package main

import "testing"

func TestAddingToMemory(t *testing.T) {
	i := Item{Title: "Hello World"}
	m := CreateMemory()

	m.Add(&i)

	if len(m.UnreadItems) != 1 {
		t.Error("item not added")
	}

	m.Add(&i)

	if len(m.UnreadItems) != 1 {
		t.Error("item added twice")
	}

}

func TestReadingAnItem(t *testing.T) {
	i := &Item{Title: "Hello World"}
	m := CreateMemory()

	m.Add(i)

	if len(m.UnreadItems) != 1 {
		t.Error("item not added")
	}

	m.Remove(i)

	if len(m.UnreadItems) != 0 {
		t.Error("item not removed")
	}
}

func TestSavingAndLoadingMemory(t *testing.T) {
	i := &Item{Title: "Hello World"}
	m := CreateMemory()

	m.Add(i)
	m.Save()

	b := CreateMemory()

	if len(b.UnreadItems) != 0 {
		t.Error("what is going here?")
	}

	b.Load()

	if len(b.UnreadItems) != 1 {
		t.Error("Memory not loaded")
	}

}
