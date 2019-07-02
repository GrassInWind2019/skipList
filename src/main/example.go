// skipList project example.go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/skipList/src/skipList"
)

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

type myInt int

func (a *myInt) Compare(b skipList.SkipListObj) bool {
	return *a < *b.(*myInt)
}

func (a *myInt) PrintObj() {
	fmt.Print(*a)
}

func searchRangeExample(s *skipList.SkipList) {
	var minObj, maxObj myInt
	minObj = 0
	maxObj = 30
	sliceObj, err := s.SearchRange(&minObj, &maxObj)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("search range result:")
	for _, sobj := range sliceObj {
		fmt.Printf("%d ", *sobj.(*myInt))
	}
	fmt.Println()
}

func operationsExample(s *skipList.SkipList) {
	var obj myInt
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		insertObj := new(myInt)
		*insertObj = myInt(rand.Intn(50))
		t, err := s.Insert(insertObj)
		if t == true {
			fmt.Println("insert obj ", *insertObj, " success")
		} else {
			fmt.Printf("insert obj %d failed: ", *insertObj, err)
		}
		//sleep 1ms
		time.Sleep(1000000)
		rand.Seed(time.Now().UnixNano())
		obj = myInt(rand.Intn(50))
		//search and delete a random generated data
		s.Search(&obj)
		s.RemoveNode(&obj)
		_, err = s.Search(&obj)
		_, err2 := s.RemoveNode(&obj)
		if err == nil && err2 != nil {
			fmt.Printf("remove obj %d failed: ", obj, err2)
			fmt.Println()
		} else {
			fmt.Printf("remove obj %d success\n", obj)
		}
	}
}

func testSkipList(wg *sync.WaitGroup, s *skipList.SkipList, i int) {
	defer wg.Done()

	operationsExample(s)
	searchRangeExample(s)
	fmt.Printf("Goroutine %d start print the skip list\n", i)
	s.Traverse()
	length, err := s.LenOfSkipList()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Goroutine %d skip list length is %d\n", i, length)
}

func main() {
	wg := sync.WaitGroup{}

	minObj := new(myInt)
	*minObj = myInt(INT_MIN)
	s, err := skipList.CreateSkipList(minObj, 10, 2)
	if s == nil {
		fmt.Println("create skip list failed: ", err)
		return
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go testSkipList(&wg, s, i)
	}
	//wait all goroutines finish
	wg.Wait()
}
