package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Memory struct {
	// keeps a log of all Item.Signatures ever seen
	Index map[string]bool

	// The list of unread Items, once read they are gone for good
	UnreadItems map[string]*Item
}

func CreateMemory() *Memory {
	m := &Memory{}
	m.UnreadItems = make(map[string]*Item)
	m.Index = make(map[string]bool)

	// load Index
	// load the UnReadItems

	return m
}

func (m *Memory) Save() bool {
	b, e := json.Marshal(m)

	if e != nil {
		log.Fatal(e)
	}

	e = ioutil.WriteFile("memory.json", b, 0666)

	return true
}

func (m *Memory) Load() bool {
	b, e := ioutil.ReadFile("memory.json")
	if e != nil {
		return false
	}

	json.Unmarshal(b, m)
	return true
}

func (m *Memory) Add(i *Item) bool {
	i.Sig()

	if _, ok := m.Index[i.Signature]; ok == false {
		m.UnreadItems[i.Signature] = i
		m.Index[i.Signature] = true
		return true
	} else {
		return false
	}
}

func (m *Memory) Remove(i *Item) bool {
	delete(m.UnreadItems, i.Signature)
	return true
}
