// skipList_test project skipList_test.go
package skipList

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/skipList/src/userDefined"
)

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func TestCreateSkipList(t *testing.T) {
	var obj userDefined.Obj
	obj.Data = INT_MIN
	s, err := CreateSkipList(&obj, 10)
	if s == nil {
		fmt.Print(err)
		t.Errorf("create list failed")
	}
}

func TestOperations(t *testing.T) {
	var obj userDefined.Obj
	obj.Data = INT_MIN
	s, err := CreateSkipList(&obj, 10)
	if s == nil {
		fmt.Print(err)
		t.Errorf("create skip list failed")
	}

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		obj.Data = rand.Intn(50)
		res, err := s.Insert(&obj)
		if res == true {
			fmt.Println("insert obj ", obj.Data, " success")
		} else {
			fmt.Print(err)
			t.Errorf("insert obj %d failed: ", obj.Data)
		}
		//sleep 10ms
		time.Sleep(10000000)
		rand.Seed(time.Now().UnixNano())
		obj.Data = rand.Intn(50)
		_, err = s.Search(&obj)
		_, err2 := s.RemoveNode(&obj)
		if err == nil && err2 != nil {
			fmt.Print(err)
			t.Errorf("remove obj %d failed: ", obj.Data)
		} else {
			fmt.Printf("remove obj %d success\n", obj.Data)
		}
	}
	fmt.Println("start print the skip list")
	s.Traverse()
}
