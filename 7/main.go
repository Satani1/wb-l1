package main

import (
	"errors"
	"log"
	"sync"
)

type ConcMap struct {
	sync.RWMutex
	Numbers map[int]int
}

func NewConcMap() ConcMap {
	return ConcMap{
		Numbers: make(map[int]int, 0),
	}
}

func (cm *ConcMap) Get(key int) (int, error) {
	//При .RLock() одновременно несколько горутин могут читать из мапы, тк
	//.RLock() блокирует все вызовы .Lock(), но вызовы с .RLock() пройдут спокойно
	cm.RLock()
	defer cm.RUnlock()

	value, ok := cm.Numbers[key]
	if ok {
		return value, nil
	}
	return 0, errors.New("Value doesn't exist ")
}

func (cm *ConcMap) Write(value, key int) {
	//.Lock() блокирует все остальные вызовы и заставляет другие горутины ждать,
	//пока запись в мапу не будет произведена
	cm.Lock()
	defer cm.Unlock()

	cm.Numbers[key] = value
}

func main() {
	wg := &sync.WaitGroup{}
	numsMap := NewConcMap()

	times := 5

	//запись в мапу
	for i := 0; i < times; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			value := (i + 1) * 10
			numsMap.Write(value, i)
		}(i)
	}

	//чтение мапы
	for i := 0; i < times; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			value, err := numsMap.Get(i)
			if err != nil {
				log.Printf("ERROR: %v\n", err)
			} else {
				log.Printf("Value '%v' with key '%v'\n", value, i)
			}
		}(i)
	}

	wg.Wait()
}
