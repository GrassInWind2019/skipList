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
跳表示意图
![图片不能显示](https://github.com/GrassInWind2019/skipList/skipList.png)

本人实现的skipList结构如下(在src/skipList/skipList.go)  
type Node struct {  
	O       *userDefined.Obj  
	forward []*Node  
}  

type SkipList struct {  
	head   *Node  
	length int  
}  

目前只实现了针对int类型数据的接口，若使用其他类型，需要实现下面这个接口(在src/userDefined/userDefined.go)  
type SkipListObj interface {  
	Compare(obj SkipListObj) bool  
	PrintObj()  
}  
其中Compare()表示数据大小比较，对于int类型如下   
由于Go内置类型用户不能新建接口（会编译失败），可以通过type给int取个别名来实现接口  
type myInt int  

func (a *myInt) Compare(b skipList.SkipListObj) bool {  
	return *a < *b.(*myInt)  
}  

PrintObj()用于打印数据的，主要是用于遍历显示跳表的Traverse()，对于int类型如下  
func (a *myInt) PrintObj() {  
	fmt.Print(*a)  
} 

使用示例在src/main/example.go  
UT示例为src/skipList/skipList_test.go  

使用LiteIDE（一个轻量级的Go IDE）运行example结果如下：
![图片不能显示](https://github.com/GrassInWind2019/skipList/example_run_result.png)
