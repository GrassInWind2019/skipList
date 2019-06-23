// skipList project example.go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/skipList/src/skipList"
	"github.com/skipList/src/userDefined"
)

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func searchRangeExample(s *skipList.SkipList) {
	var obj, obj2 userDefined.Obj
	obj.Data = 0
	obj2.Data = 30
	sliceObj, err := s.SearchRange(&obj, &obj2)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("search range result:")
	for _, sobj := range sliceObj {
		fmt.Printf("%d ", sobj.Data)
	}
	fmt.Println()
}

func operationsExample(s *skipList.SkipList) {
	var obj userDefined.Obj
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		insertObj := new(userDefined.Obj)
		insertObj.Data = rand.Intn(50)
		t, err := s.Insert(insertObj)
		if t == true {
			fmt.Println("insert obj ", insertObj.Data, " success")
		} else {
			fmt.Printf("insert obj %d failed: ", insertObj.Data, err)
		}
		//sleep 10ms
		time.Sleep(10000000)
		rand.Seed(time.Now().UnixNano())
		obj.Data = rand.Intn(50)
		//search and delete a random generated data
		_, err = s.Search(&obj)
		_, err2 := s.RemoveNode(&obj)
		if err == nil && err2 != nil {
			fmt.Printf("remove obj %d failed: ", obj.Data, err2)
			fmt.Println()
		} else {
			fmt.Printf("remove obj %d success\n", obj.Data)
		}
	}
}

func main() {
	minObj := new(userDefined.Obj)
	minObj.Data = INT_MIN
	s, err := skipList.CreateSkipList(minObj, 10)
	if s == nil {
		fmt.Println("create skip list failed: ", err)
		return
	}

	operationsExample(s)
	searchRangeExample(s)
	fmt.Println("start print the skip list")
	s.Traverse()
}
