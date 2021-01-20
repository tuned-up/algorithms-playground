package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	minRand = 1
	maxRand = 100
)

type OpenAdressableSortedHash struct {
	bucketsNum int
	storage []*HashValue
}

func (oh OpenAdressableSortedHash) bucketLen() int {
	return cap(oh.storage) / oh.bucketsNum
}

func (oh OpenAdressableSortedHash) Add(v *HashValue) bool {
	bucket := oh.HashFunc(v.key)

	for i := 0; i < cap(oh.storage); i++ {
		probe := bucket * oh.bucketLen() + i
		hashValue := oh.storage[probe]

		if hashValue == nil {
			oh.storage[probe] = v
			return true
		}
		if hashValue.key == v.key {
			hashValue.value = v.value
			return true
		}
		if hashValue.key > v.key {
			temp := hashValue
			oh.storage[probe] = v
			oh.Add(temp)
			return true
		}

		continue
	}

	return false
}

func (oh OpenAdressableSortedHash) Find(key int) (int, bool) {
	bucket := oh.HashFunc(key)

	for i := 0; i < cap(oh.storage); i++ {
		probe := bucket * oh.bucketLen() + i
		if oh.storage[probe] == nil {
			return 0, false
		}
		hashValue := oh.storage[probe]
		if hashValue.deleted {
			continue
		}
		if hashValue.key > key {
			return 0, false
		}
		if hashValue.key == key {
			return hashValue.value, true
		}

		continue
	}

	return 0, false
}

//Ideally, this need rehashing if there are too many deleted entries
//but I am too lazy

//Too lazy to refactor Del and Find too
func (oh OpenAdressableSortedHash) Del(key int) (int, bool) {
	bucket := oh.HashFunc(key)

	for i := 0; i < cap(oh.storage); i++ {
		probe := bucket * oh.bucketLen() + i
		if oh.storage[probe] == nil {
			return 0, false
		}
		hashValue := oh.storage[probe]
		if hashValue.deleted {
			continue
		}
		if hashValue.key > key {
			return 0, false
		}
		if hashValue.key == key {
			hashValue.deleted = true
			return hashValue.value, true
		}

		continue
	}

	return 0, false
}

func (oh OpenAdressableSortedHash) HashFunc(key int) int {
	return key % (cap(oh.storage) / oh.bucketsNum)
}

func (oh OpenAdressableSortedHash) Print() {
	for i := 0; i < len(oh.storage); i++ {
		fmt.Printf(" %v ", oh.storage[i])
		if (i + 1) % oh.bucketLen() == 0 && i != 0 {
			fmt.Println()
		}
	}
}

type HashValue struct {
	key int
	value int
	deleted bool
}

func main()  {
	rand.Seed(time.Now().Unix())

	oh := OpenAdressableSortedHash{
		bucketsNum: 10,
		storage:    make([]*HashValue, 100, 100),
	}

	var toDelete int
	for i := 0; i < randomInt(); i++ {
		hv := &HashValue{key: randomInt(), value: randomInt()}
		if i == 0 {
			toDelete = hv.key
		}
		oh.Add(hv)
	}

	oh.Del(toDelete)
	oh.Print()
	fmt.Println()

	for i := 0; i < 10; i++ {
		key := randomInt()
		value, found := oh.Find(key)
		fmt.Printf("Under key %d: %d %t\n", key, value, found)
	}
	value, found := oh.Find(toDelete)
	fmt.Printf("Under deleted key %d: %d %t\n", toDelete, value, found)
}

func randomInt() int {
	return rand.Intn(maxRand - minRand) + minRand
}
