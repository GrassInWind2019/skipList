// skipList project skipList.go
package skipList

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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
	head     *Node
	length   int
	maxLevel int
}

//try to find the first node which not match the user defined Compare() condition
func (s *SkipList) searchInternal(o SkipListObj) (*Node, error) {
	p := s.head
	if p == nil {
		return nil, errors.New("skip list head is null, must use CreateSkipList() before Search")
	}

	for i := s.maxLevel - 1; i >= 0; i-- {
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

	for i := s.maxLevel - 1; i >= 0; i-- {
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
	newNode.forward = make([]*Node, s.maxLevel)
	level := s.createNewNodeLevel()

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
	s.length++

	return true, nil
}

func (s *SkipList) RemoveNode(obj SkipListObj) (bool, error) {
	if s == nil || s.head == nil {
		return false, errors.New("skip list is null, nothing to remove")
	}
	var update []*Node = make([]*Node, s.maxLevel)
	p := s.head

	for i := s.maxLevel - 1; i >= 0; i-- {
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

	for i := s.maxLevel - 1; i >= 0; i-- {
		update[i].forward[i] = p.forward[i]
	}
	s.length--

	return true, nil
}

func (s *SkipList) ClearSkipList() error {
	if s == nil {
		return errors.New("skip list not exist")
	}
	if s.head == nil {
		return errors.New("skip list head is null, must use CreateSkipList() to create skip list")
	}

	for i := s.maxLevel; i >= 0; i-- {
		s.head.forward[i] = nil
	}
	s.length = 0

	return nil
}

func (s *SkipList) LenOfSkipList() (int, error) {
	if s == nil || s.head == nil {
		return 0, errors.New("skip list is null")
	}

	return s.length, nil
}

func CreateSkipList(minObj SkipListObj, maxlevel int) (*SkipList, error) {
	if minObj == nil {
		return nil, errors.New("minObj paramter is null")
	}
	if maxlevel <= 0 {
		return nil, errors.New("Max level of skip list must greater than 0")
	}

	s := new(SkipList)
	s.head = new(Node)
	s.maxLevel = maxlevel
	s.head.forward = make([]*Node, maxlevel)
	s.head.O = minObj
	//The length of skip list didn't include the head node
	s.length = 0

	return s, nil
}

func (s *SkipList) createNewNodeLevel() int {
	var level int = 0

	rand.Seed(time.Now().UnixNano())
	for {
		if rand.Intn(2) == 1 {
			break
		}
		level++
	}
	return level
}
