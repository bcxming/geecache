// LRU 缓存淘汰策略
package LRU
import "container/list"

//go语言实现LRU算法
type Cache struct{
	max_capacity int64
	cur_capacity int64
	mylist *list.List  //双向链表
	mymap map[string]*list.Element //哈希表 键是字符串，值是双向链表中对应节点的指针 map<string,node*>C++
	callback func (key string , value Value) //回调函数
}

type entry struct{
	key string
	value Value
}

//这个接口可以被其他类型实现，只要这些类型提供了 Len() 方法的实现
type Value interface{
	Len() int
}

//实例化
func New(m_cap int64,cb func (string ,Value))*Cache{
	//在C++中算返回局部变量的引用了
	return &Cache{
		max_capacity: m_cap,
		mylist: list.New(),
		mymap: make(map[string]*list.Element),
		callback:cb,
	}
}

//查找  在 Go 语言中，函数可以关联到类型上，形成方法
func (c *Cache) Get(key string)(value Value,ok bool){
	if ele,ok:=c.mymap[key];ok{
		c.mylist.MoveToFront(ele)
		kv:=ele.Value.(*entry)
		return kv.value,true
	}
	return 
}

//删除
func (c *Cache)RemoveOld(){
	ele:=c.mylist.Back()
	if ele != nil{
		c.mylist.Remove(ele)
		kv:=ele.Value.(*entry)
		delete(c.mymap,kv.key)
		c.cur_capacity -= int64(len(kv.key))+int64(kv.value.Len())
		if c.callback!=nil{
			c.callback(kv.key,kv.value)
		}
	}
}

//添加
func (c *Cache)Add(key string,value Value){
	if ele,ok:=c.mymap[key];ok{
		c.mylist.MoveToFront(ele)
		kv:=ele.Value.(*entry)
		c.cur_capacity+=int64(value.Len())-int64(kv.value.Len())
		kv.value = value
	}else{
		ele:=c.mylist.PushFront(&entry{key,value})
		c.mymap[key] = ele
		c.cur_capacity+=int64(len(key))+int64(value.Len())
	}
	for c.max_capacity!=0&&c.max_capacity<c.cur_capacity{
		c.RemoveOld()
	}
}

func (c *Cache) Len()int{
	return c.mylist.Len()
}