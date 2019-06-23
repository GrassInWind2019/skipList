// userDefined project userDefined.go
package userDefined

import (
	"fmt"
)

type Obj struct {
	Data interface{}
}

type SkipListObj interface {
	Compare(a *Obj) bool
	PrintObj()
}

func (b *Obj) Compare(a *Obj) bool {
	return b.Data.(int) < a.Data.(int)
}

func (obj *Obj) PrintObj() {
	fmt.Print(obj.Data.(int))
}
