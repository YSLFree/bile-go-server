package datautil

//List  集合
type List interface {
	newArray() []interface{}
	get(index int) interface{}
	add(index int, element interface{})
	isEmpty() bool
	isContains(element interface{}) bool
	remove(index int)
	removeAll()
}

// ArrayList 集合
type ArrayList struct {
	List []interface{}
	size int
}

//NewArrayList 创建一个新的默认长度集合
func NewArrayList() *ArrayList {
	var list = &ArrayList{}
	list.List = make([]interface{}, 10)
	list.size = 0
	return list
}

//NewArraySizeList 创建一个制定长度的集合
func NewArraySizeList(size int32) *ArrayList {
	var list = &ArrayList{}
	list.List = make([]interface{}, size)
	list.size = 0
	return list
}

//Size  集合长度
func (list *ArrayList) Size() int {
	return len(list.List)
}

//Get 根据集合索引获取单个元素
func (list *ArrayList) Get(index int) interface{} {
	if index < 0 || index > list.size-1 {
		panic("ArratList index out of bround!")
	} else {
		return list.List[index]
	}
}

//Add 根据索引添加单个元素
func (list *ArrayList) Add(index int, element interface{}) {
	if index < 0 || index > list.size {
		panic("ArratList index out of bround!")
	} else {
		v := cap(list.List)
		if index == 0 {
			list.List = append([]interface{}{element}, list.List...)
		} else if v == index {
			list.List = append(list.List, element)
		} else {
			var li = list.List
			list.List = append(li[:index], element)
			list.List = append(list.List, li[index+1:]...)
		}
		list.size = list.size + 1
	}
}

//IsEmpty 判断集合是否为空
func (list *ArrayList) IsEmpty() bool {
	return len(list.List) == 0 
}

//IsContains 判断集合是否含有某个元素
func (list *ArrayList) IsContains(element interface{}) bool {
	for i := 0; i < list.size; i++ {
		if element == list.List[i] {
			return true
		}
	}
	return false
}

//Remove 根据索引移除集合中某个元素
func (list *ArrayList) Remove(index int) {
	if list.size == 0 {
		return
	}
	if index < 0 || index > list.size-1 {
		panic("ArratList index out of bround!")
	} else {
		var li = list.List
		list.List = append(li[:index-1], li[index+1:])
		list.size = list.size - 1
	}
}

//RemoveAll 移除集合中所有元素
func (list *ArrayList) RemoveAll() {
	list.List = make([]interface{}, 10)
	list.size = 0
}
