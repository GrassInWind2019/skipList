// skipList project skipList.go
package skipList

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/skipList/src/userDefined"
)

type SkipListObj interface {
	Compare(obj SkipListObj) bool
	PrintObj()
}

type Node struct {
	O       SkipListObj
	forward []*Node
}

type SkipList struct {
	head   *Node
	length int
}

var maxLevel int

//try to find the first node which not match the user defined Compare() condition
func (s *SkipList) searchInternal(o SkipListObj) (*Node, error) {
	p := s.head
	if p == nil {
		return nil, errors.New("skip list head is null, must use CreateSkipList() before Search")
	}

	for i := maxLevel - 1; i >= 0; i-- {
		for {
			if p.forward[i] != nil && p.forward[i].O.Compare(o) {
				p = p.forward[i]
			} else {
				break
			}
		}
	}
	p = p.forward[0]
	if p == nil {
		return nil, errors.New("No matched object in skip list")
	} else {
		return p, nil
	}
}

func (s *SkipList) Search(o SkipListObj) (SkipListObj, error) {
	if s == nil {
		return o, errors.New("skiplist pointer is nil")
	}

	res, err := s.searchInternal(o)
	if err == nil {
		if !res.O.Compare(o) && !o.Compare(res.O) {
			return res.O, nil
		} else {
			return o, errors.New("cannot find object in skip list")
		}
	} else {
		return o, err
	}
}

func (s *SkipList) SearchRange(minObj, maxObj SkipListObj) ([]SkipListObj, error) {
	res := make([]SkipListObj, 0)
	if s == nil {
		return res, errors.New("skip list pointer is nil")
	}

	p, err := s.searchInternal(minObj)
	if err != nil {
		return res, err
	}

	for {
		if p != nil && p.O.Compare(maxObj) {
			res = append(res, p.O)
			p = p.forward[0]
		} else {
			break
		}
	}

	return res, nil
}

func (s *SkipList) Traverse() {
	if s == nil {
		return
	}

	var p *Node = s.head
	if p == nil {
		return
	}

	for i := maxLevel - 1; i >= 0; i-- {
		for {
			if p != nil {
				p.O.PrintObj()
				if p.forward[i] != nil {
					fmt.Print("-->")
				}
				p = p.forward[i]
			} else {
				break
			}
		}
		fmt.Println()
		p = s.head
	}
}

func (s *SkipList) Insert(obj SkipListObj) (bool, error) {
	var p *Node = s.head
	if s.head == nil {
		return false, errors.New("skip list head is null, must use CreateSkipList() before insert")
	}
	newNode := new(Node)
	newNode.O = obj
	newNode.forward = make([]*Node, maxLevel)
	level := createNewNodeLevel()

	for i := level; i >= 0; i-- {
		for {
			if p.forward[i] != nil && p.forward[i].O.Compare(obj) {
				p = p.forward[i]
			} else {
				break
			}
		}
		//find the last Node which match user defined Compare() condition in i level
		newNode.forward[i] = p.forward[i]
		p.forward[i] = newNode
	}
	return true, nil
}

func (s *SkipList) RemoveNode(obj SkipListObj) (bool, error) {
	var update []*Node = make([]*Node, maxLevel)
	p := s.head

	if s == nil || s.head == nil {
		return false, errors.New("skip list is null, nothing to remove")
	}

	for i := maxLevel - 1; i >= 0; i-- {
		for {
			if p.forward[i] != nil && p.forward[i].O.Compare(obj) {
				p = p.forward[i]
			} else {
				break
			}
		}
		update[i] = p
	}
	p = p.forward[0]

	if p == nil || p.O.Compare(obj) || obj.Compare(p.O) {
		return false, errors.New("cannot find object")
	}

	for i := maxLevel - 1; i >= 0; i-- {
		update[i].forward[i] = p.forward[i]
	}

	return true, nil
}

func CreateSkipList(minObj SkipListObj, maxlevel int) (*SkipList, error) {
	if minObj == nil {
		return nil, errors.New("minObj paramter is null")
	}
	if maxlevel <= 0 {
		return nil, errors.New("Max level of skip list must greater than 0")
	}

	maxLevel = maxlevel
	s := new(SkipList)
	s.head = new(Node)
	s.head.forward = make([]*Node, maxLevel)
	s.head.O = minObj

	return s, nil
}

func createNewNodeLevel() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxLevel - 1)
}
