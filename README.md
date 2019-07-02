# 目录  
- [跳表简介](#跳表简介)
- [skipList](#skiplist)
  * [跳表示意图](#跳表示意图)
  * [skip list实现简要介绍](#skip list实现简要介绍)
  * [skip list使用介绍](#skip list使用介绍)

# 跳表简介  
跳跃链表是对链表的改进，链表虽然节省空间，但是查找时间复杂度为O（n），效率比较低。
跳跃链表的查找、插入、删除等操作的期望时间复杂度为O(logn)，效率比链表提高了很多。
跳表的基本性质
1. 有很多层结构组成
2. 每一层都是一个有序的链表
3. 最底层(level 0)的链表包含所有元素
4. 如果一个元素出现在level i的链表中，则在level i之下的链表都会出现
5. 每个节点包含两个指针，一个指向同一链表的下一个元素，一个指向下面一层的元素
  
# skipList  
## 跳表示意图  
![图片不能显示](https://github.com/GrassInWind2019/skipList/blob/master/skipList.png)
## skip list实现简要介绍  
本人实现的skipList结构如下(在src/skipList/skipList.go)  
```
type Node struct {  
	O        SkipListObj  
	forward  []*Node  
	curLevel int  
}

type SkipList struct {  
	head     *Node  
	length   int  
	maxLevel int  
	lockType int  
	lock     sync.Locker  
}  
```  
目前只实现了针对int类型数据的接口，若使用其他类型，需要实现下面这个接口(示例实现见src/main/example.go)  
```
type SkipListObj interface {  
	Compare(obj SkipListObj) bool  
	PrintObj()  
}  
```  
其中Compare()表示数据大小比较，   
一开始本想用Compare(obj interface{}) bool作为接口，但是编译报错，提示type interface {} is interface with no methods，于是改用了上面的接口  
由于Go内置类型用户不能新建接口（会编译失败，对于int类型提示cannot define new methods on non-local type int），可以通过type给int取个别名来实现接口  
```
type myInt int  

func (a *myInt) Compare(b skipList.SkipListObj) bool {  
	return *a < *b.(*myInt)  
}  
```  
PrintObj()用于打印数据的，主要是用于遍历显示跳表的Traverse()，对于int类型如下 
```
func (a *myInt) PrintObj() {  
	fmt.Print(*a)  
} 
```  
CreateSkipList用于创建一个skip list，需要传入一个对任意其他对象Compare()函数都返回false的对象（例如对于example的实现就是传入最小的int整数），  
还需要传入一个表示skip list最大层数以及一个可选的参数mode表示创建的skip list的锁类型，若  
mode = 1, 则为互斥锁  
mode = 2, 则为读写锁  
其他为无锁  
func CreateSkipList(minObj SkipListObj, args ...int) (*SkipList, error)  
  
lockList主要用于在操作skip list前上锁，可根据配置来决定使用：
1. 无锁  
2. 互斥锁  
3. 读写锁 
```
func (s *SkipList) lockList(mode int) {  
	if s.lockType == 0 {  
		return  
	}  
	switch mode {  
	case 1:  
		s.lock.Lock()  
	case 2:  
		if s.lockType == 1 {  
			s.lock.Lock()  
		} else if s.lockType == 2 {  
			s.lock.(*sync.RWMutex).RLock()  
		}  
	default:  
		return  
	}  
}  
```
  
unLockList相应地解锁
func (s *SkipList) unLockList(mode int)

Search()用于在skip list中查找指定的对象  
func (s *SkipList) Search(o SkipListObj) (SkipListObj, error)  

SearchRange()用于在skip list中按照指定范围查找对象  
func (s *SkipList) SearchRange(minObj, maxObj SkipListObj) ([]SkipListObj, error)  

Insert()用于在skip list中插入一个对象
func (s *SkipList) Insert(obj SkipListObj) (bool, error)

createNewNodeLevel()用于Insert()方法中，为新对象产生其所在的层级  
```
func (s *SkipList) createNewNodeLevel() int {  
	var level int = 0  

        rand.Seed(time.Now().UnixNano())  
        for {  
                if rand.Intn(2) == 1 || level >= s.maxLevel-1 {  
                        break  
                }  
                level++  
        }  
        return level  
} 
```  
Traverse()用于遍历打印skip list  
func (s *SkipList) Traverse()  
```
func (s *SkipList) Traverse() {
	v, _ := checkSkipListValid(s)
	if v == false {
		return
	}

	var p *Node = s.head

	s.lockList(2)
	defer s.unLockList(2)

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
```  
  
##  skip list使用介绍
首先必须调用CreateSkipList()创建一个skip list  
创建一个最大10层无锁skip list  
s, err := skipList.CreateSkipList(minObj, 10）  
创建一个最大10层使用互斥锁的skip list  
s, err := skipList.CreateSkipList(minObj, 10, 1)  
创建一个最大10层使用读写锁的skip list  
s, err := skipList.CreateSkipList(minObj, 10, 2)  
  
插入一个对象(以myInt为例)  
```
insertObj := new(myInt)
*insertObj = myInt(30)
t, err := s.Insert(insertObj)
if t == true {
	fmt.Println("insert obj ", *insertObj, " success")
} else {
	fmt.Printf("insert obj %d failed: ", *insertObj, err)
}
```
范围搜索skip list示例(以myInt为例)
```
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
```
使用示例在src/main/example.go  
UT示例为src/skipList/skipList_test.go  

使用LiteIDE（一个轻量级的Go IDE）运行example结果如下：
![图片不能显示](https://github.com/GrassInWind2019/skipList/blob/master/example_run_result.png)
