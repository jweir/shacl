package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestAddingToMemory(t *testing.T) {
	i := Item{Title: "Hello World"}
	m := CreateMemory("test.json")

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
	m := CreateMemory("test.json")

	m.Add(i)

	if len(m.UnreadItems) != 1 {
		t.Error("item not added")
	}

	m.Remove(i)

	if len(m.UnreadItems) != 0 {
		t.Error("item not removed")
	}

	m.Add(i)

	if len(m.UnreadItems) != 0 {
		t.Error("item readded")
	}

}

func TestSavingAndLoadingMemory(t *testing.T) {
	i := &Item{Title: "Hello World"}
	m := CreateMemory("test.json")

	m.Add(i)
	m.Save()

	b := CreateMemory("test.json")

	if len(b.UnreadItems) != 0 {
		t.Error("what is going here?")
	}

	b.Load()

	if len(b.UnreadItems) != 1 {
		t.Error("Memory not loaded")
	}

	m.Destroy()
}

func TestParsing(t *testing.T) {
	b, e := ioutil.ReadFile("sample.xml")

	if e != nil {
		log.Fatal(e)
	}

	doc := parse(b)

	if len(doc.Items) != 3 {
		t.Errorf("wrong items %d instead of 3", len(doc.Items))
	}

}
