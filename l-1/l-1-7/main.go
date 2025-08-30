package main

import (
	"fmt"
	"sync"
)

//Реализовать безопасную для конкуренции запись данных в структуру map.

type syncMap struct {
	storage map[string]interface{}
	// поменять на rwmutex если будет нагрузка на чтение в основном
	sync.Mutex
}

func newMap(size int) *syncMap {
	return &syncMap{}
}

func (m *syncMap) add(k string, v interface{}) {
	m.Lock()
	defer m.Unlock()
	m.storage[k] = v
}

func (m *syncMap) del(k string) {
	m.Lock()
	defer m.Unlock()
	delete(m.storage, k)
}

func (m *syncMap) get(k string) interface{} {
	m.Lock()
	defer m.Unlock()
	return m.storage[k]
}

func main() {
	person := newMap(5)

	person.add("name", "mike")
	person.add("age", 18)
	person.add("balance", 123.123)

	fmt.Println(person.storage)

	person.del("name")
	fmt.Println(person.storage)

	fmt.Println(person.get("balance"))
}
